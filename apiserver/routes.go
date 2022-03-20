package apiserver

import "github.com/azamt18/post-service-grpc-api-gateway/apiserver/controllers"

func (apiServer *apiServer) registerRoutes() {
	checkPerformanceController := controllers.NewCheckPerformanceController()
	postController := controllers.NewPostController(apiServer.database, apiServer.postsLoaderService)

	apiServer.server.GET("/ping", checkPerformanceController.LoadPosts)

	group := apiServer.server.Group("/api")
	group.POST("/post/load", postController.LoadPosts)

	group.GET("/post", postController.ListPosts)
	group.GET("/post/:id", postController.ReadPost)
	group.PUT("/post", postController.UpdatePost)
	group.DELETE("/post/:id", postController.DeletePost)
}
