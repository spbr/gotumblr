package gotumblr

// BlogInfo holds information about a Tumblr blog
type BlogInfo struct {
	Title       string
	Posts       int64
	Name        string
	URL         string
	Updated     int64
	Description string
	Ask         bool
	AskAnon     bool
	Likes       int64
}
