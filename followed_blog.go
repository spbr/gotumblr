package gotumblr

// FollowedBlog holds information about a blog that the user follows
type FollowedBlog struct {
	Name        string
	URL         string
	Updated     int64
	Title       string
	Description string
}
