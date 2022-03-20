package apiserver

import (
	"github.com/azamt18/post-service-grpc-api-gateway/apiserver/middlewares"
	"github.com/azamt18/post-service-grpc-api-gateway/db"
	postLoaderService "github.com/azamt18/post-service-grpc-api-gateway/services/post/external/loader"
	postOperationsService "github.com/azamt18/post-service-grpc-api-gateway/services/post/external/operations"
	"github.com/gin-gonic/gin"
)

type apiServer struct {
	bindAddr              string
	server                *gin.Engine
	database              db.Database
	postsLoaderService    postLoaderService.PostsLoaderService
	postOperationsService postOperationsService.PostOperationsService
}

func New(bindAddr string, database db.Database, postsLoaderService postLoaderService.PostsLoaderService, postOperationsService postOperationsService.PostOperationsService) *apiServer {
	apiServer := &apiServer{
		bindAddr:              bindAddr,
		server:                gin.Default(),
		database:              database,
		postsLoaderService:    postsLoaderService,
		postOperationsService: postOperationsService,
	}

	apiServer.server.Use(middlewares.Cors())
	apiServer.registerRoutes()

	return apiServer
}

func (apiServer *apiServer) Start() error {
	return apiServer.server.Run(apiServer.bindAddr)
}
