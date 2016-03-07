package gotumblr

// OwnedBlog contains the information about a blog the user owns
type OwnedBlog struct {
	Name      string
	URL       string
	Title     string
	Primary   bool
	Followers int64
	Tweet     string
	Facebook  string
	Type      string
}
