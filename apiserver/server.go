package apiserver

import (
	"github.com/azamt18/post-service-grpc-api-gateway/apiserver/middlewares"
	"github.com/azamt18/post-service-grpc-api-gateway/db"
	postLoader "github.com/azamt18/post-service-grpc-api-gateway/services/post/external/loader"
	"github.com/gin-gonic/gin"
)

type apiServer struct {
	bindAddr           string
	server             *gin.Engine
	database           db.Database
	postsLoaderService postLoader.PostsLoaderService
}

func New(bindAddr string, database db.Database, postsLoaderService postLoader.PostsLoaderService) *apiServer {
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
