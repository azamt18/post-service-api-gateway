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
	"sync"
)

const (
	goRestApiHost       = "GOREST_API_HOST"
	pageNumberToFetchTo = 50
)

type server struct {
	goRestApiHost string
	database      db.Database
}

type Post struct {
	PostId int64  `json:"id,omitempty" bson:"_id,omitempty"`
	UserId int64  `json:"user_id,omitempty"`
	Title  string `json:"title,omitempty"`
	Body   string `json:"body,omitempty"`
}

var posts []Post

func GetPosts(page int) ([]Post, error) {
	// make an API request to load posts
	params := url.Values{}
	params.Add("page", strconv.Itoa(page))

	resp, err := http.Get(goRestApiHost + "/public/v2/posts?" + params.Encode())
	if err != nil {
		log.Printf("Request Failed: %s", err)
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Log the request body
	bodyString := string(body)
	log.Print(bodyString)

	// Unmarshal result
	err = json.Unmarshal(body, &posts)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
		return nil, err.Error()
	}

	return posts, nil
}

func FetchPostsFromChannel(i int, ch chan<- []Post, wg *sync.WaitGroup) {
	defer wg.Done()

	getPostsRes, err := GetPosts(i)
	if err != nil {
		panic(err)
	}

	ch <- getPostsRes
}

func (s server) LoadPosts(ctx context.Context, request *post_loader.LoadPostsRequest) (*post_loader.LoadPostsResponse, error) {
	fmt.Println("Load posts request...")

	ch := make(chan []Post)
	responses := make([]Post, 0)
	var wg sync.WaitGroup
	loadedPostsCount := int64(0)

	// fetch data from 50 pages
	for i := 1; i <= pageNumberToFetchTo; i++ {
		wg.Add(1)
		go FetchPostsFromChannel(i, ch, &wg)
	}

	// close the channel in the background
	go func() {
		wg.Wait()
		close(ch)
	}()

	// read from channel as they come in until its closed
	for res := range ch {
		responses = append(responses, res...)
		{
			// Iterate over the values
			for i, post := range posts {
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
		}
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
