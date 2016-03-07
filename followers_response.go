package gotumblr

// FollowersResponse holds information about the users that follow a Tumblr blog
type FollowersResponse struct {
	TotalUsers int64
	Users      []User
}
