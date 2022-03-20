package operations

import "github.com/azamt18/post-service-grpc-api-gateway/db/entity"

type ReadPostModel struct {
	Id string
}

type ReadPostResult struct {
	Success bool
	Error   error
	Post    *entity.Post
}
