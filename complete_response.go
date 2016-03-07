package gotumblr

import "encoding/json"

// CompleteResponse holds the entire response from the Tumblr API
type CompleteResponse struct {
	Meta     MetaInfo
	Response json.RawMessage
}
