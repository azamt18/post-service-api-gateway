package controllers

import (
	request_models "github.com/azamt18/post-service-grpc-api-gateway/apiserver/request-models"
	"github.com/azamt18/post-service-grpc-api-gateway/apiserver/viewmodels"
	"github.com/azamt18/post-service-grpc-api-gateway/db"
	"github.com/azamt18/post-service-grpc-api-gateway/db/entity"
	posts_loader_service "github.com/azamt18/post-service-grpc-api-gateway/services/post/external/loader"
	post_operations_service "github.com/azamt18/post-service-grpc-api-gateway/services/post/external/operations"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type PostController interface {
	LoadPosts(ginContext *gin.Context)
	ListPosts(ginContext *gin.Context)
	ReadPost(ginContext *gin.Context)
	UpdatePost(ginContext *gin.Context)
	DeletePost(ginContext *gin.Context)
}

type postController struct {
	db                    db.Database
	postsLoaderService    posts_loader_service.PostsLoaderService
	postsOperationService post_operations_service.PostOperationsService
}

func (p *postController) LoadPosts(ginContext *gin.Context) {
	success, count, err := p.postsLoaderService.LoadPosts()

	response := viewmodels.LoadPostsViewModel{
		Success: success,
		Count:   count,
		Error:   err.Error(),
	}

	if err != nil {
		response.Success = false
		response.Count = 0
		response.Error = err.Error()

		ginContext.JSON(http.StatusInternalServerError, response)
		return
	}

	ginContext.JSON(http.StatusOK, response)

	return
}

func (p *postController) ListPosts(ginContext *gin.Context) {
	var (
		requestModel request_models.ListPostsRequestModel
		err          error = nil
	)

	response := viewmodels.ListPostsViewModel{}
	// bind request model
	{
		if err = ginContext.ShouldBindJSON(&requestModel); err != nil {
			ginContext.JSON(http.StatusInternalServerError, err)
			return
		}
	}

	res := p.postsOperationService.ListPosts(post_operations_service.ListPostsModel{
		Limit: requestModel.Limit,
		Skip:  requestModel.Skip,
	})

	if res.Error != nil {
		response.Success = false
		response.Error = response.Error.Error()
	}

	response.Success = true
	response.Post = res.Data

	ginContext.JSON(http.StatusOK, response)
	return
}

func (p *postController) ReadPost(ginContext *gin.Context) {
	response := viewmodels.ReadPostViewModel{}

	id, err := strconv.ParseInt(ginContext.Param("id"), 10, 64)
	if err != nil {
		response.Success = false
		response.Error = err.Error()

		ginContext.JSON(http.StatusBadRequest, response)
		return
	}

	res := p.postsOperationService.GetPost(id)
	if res.Error != nil {
		response.Success = false
		response.Error = err.Error()

		ginContext.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Success = true
	response.Post = res.Post

	ginContext.JSON(http.StatusOK, response)
	return
}

func (p *postController) UpdatePost(ginContext *gin.Context) {
	var (
		requestModel request_models.UpdatePostRequestModel
		err          error = nil
	)

	response := viewmodels.UpdatePostViewModel{}
	// bind request model
	{
		if err = ginContext.ShouldBindJSON(&requestModel); err != nil {
			ginContext.JSON(http.StatusInternalServerError, err)
			return
		}
	}

	id, err := strconv.ParseInt(ginContext.Param("id"), 10, 64)
	if err != nil {
		response.Success = false
		response.Error = err.Error()

		ginContext.JSON(http.StatusBadRequest, response)
		return
	}

	updatePostModel := &entity.Post{
		UserId: requestModel.UserId,
		Title:  requestModel.Title,
		Body:   requestModel.Body,
	}

	res := p.postsOperationService.UpdatePost(id, post_operations_service.UpdatePostModel{
		Post: updatePostModel,
	})

	if res.Error != nil {
		response.Success = false
		response.Error = response.Error.Error()
	}

	response.Success = true
	response.Post = res.Post

	ginContext.JSON(http.StatusOK, response)
	return
}

func (p *postController) DeletePost(ginContext *gin.Context) {
	response := viewmodels.DeletePostViewModel{}

	id, err := strconv.ParseInt(ginContext.Param("id"), 10, 64)
	if err != nil {
		response.Success = false
		response.Error = err.Error()

		ginContext.JSON(http.StatusBadRequest, response)
		return
	}

	res := p.postsOperationService.DeletePost(id)
	if res.Error != nil {
		response.Success = false
		response.Error = err.Error()

		ginContext.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Success = true

	ginContext.JSON(http.StatusOK, response)
	return
}

func NewPostController(db db.Database, postsLoaderService posts_loader_service.PostsLoaderService) PostController {
	return &postController{
		db:                 db,
		postsLoaderService: postsLoaderService,
	}
}
