package grpcserver

import (
	"context"
	"fmt"
	"github.com/azamt18/post-service-grpc-api-gateway/db"
	post_operations "github.com/azamt18/post-service-grpc-api-gateway/services/post/inner/operations/internals/protobuff/post.operations.v1"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	db db.Database
}

func (s server) ListPosts(request *post_operations.ListPostsRequest, stream post_operations.PostOperationsService_ListPostsServer) error {
	fmt.Println("Get posts request...")

	cursor, err := s.db.PostsCollection().Find(context.Background(), primitive.D{{}}) // D - used because of the order of the elements matters
	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal err: %v", err),
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
		data := &PostItem{}
		if err := cursor.Decode(data); err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error while decoding data from MongoDB: %v", err),
			)
		}

		// send a blog via stream
		if err := stream.Send(&post_operations.ListPostsResponse{Post: dataToPostPb(data)}); err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error while sending a stream: %v", err),
			)
		}
	}

	if err := cursor.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal err: %v", err),
		)
	}

	return nil
}

func (s server) ReadPost(ctx context.Context, request *post_operations.ReadPostRequest) (*post_operations.ReadPostResponse, error) {
	fmt.Println("Read post request...")

	postId := request.GetPostId()
	oid, err := primitive.ObjectIDFromHex(postId)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse PostId"),
		)
	}

	// create an empty struct
	data := &PostItem{}
	filter := bson.M{"_id": oid, "id": postId} // Also apply a filter collection by post id("id" int64)

	// perform search operation
	result := s.db.PostsCollection().FindOne(ctx, filter)
	if error := result.Decode(data); error != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find post with specified PostId: %v", error),
		)
	}

	// prepare response
	response := &post_operations.ReadPostResponse{
		Post: dataToPostPb(data),
	}

	return response, nil
}

func (s server) UpdatePost(ctx context.Context, request *post_operations.UpdatePostRequest) (*post_operations.UpdatePostResponse, error) {
	fmt.Println("Update blog request...")
	post := request.GetPost()
	oid, error := primitive.ObjectIDFromHex(string(post.GetId()))
	if error != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse PostId"),
		)
	}

	// create an empty struct
	data := &PostItem{}
	filter := bson.M{"_id": oid}

	result := s.db.PostsCollection().FindOne(ctx, filter)
	if error := result.Decode(data); error != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Can not find a post with given PostId: %v", error),
		)
	}

	// perform update operation
	data.UserId = post.GetUserId()
	data.Title = post.GetTitle()
	data.Body = post.GetBody()

	_, updateError := s.db.PostsCollection().ReplaceOne(context.Background(), filter, data)
	if updateError != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Can not update object in the Db: %v", updateError),
		)
	}

	// prepare response
	response := &post_operations.UpdatePostResponse{
		Post: dataToPostPb(data),
	}

	return response, nil
}

func (s server) DeletePost(ctx context.Context, request *post_operations.DeletePostRequest) (*post_operations.DeletePostResponse, error) {
	fmt.Println("Delete blog request...")
	oid, error := primitive.ObjectIDFromHex(request.GetPostId())
	if error != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse PostId"),
		)
	}

	filter := bson.M{"_id": oid}
	deleteResult, deleteError := s.db.PostsCollection().DeleteOne(context.Background(), filter)
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

	return &post_operations.DeletePostResponse{
		PostId: request.GetPostId(),
	}, nil
}

func dataToPostPb(data *PostItem) *post_operations.Post {
	return &post_operations.Post{
		Id:     data.Id,
		UserId: data.UserId,
		Title:  data.Title,
		Body:   data.Body,
	}
}

type PostItem struct {
	Id     int64  `json:"_id,omitempty" bson:"_id,omitempty"`
	UserId int64  `bson:"user_id"`
	Title  string `bson:"title"`
	Body   string `bson:"body"`
}

func NewGrpcServer() post_operations.PostOperationsServiceServer {
	return &server{}
}
