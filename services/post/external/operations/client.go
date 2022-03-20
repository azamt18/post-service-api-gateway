package operations

import (
	post_operations_grpc_server "github.com/azamt18/post-service-grpc-api-gateway/services/post/inner/operations/internals/protobuff/post.operations.v1"
	"google.golang.org/grpc"
	"log"
	"os"
)

const (
	postOperationsGrpcServerAddress = "POST_OPERATIONS_GRPC_SERVER_ADDRESS"
)

func NewClient() post_operations_grpc_server.PostOperationsServiceClient {
	// Set up connection with the grpc server
	conn, err := grpc.Dial(os.Getenv(postOperationsGrpcServerAddress), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error while making connection, %v", err)
	}

	// Create a client instance
	return post_operations_grpc_server.NewPostOperationsServiceClient(conn)
}
