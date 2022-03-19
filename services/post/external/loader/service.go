package loader

import (
	"context"
	post_loader_grpc_client "github.com/azamt18/post-service-grpc-api-gateway/services/post/inner/loader/internals/protobuff/post.loader.v1"
)

type PostsLoaderService interface {
	LoadPosts() (bool, int64, error)
}

type postsLoaderService struct {
	postLoaderGrpcClient post_loader_grpc_client.LoadPostsServiceClient
}

func (p postsLoaderService) LoadPosts() (bool, int64, error) {

	res, err := p.postLoaderGrpcClient.LoadPosts(context.TODO(), &post_loader_grpc_client.LoadPostsRequest{
		PageNumber: 1, //todo remove page number -> make loading all
	})
	if err != nil {
		return false, 0, err
	}

	return res.GetSuccess(), res.GetLoadedPostsCount(), nil
}

func NewPostsLoaderService(postLoaderGrpcClient post_loader_grpc_client.LoadPostsServiceClient) PostsLoaderService {
	return &postsLoaderService{
		postLoaderGrpcClient: postLoaderGrpcClient,
	}
}
