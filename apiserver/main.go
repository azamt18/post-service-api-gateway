package apiserver

import (
	db "github.com/azamt18/post-service-grpc-api-gateway/db"
	posts_loader_service "github.com/azamt18/post-service-grpc-api-gateway/services/inner/posts-loader-service"
)

const (
	webApiServerAddress        = "WEB_API_SERVER_SERVER"
	mongoConnectionString      = "MONGO_CONNECTION_STRING"
	mongoDbName                = "MONGO_DATABASE_NAME"
	postsLoaderServiceGrpcHost = "POSTS_LOADER_GRPC_SERVICE_HOST"
)

func main() {
	database := db.NewDatabase(mongoConnectionString, mongoDbName)
	postLoaderService := posts_loader_service.NewPostsLoaderService(postsLoaderServiceGrpcHost)

	webApiServer := New(webApiServerAddress, database, postLoaderService)
	webApiServer.Start()
}
