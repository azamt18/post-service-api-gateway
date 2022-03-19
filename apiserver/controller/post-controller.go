package controller

import (
	"github.com/azamt18/post-service-grpc-api-gateway/apiserver/viewmodels"
	"github.com/azamt18/post-service-grpc-api-gateway/db"
	posts_loader_service "github.com/azamt18/post-service-grpc-api-gateway/services/post/external/loader"
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
	success, count, err := p.postsLoaderService.LoadPosts()

	response := viewmodels.LoadPostsViewModel{
		Success: success,
		Count:   count,
		Error:   err,
	}

	if err != nil {
		context.JSON(http.StatusInternalServerError, response)
	}

	context.JSON(http.StatusOK, response)
}

func NewPostController(db db.Database, postsLoaderService posts_loader_service.PostsLoaderService) PostController {
	return &postController{
		db:                 db,
		postsLoaderService: postsLoaderService,
	}
}
