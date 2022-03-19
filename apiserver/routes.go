package apiserver

import "github.com/azamt18/post-service-grpc-api-gateway/apiserver/controller"

func (apiServer *apiServer) registerRoutes() {
	checkPerformanceController := controller.NewCheckPerformanceController()
	postController := controller.NewPostController(apiServer.database, apiServer.postsLoaderService)

	apiServer.server.GET("/ping", checkPerformanceController.LoadPosts)

	group := apiServer.server.Group("/api")
	group.POST("/posts/load", postController.LoadPosts)
}
