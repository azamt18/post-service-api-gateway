package main

import (
	"bufio"
	"github.com/azamt18/post-service-grpc-api-gateway/services/post/inner/loader/grpcserver"
	post_loader_grpc_server "github.com/azamt18/post-service-grpc-api-gateway/services/post/inner/loader/internals/protobuff/post.loader.v1"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"strings"
)

const (
	postLoaderGrpcServerAddress = "POST_LOADER_GRPC_SERVER_ADDRESS"
	goRestApiHost               = "GOREST_API_HOST"
)

func initEnvVariables() {
	file, err := os.Open(".env")
	if err != nil {
		log.Fatalln(err)
		return
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		envKV := strings.SplitN(line, "=", 2)
		if err = os.Setenv(envKV[0], envKV[1]); err != nil {
			log.Fatalln(err)
		}
	}
	if err = file.Close(); err != nil {
		log.Fatalln(err)
	}
}

func init() {
	initEnvVariables()
}

func main() {
	grpcChan := make(chan int, 1)
	go startGRPCServer(grpcChan)

	<-grpcChan
}

func startGRPCServer(channel chan int) {
	// NewServer creates a gRPC server which has no service registered and has not started
	// to accept requests yet.
	s := grpc.NewServer()
	lis, err := net.Listen("tcp", os.Getenv(postLoaderGrpcServerAddress))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpcserver.NewGrpcServer(os.Getenv(goRestApiHost))
	post_loader_grpc_server.RegisterLoadPostsServiceServer(s, server)

	// Serve accepts incoming connections on the listener lis, creating a new ServerTransport
	// and service goroutine for each. The service goroutines read gRPC requests and then
	// call the registered handlers to reply to them.
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	channel <- 0
}
