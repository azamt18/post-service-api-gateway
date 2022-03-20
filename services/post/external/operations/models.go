package operations

import "github.com/azamt18/post-service-grpc-api-gateway/db/entity"

type ReadPostModel struct{}

type ReadPostResult struct {
	Success bool
	Error   error
	Post    *entity.Post
}

type ListPostsModel struct {
	Limit int64
	Skip  int64
}

type ListPostsResult struct {
	Success bool
	Error   error
	Data    []*entity.Post
}

type UpdatePostModel struct {
	Post *entity.Post
}

type UpdatePostResult struct {
	Success bool
	Error   error
	Post    *entity.Post
}

type DeletePostModel struct{}

type DeletePostResult struct {
	Success bool
	PostId  int64
	Error   error
}
