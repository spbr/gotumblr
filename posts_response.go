package gotumblr

import "encoding/json"

// PostsResponse holds the information returned from a API call fetching posts
type PostsResponse struct {
	Blog       BlogInfo
	Posts      []json.RawMessage
	TotalPosts int64 `json:"total_posts"`
}
