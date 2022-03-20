package grpcserver

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/azamt18/post-service-grpc-api-gateway/db"
	post_loader "github.com/azamt18/post-service-grpc-api-gateway/services/post/inner/loader/internals/protobuff/post.loader.v1"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type server struct {
	goRestApiHost string
	database      db.Database
}

var response []Post

type Post struct {
	PostId int64  `json:"id,omitempty" bson:"_id,omitempty"`
	UserId int64  `json:"user_id,omitempty"`
	Title  string `json:"title,omitempty"`
	Body   string `json:"body,omitempty"`
}

func (s server) LoadPosts(ctx context.Context, request *post_loader.LoadPostsRequest) (*post_loader.LoadPostsResponse, error) {
	fmt.Println("Load posts request...")

	//todo implement async downloading via go routines

	// make an API request to load posts
	params := url.Values{}
	params.Add("page", strconv.Itoa(int(request.GetPageNumber())))

	resp, err := http.Get(s.goRestApiHost + "/public/v2/posts?" + params.Encode())
	if err != nil {
		log.Printf("Request Failed: %s", err)
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	// Log the request body
	bodyString := string(body)
	log.Print(bodyString)

	// Unmarshal result
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
		return nil, err
	}

	// Iterate over the values
	loadedPostsCount := int64(0)
	for i, post := range response {
		fmt.Printf("%v) Post title: %v\n", i, post.Title)
		postObject := Post{
			PostId: post.PostId,
			UserId: post.UserId,
			Title:  post.Title,
			Body:   post.Body,
		}

		// Saving to db
		res, err := s.database.PostsCollection().InsertOne(context.TODO(), postObject)
		if err != nil {
			log.Fatalf("Internal error: %v", err)
		}

		// Check the insertion result
		objectId, ok := res.InsertedID.(primitive.ObjectID)
		if !ok {
			log.Fatalf("Can not convert to ObjectId")
		}

		loadedPostsCount++
		log.Printf("Inserted objectId: %v", objectId)
	}

	return &post_loader.LoadPostsResponse{
		Success:          true,
		LoadedPostsCount: loadedPostsCount,
	}, nil
}

func NewGrpcServer(goRestApiHost string) post_loader.LoadPostsServiceServer {
	return &server{
		goRestApiHost: goRestApiHost,
	}
}
