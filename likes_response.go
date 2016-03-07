package gotumblr

import "encoding/json"

// LikesResponse holds information about the posts a user liked on Tumblr
type LikesResponse struct {
	LikedPosts []json.RawMessage
	LikedCount int64
}
