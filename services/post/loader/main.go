package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/azamt18/post-service-grpc-api-gateway/db"
	pb "github.com/azamt18/post-service-grpc-api-gateway/services/post/loader/api/goclient/v1"
	"github.com/azamt18/post-service-grpc-api-gateway/services/post/loader/api/v1/handlers"
	"github.com/azamt18/post-service-grpc-api-gateway/services/post/loader/api/v1/services/gorest-api-service"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
)

const (
	listenAddress     = "0.0.0.0:50051"
	dbHost            = "mongodb://localhost:27017"
	dbName            = "mydb"
	collectionName    = "posts"
	goRestApiHost     = "https://gorest.co.in"
	grpcServerAddress = "POSTS_LOADER_GRPC_SERVER_ADDRESS"
)

var collection *mongo.Collection

type postLoaderService struct{}

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

func (p postLoaderService) LoadPosts(ctx context.Context, request *pb.LoadPostsRequest) (*pb.LoadPostsResponse, error) {
	//todo implement me
	panic("implement me")
}

func main() {
	// if the go code is crushed -> get the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	database := db.NewDatabase(dbHost, dbName)
	defer database.Disconnect()

	goRestApiService := initializeGoRestApiService()

	webApiChan := make(chan int, 1)
	go startWebApi(database, goRestApiService, webApiChan)
	<-webApiChan

	// initialize connection to db
	collection = database.PostsCollection()

	// starting a service (for internal connections)
	fmt.Println("Start Posts Loading Service...")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Panicf("FAILED TO LISTEN %v", err)
	}
	fmt.Printf("Posts Loading is listening on: %v", lis.Addr())

	// starting an API server (for public connections)
	if startApiServer(collection); err != nil {
		log.Panicf("FAILED TO START API SERVER: %v", err)
	}

	// closing the connection to db is not necessary
	// because of it is regulated
	defer database.Disconnect()
}

func startWebApi(database db.Database, goRestApiService gorest_api_service.GoRestApiService, channel chan int) {
	// init and start server
	fmt.Println("Start Posts Loading Service...")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("FAILED TO LISTEN %v", err)
	}

	var options []grpc.ServerOption
	s := grpc.NewServer(options...)
	pb.RegisterLoadPostsServiceServer(s, &postLoaderService{})

	// Register reflection service on gRPC server
	reflection.Register(s)

	go func() {
		fmt.Println("Starting Blog Server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Wait for Control C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block until a signal is received
	<-ch
}

func startApiServer(collection *mongo.Collection) error {
	postsController := handlers.NewController(collection)
	router := mux.NewRouter()

	router.HandleFunc("/posts/load", postsController.LoadPosts).Methods(http.MethodGet)

	log.Println("API is running")
	return http.ListenAndServe(":4000", router)
}

func initializeGoRestApiService() gorest_api_service.GoRestApiService {
	host := os.Getenv(goRestApiHost)

	return gorest_api_service.NewGoRestApiService(host)
}
