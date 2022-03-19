package apiserver

import (
	"bufio"
	"github.com/azamt18/post-service-grpc-api-gateway/apiserver"
	"github.com/azamt18/post-service-grpc-api-gateway/db"
	posts_loader_service "github.com/azamt18/post-service-grpc-api-gateway/services/inner/posts-loader-service"
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

	postsLoaderServiceGrpcHost = "POSTS_LOADER_GRPC_SERVICE_HOST"
)

func main() {
	database := createDatabase()
	defer database.Disconnect()

	postsLoaderService := createPostsLoaderService()

	webApiChan := make(chan int, 1)
	go startWebApi(database, postsLoaderService, webApiChan)

	<-webApiChan
}

func startWebApi(database db.Database, postsLoaderService posts_loader_service.PostsLoaderService, channel chan int) {
	// init and start server
	server := apiserver.New(os.Getenv(apiServerBindAddr), database, postsLoaderService)
	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}

	channel <- 0
}

func createPostsLoaderService() posts_loader_service.PostsLoaderService {
	host := os.Getenv(postsLoaderServiceGrpcHost)

	return posts_loader_service.NewPostsLoaderService(host)
}

func createDatabase() db.Database {
	return db.NewDatabase(os.Getenv(mongoConnectionString), os.Getenv(mongoDatabaseName))
}
