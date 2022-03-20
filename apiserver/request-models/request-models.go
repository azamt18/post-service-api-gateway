package request_models

type ListPostsRequestModel struct {
	Limit int64 `json:"limit"`
	Skip  int64 `json:"skip"`
}

type UpdatePostRequestModel struct {
	UserId int64  `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
