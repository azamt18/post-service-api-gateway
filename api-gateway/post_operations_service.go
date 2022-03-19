package main

import (
	"context"
	"fmt"
	postOperationsSvcV1 "github.com/azamt18/post-service-grpc-api-gateway/services/post/operations/api/goclient/v1"
	"io"
	"log"

	pb "github.com/azamt18/post-service-grpc-api-gateway/api-gateway/api/goclient/v1"
)

type postOperationsService struct {
	postOperationsClient postOperationsSvcV1.PostServiceClient
}

func NewPostOperationsService(postOperationsClient postOperationsSvcV1.PostServiceClient) *postOperationsService {
	return &postOperationsService{
		postOperationsClient: postOperationsClient,
	}
}

func (u *postOperationsService) ReadPost(ctx context.Context, request *pb.ReadPostRequest) (*pb.ReadPostResponse, error) {
	req := &pb.ReadPostRequest{PostId: request.GetPostId()}
	res, err := u.postOperationsClient.ReadPost(context.Background(), req)
	if err != nil {
		fmt.Printf("Error happened while reading: %v\n", err)
		return nil, err
	}

	post := &pb.Post{
		Id:     res.GetPost().GetId(),
		UserId: res.GetPost().GetUserId(),
		Title:  res.GetPost().GetTitle(),
		Body:   res.GetPost().GetBody(),
	}

	// in this case the messages are quite similar (for now) but we have to translate
	// them between API structs and internal structs
	return &pb.ReadPostResponse{
		Post: post,
	}, nil
}

func (u *postOperationsService) UpdatePost(ctx context.Context, request *pb.UpdatePostRequest) (*pb.UpdatePostResponse, error) {
	newPost := &pb.Post{
		Id:         request.GetPost().GetId(),
		UserId: 	request.GetPost().GetUserId(),
		Title:      request.GetPost().GetTitle(),
		Body:  		request.GetPost().GetBody(),
	}

	res, err := u.postOperationsClient.UpdatePost(context.Background(), &pb.UpdatePostRequest{
		Post: newPost,
	})

	if err != nil {
		fmt.Printf("Error happened while updating: %v\n", err)
		return nil, err
	}

	return &pb.UpdatePostResponse{
		Post: newPost,
	}, nil
}

func (u *postOperationsService) DeletePost(ctx context.Context, request *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
	res, err := u.postOperationsClient.DeletePost(context.Background(), &pb.DeletePostRequest{
		PostId: request.GetPostId()
	})

	if err != nil {
		fmt.Printf("Error while deleting: %v", err)
		return nil, err
	}

	fmt.Printf("Blog was deleted: %v", res)
	return &pb.DeletePostResponse{
		PostId: request.GetPostId(),
	}, nil
}

func (u *postOperationsService) ListPosts(ctx context.Context, request *pb.ListPostsRequest) (*pb.ListPostsResponse, error) {
	stream, err := u.postOperationsClient.ListPosts(context.Background(), &pb.ListPostsRequest{})
	if err != nil {
		log.Fatalf("Error while calling ListBlog: %v", err)
	}
	for {
		res, err := stream.Recv()

		// break when a stream riches to the end
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error in receiving a stream: %v", err)
			return nil, err
		}

		fmt.Printf("ListBlog response: %v", res.GetPost())
		return &pb.ListPostsResponse{
			Post: res.GetPost(),
		}, nil
	}

}