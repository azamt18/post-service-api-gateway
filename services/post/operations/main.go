package main

import (
	"context"
	"fmt"
	pb "github.com/azamt18/post-service-grpc-api-gateway/services/post/operations/api/goclient/v1"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
	"os/signal"
)

const (
	listenAddress  = "0.0.0.0:9090"
	dbHost         = "mongodb://localhost:27017"
	dbName         = "mydb"
	collectionName = "posts"
)

var collection *mongo.Collection

type postOperationsService struct{}

func (s *postOperationsService) ReadPost(ctx context.Context, request *pb.ReadPostRequest) (*pb.ReadPostResponse, error) {
	fmt.Println("Read post request...")

	postId := request.GetPostId()
	oid, error := primitive.ObjectIDFromHex(postId)
	if error != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse ID"),
		)
	}

	// create an empty struct
	data := &postItem{}
	filter := bson.M{"_id": oid} // NewDocument

	// perform find operation
	result := collection.FindOne(ctx, filter)
	if error := result.Decode(data); error != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find post with specified ID: %v", error),
		)
	}

	// prepare response
	response := &pb.ReadPostResponse{
		Post: dataToPostPb(data),
	}

	return response, nil
}

func (s *postOperationsService) UpdatePost(ctx context.Context, request *pb.UpdatePostRequest) (*pb.UpdatePostResponse, error) {
	fmt.Println("Update blog request...")
	post := request.GetPost()
	oid, error := primitive.ObjectIDFromHex(post.GetId())
	if error != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse ID"),
		)
	}

	// create an empty struct
	data := &postItem{}
	filter := bson.M{"_id": oid}

	result := collection.FindOne(ctx, filter)
	if error := result.Decode(data); error != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Can not find a post with given ID: %v", error),
		)
	}

	// perform update operation
	data.UserID = post.GetUserId()
	data.Title = post.GetTitle()
	data.Body = post.GetBody()

	_, updateError := collection.ReplaceOne(context.Background(), filter, data)
	if updateError != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Can not update object in the Db: %v", updateError),
		)
	}

	// prepare response
	response := &pb.UpdatePostResponse{
		Post: dataToPostPb(data),
	}

	return response, nil
}

func (s *postOperationsService) DeletePost(ctx context.Context, request *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
	fmt.Println("Delete blog request...")
	oid, error := primitive.ObjectIDFromHex(request.GetPostId())
	if error != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse ID"),
		)
	}

	filter := bson.M{"_id": oid}
	deleteResult, deleteError := collection.DeleteOne(context.Background(), filter)
	if deleteError != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Can not delete object in the Db: %v", deleteError),
		)
	}

	if deleteResult.DeletedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Can not find blog in the Db: %v", deleteError),
		)
	}

	return &pb.DeletePostResponse{
		PostId: request.GetPostId(),
	}, nil
}

func (s *postOperationsService) ListPosts(request *pb.ListPostsRequest, stream pb.PostService_ListPostsServer) error {
	fmt.Println("Get posts request...")

	cursor, error := collection.Find(context.Background(), primitive.D{{}}) // D - used because of the order of the elements matters
	if error != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", error),
		)
	}

	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			fmt.Printf("Error while closing a cursor: %v", err)
		}
	}(cursor, context.Background())

	for cursor.Next(context.Background()) {
		// create an empty struct for response
		data := &postItem{}
		if error := cursor.Decode(data); error != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error while decoding data from MongoDB: %v", error),
			)
		}

		// send a blog via stream
		if error := stream.Send(&pb.ListPostsResponse{Post: dataToPostPb(data)}); error != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error while sending a stream: %v", error),
			)
		}
	}

	if err := cursor.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", error),
		)
	}

	return nil
}

func dataToPostPb(data *postItem) *pb.Post {
	return &pb.Post{
		Id:     data.ID.Hex(),
		UserId: data.UserID,
		Title:  data.Title,
		Body:   data.Body,
	}
}

type postItem struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID string             `bson:"user_id"`
	Title  string             `bson:"title"`
	Body   string             `bson:"body"`
}

func main() {
	// if the go code is crushed -> get the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Connection to MongoDb")
	// Create client
	client, err := mongo.NewClient(options.Client().ApplyURI(dbHost))
	if err != nil {
		log.Fatal(err)
	}

	// Create connection
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database(dbName).Collection(collectionName)

	if err != nil {
		log.Fatalf("FAILED TO LISTEN %v", err)
	}

	var options []grpc.ServerOption
	s := grpc.NewServer(options...)
	pb.RegisterPostServiceServer(s, &postOperationsService{})

	log.Printf("Post operations service starting on %s", listenAddress)
	lis, err := net.Listen("tcp", listenAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	// Register reflection service on gRPC postOperationsService
	reflection.Register(s)

	go func() {
		fmt.Println("Starting Posts GRUD Server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Wait for Control C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block until a signal is received
	<-ch

	// 1st: close the connection with db
	fmt.Println("Closing MongoDb connection")
	client.Disconnect(context.TODO())

	// 2nd: close the listener
	fmt.Println("Closing the listener")
	lis.Close()

	// Finally, stop the postOperationsService
	fmt.Println("Stopping the postOperationsService")
	s.Stop()

	fmt.Println("End of Program")
}
