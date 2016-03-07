package gotumblr

// FollowingResponse holds information about the blogs a user follows
type FollowingResponse struct {
	TotalBlogs int64
	Blogs      []FollowedBlog
}
