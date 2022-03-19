package models

type LoadPostsResponseViewModel struct {
	Success          bool   `json:"success,omitempty"`
	LoadedPostsCount int64  `json:"loaded_posts_count,omitempty"`
	Error            string `json:"error,omitempty"`
}
