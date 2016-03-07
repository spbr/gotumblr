package gotumblr

// TextPost holds the data for a Tumblr text post
type TextPost struct {
	BasePost
	Title, Body string
}

// AltSize holds the data for an images alternative size
type AltSize struct {
	Width, Height int64
	URL           string
}

// PhotoObject holds information about a photo
type PhotoObject struct {
	Caption  string
	AltSizes []AltSize
}

// PhotoPost holds the information for a Tumblr photo post
type PhotoPost struct {
	BasePost
	Photos        []PhotoObject
	Caption       string
	Width, Height int64
}

// QuotePost holds the information for a Tumblr quote post
type QuotePost struct {
	BasePost
	Text, Source string
}

// LinkPost holds the information for a Tumblr link post
type LinkPost struct {
	BasePost
	Title, URL, Description string
}

// DialogueInfo holds the dialog information from a chat item
type DialogueInfo struct {
	Name, Label, Phrase string
}

// ChatPost holds the information for a Tumblr chat post
type ChatPost struct {
	BasePost
	Title, Body string
	Dialogue    []DialogueInfo
}

// AudioPost holds the information for a Tumblr audio post
type AudioPost struct {
	BasePost
	Caption     string
	Player      string
	Plays       int64
	AlbumArt    string
	Artist      string
	Album       string
	TrackName   string
	TrackNumber int64
	Year        int64
}

// PlayerInfo holds infromation about a video player
type PlayerInfo struct {
	Width     int64
	EmbedCode string
}

// VideoPost contains the information for a Tumblr video post
type VideoPost struct {
	BasePost
	Caption string
	Player  []PlayerInfo
}

// AnswerPost contains the information for a Tumblr answer post
type AnswerPost struct {
	BasePost
	AskingName string
	AskingURL  string
	Question   string
	Answer     string
}
