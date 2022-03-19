package posts_loader_service

import (
	"github.com/azamt18/post-service-grpc-api-gateway/services/inner/posts-loader-service/models"
)

type PostsLoaderService interface {
	LoadPosts() (*models.LoadPostsResponseViewModel, error)
}

type postsLoaderService struct {
	host string
}

func (p postsLoaderService) LoadPosts() (*models.LoadPostsResponseViewModel, error) {
	//TODO implement me
	panic("implement me")
}

func NewPostsLoaderService(host string) PostsLoaderService {
	return &postsLoaderService{
		host: host,
	}
}
