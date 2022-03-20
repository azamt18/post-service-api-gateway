package operations

import (
	"context"
	"github.com/azamt18/post-service-grpc-api-gateway/db/entity"
	post_operations_grpc_client "github.com/azamt18/post-service-grpc-api-gateway/services/post/inner/operations/internals/protobuff/post.operations.v1"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostOperationsService interface {
	GetPost(model ReadPostModel) *ReadPostResult
}

type postOperationsService struct {
	postOperationsGrpcClient post_operations_grpc_client.PostOperationsServiceClient
}

func (p postOperationsService) GetPost(model ReadPostModel) (result *ReadPostResult) {
	result = &ReadPostResult{}

	var post *entity.Post

	res, err := p.postOperationsGrpcClient.ReadPost(context.TODO(), &post_operations_grpc_client.ReadPostRequest{
		PostId: model.Id,
	})

	if err != nil {
		return
	}

	postId := res.GetPost().GetId()
	oid, error := primitive.ObjectIDFromHex(string(postId))
	if error != nil {
		result.Success = false
		result.Error = err
		return
	}

	post = &entity.Post{
		Id:     oid,
		UserId: int(res.GetPost().GetUserId()),
		Title:  res.GetPost().GetTitle(),
		Body:   res.GetPost().GetBody(),
	}

	result.Success = true
	result.Post = post

	return
}

func NewPostsLoaderService(postOperationsGrpcClient post_operations_grpc_client.PostOperationsServiceClient) PostOperationsService {
	return &postOperationsService{
		postOperationsGrpcClient: postOperationsGrpcClient,
	}
}
