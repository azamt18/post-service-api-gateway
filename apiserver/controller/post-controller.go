package controller

import (
	"github.com/azamt18/post-service-grpc-api-gateway/db"
	posts_loader_service "github.com/azamt18/post-service-grpc-api-gateway/services/inner/posts-loader-service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PostController interface {
	LoadPosts(context *gin.Context)
}

type postController struct {
	db                 db.Database
	postsLoaderService posts_loader_service.PostsLoaderService
}

func (p *postController) LoadPosts(context *gin.Context) {
	response, _ := p.postsLoaderService.LoadPosts()

	context.JSON(http.StatusOK, response.LoadedPostsCount)
}

func NewPostController(db db.Database, postsLoaderService posts_loader_service.PostsLoaderService) PostController {
	return &postController{
		db:                 db,
		postsLoaderService: postsLoaderService,
	}
}
