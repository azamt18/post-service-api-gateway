package operations

import (
	"context"
	"github.com/azamt18/post-service-grpc-api-gateway/db/entity"
	post_operations_grpc_client "github.com/azamt18/post-service-grpc-api-gateway/services/post/inner/operations/internals/protobuff/post.operations.v1"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
)

type PostOperationsService interface {
	GetPost(id int64) *ReadPostResult
	ListPosts(model ListPostsModel) (result *ListPostsResult)
	UpdatePost(id int64, model UpdatePostModel) (result *UpdatePostResult)
	DeletePost(id int64) (result *DeletePostResult)
}

type postOperationsService struct {
	postOperationsGrpcClient post_operations_grpc_client.PostOperationsServiceClient
}

func (p postOperationsService) GetPost(id int64) (result *ReadPostResult) {
	result = &ReadPostResult{}

	var post *entity.Post

	res, err := p.postOperationsGrpcClient.ReadPost(context.TODO(), &post_operations_grpc_client.ReadPostRequest{
		PostId: id,
	})

	if err != nil {
		result.Success = false
		result.Error = err
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
		UserId: res.GetPost().GetUserId(),
		Title:  res.GetPost().GetTitle(),
		Body:   res.GetPost().GetBody(),
	}

	result.Success = true
	result.Post = post

	return
}

func (p postOperationsService) ListPosts(model ListPostsModel) (result *ListPostsResult) {
	result = &ListPostsResult{}

	stream, err := p.postOperationsGrpcClient.ListPosts(context.TODO(), &post_operations_grpc_client.ListPostsRequest{
		Limit: model.Limit,
		Skip:  model.Skip,
	})

	if err != nil {
		result.Success = false
		result.Error = err
		return
	}

	var posts []*entity.Post

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return
		}

		var post *entity.Post
		postId := res.GetPost().GetId()
		oid, error := primitive.ObjectIDFromHex(string(postId))
		if error != nil {
			result.Success = false
			result.Error = err
			return
		}

		post = &entity.Post{
			Id:     oid,
			UserId: res.GetPost().GetUserId(),
			Title:  res.GetPost().GetTitle(),
			Body:   res.GetPost().GetBody(),
		}

		posts = append(posts, post)
	}

	result.Success = true
	result.Data = posts
	return
}

func (p postOperationsService) UpdatePost(id int64, model UpdatePostModel) (result *UpdatePostResult) {
	result = &UpdatePostResult{}

	postModel := &post_operations_grpc_client.Post{
		UserId: model.Post.UserId,
		Title:  model.Post.Title,
		Body:   model.Post.Body,
	}

	res, err := p.postOperationsGrpcClient.UpdatePost(context.TODO(), &post_operations_grpc_client.UpdatePostRequest{
		PostId: id,
		Post:   postModel,
	})

	if err != nil {
		result.Success = false
		result.Error = err
		return
	}

	postId := res.GetPost().GetId()
	oid, error := primitive.ObjectIDFromHex(string(postId))
	if error != nil {
		result.Success = false
		result.Error = err
		return
	}

	post := &entity.Post{
		Id:     oid,
		UserId: res.GetPost().GetUserId(),
		Title:  res.GetPost().GetTitle(),
		Body:   res.GetPost().GetBody(),
	}

	result.Success = true
	result.Post = post

	return
}

func (p postOperationsService) DeletePost(id int64) (result *DeletePostResult) {
	result = &DeletePostResult{}

	res, err := p.postOperationsGrpcClient.DeletePost(context.TODO(), &post_operations_grpc_client.DeletePostRequest{
		PostId: id,
	})

	if err != nil {
		result.Success = false
		result.Error = err
		return
	}

	result.Success = true
	result.PostId = res.GetPostId()

	return
}

func NewPostOperationsService(postOperationsGrpcClient post_operations_grpc_client.PostOperationsServiceClient) PostOperationsService {
	return &postOperationsService{
		postOperationsGrpcClient: postOperationsGrpcClient,
	}
}
