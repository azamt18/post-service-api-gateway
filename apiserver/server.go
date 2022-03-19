package apiserver

import (
	"github.com/azamt18/post-service-grpc-api-gateway/apiserver/middlewares"
	"github.com/azamt18/post-service-grpc-api-gateway/db"
	posts_loader_service "github.com/azamt18/post-service-grpc-api-gateway/services/inner/posts-loader-service"
	"github.com/gin-gonic/gin"
)

type apiServer struct {
	bindAddr           string
	server             *gin.Engine
	database           db.Database
	postsLoaderService posts_loader_service.PostsLoaderService
}

func New(bindAddr string, database db.Database, postsLoaderService posts_loader_service.PostsLoaderService) *apiServer {
	apiServer := &apiServer{
		bindAddr:           bindAddr,
		server:             gin.Default(),
		database:           database,
		postsLoaderService: postsLoaderService,
	}

	apiServer.server.Use(middlewares.Cors())
	apiServer.registerRoutes()

	return apiServer
}

func (apiServer *apiServer) Start() error {
	return apiServer.server.Run(apiServer.bindAddr)
}
