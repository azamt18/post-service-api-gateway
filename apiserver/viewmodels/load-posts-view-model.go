package viewmodels

import "github.com/azamt18/post-service-grpc-api-gateway/db/entity"

type LoadPostsViewModel struct {
	Success bool   `json:"success"`
	Count   int64  `json:"count"`
	Error   string `json:"error"`
}

type ListPostsViewModel struct {
	Success bool
	Post    []*entity.Post
	Error   string
}

type ReadPostViewModel struct {
	Success bool
	Post    *entity.Post
	Error   string
}

type UpdatePostViewModel struct {
	Success bool
	Post    *entity.Post
	Error   string
}

type DeletePostViewModel struct {
	Success bool
	Error   string
}
