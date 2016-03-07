package gotumblr

//UserInfo holds the tumbler user data
type User struct {
	Name      string
	Following bool
	URL       string
	Updated   int64
}
