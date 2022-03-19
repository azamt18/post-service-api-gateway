package main

import (
	"bufio"
	"github.com/azamt18/post-service-grpc-api-gateway/apiserver"
	"github.com/azamt18/post-service-grpc-api-gateway/db"
	postLoader "github.com/azamt18/post-service-grpc-api-gateway/services/post/external/loader"
	postLoaderGrpcClient "github.com/azamt18/post-service-grpc-api-gateway/services/post/inner/loader/client"
	"log"
	"os"
	"strings"
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

const (
	apiServerBindAddr = "API_SERVER_BIND_ADDRESS"

	mongoConnectionString = "MONGO_CONNECTION_STRING"
	mongoDatabaseName     = "MONGO_DATABASE_NAME"
)

func main() {
	database := createDatabase()
	defer database.Disconnect()

	postLoaderGrpcClient := postLoaderGrpcClient.NewClient()
	postLoaderService := postLoader.NewPostsLoaderService(postLoaderGrpcClient)

	webApiChan := make(chan int, 1)
	go startWebApi(database, postLoaderService, webApiChan)

	<-webApiChan
}

func startWebApi(database db.Database, postsLoaderService postLoader.PostsLoaderService, channel chan int) {
	// init and start server
	server := apiserver.New(os.Getenv(apiServerBindAddr), database, postsLoaderService)
	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}

	channel <- 0
}

func createDatabase() db.Database {
	return db.NewDatabase(os.Getenv(mongoConnectionString), os.Getenv(mongoDatabaseName))
}
