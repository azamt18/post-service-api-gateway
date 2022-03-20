package loader

import (
	post_loader_grpc_server "github.com/azamt18/post-service-grpc-api-gateway/services/post/inner/loader/internals/protobuff/post.loader.v1"
	"google.golang.org/grpc"
	"log"
	"os"
)

const (
	postLoaderGrpcServerAddress = "POST_LOADER_GRPC_SERVER_ADDRESS"
)

func NewClient() post_loader_grpc_server.LoadPostsServiceClient {
	// Set up connection with the grpc server
	conn, err := grpc.Dial(os.Getenv(postLoaderGrpcServerAddress), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error while making connection, %v", err)
	}

	// Create a client instance
	return post_loader_grpc_server.NewLoadPostsServiceClient(conn)
}
