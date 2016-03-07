package gotumblr

import "encoding/json"

// DraftsResponse holds all the draft posts
type DraftsResponse struct {
	Posts []json.RawMessage
}
