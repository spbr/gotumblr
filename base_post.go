package gotumblr

//BasePost is the basic information common to all Tumblr posts
type BasePost struct {
	BlogName    string
	ID          int64
	PostURL     string
	PostType    string `json:"type"`
	Timestamp   int64
	Date        string
	Format      string
	ReblogKey   string
	Tags        []string
	Bookmarklet bool
	Mobile      bool
	SourceURL   string
	SourceTitle string
	Liked       bool
	State       string
	TotalPosts  int64
}
