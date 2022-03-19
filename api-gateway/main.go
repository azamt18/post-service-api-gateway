package main

import (
	"context"
	"fmt"
	postOperationsSvcV1 "github.com/azamt18/post-service-grpc-api-gateway/services/post/operations/api/goclient/v1"
	"google.golang.org/grpc"
	//"github.com/golang/protobuf/protoc-gen-go/grpc"
	"log"
)

const (
	listenAddress     = "0.0.0.0:9090"
	postOperationsSvc = "users:9090"
)

func newPostOperationsSvcClient() (postOperationsSvcV1.PostServiceClient, error) {
	conn, err := grpc.DialContext(context.TODO(), postOperationsSvc, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("users client: %w", err)
	}

	return postOperationsSvcV1.NewPostServiceClient(conn), nil
}

func logger(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("method %q called\n", info.FullMethod)
	resp, err := handler(ctx, req)
	if err != nil {
		log.Printf("method %q failed: %s\n", info.FullMethod, err)
	}
	return resp, err
}

func main() {
	log.Printf("APIGW service starting on %s", listenAddress)

	// connect to post operations svc
	operationsSvcClient, err := newPostOperationsSvcClient()
	if err != nil {
		panic(err)
	}

	//lis, err := net.Listen("tcp", listenAddress)
	//if err != nil {
	//	log.Fatalf("failed to listen: %v", err)
	//}
	//s := grpc.NewServer(grpc.UnaryInterceptor(logger))

	//pb.RegisterPostServiceServer(s, NewPostOperationsService(operationsSvcClient))

	//if err := s.Serve(lis); err != nil {
	//	log.Fatalf("failed to serve: %v", err)
	//}
}
