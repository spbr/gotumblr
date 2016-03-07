package gotumblr

// UserInfo holds the Tumbler userinfo data
type UserInfo struct {
	Following         int64
	DefaultPostFormat string
	Name              string
	Likes             int64
	Blogs             []OwnedBlog
}
