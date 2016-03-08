//A Go Tumblr API v2 Client.

package gotumblr

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

//TumblrRestClient defines a Go Client for the Tumblr API.
type TumblrRestClient struct {
	request *TumblrRequest
}

//NewTumblrRestClient initializes the TumblrRestClient, creating TumblrRequest that deals with all request formatting.
//consumerKey is the consumer key of your Tumblr Application.
//consumerSecret is the consumer secret of your Tumblr Application.
//oauthToken is the user specific token, received from the /access_token endpoint.
//oauthSecret is the user specific secret, received from the /access_token endpoint.
//host is the host that you are tryng to send information to (e.g. http://api.tumblr.com).
func NewTumblrRestClient(consumerKey, consumerSecret, oauthToken, oauthSecret, callbackURL, host string) *TumblrRestClient {
	return &TumblrRestClient{NewTumblrRequest(consumerKey, consumerSecret, oauthToken, oauthSecret, callbackURL, host)}
}

//Info retrieves the user information.
func (trc *TumblrRestClient) Info() (*UserInfoResponse, error) {
	data, err := trc.request.Get("/v2/user/info", map[string]string{})
	if err != nil {
		return nil, err
	}
	var result UserInfoResponse
	json.Unmarshal(data.Response, &result)
	return &result, nil
}

//Avatar etrieves the url of the blog's avatar.
//size can be: 16, 24, 30, 40, 48, 64, 96, 128 or 512.
func (trc *TumblrRestClient) Avatar(blogname string, size int) (*AvatarResponse, error) {
	requestURL := trc.request.host + fmt.Sprintf("/v2/blog/%s/avatar/%d", blogname, size)
	httpRequest, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return nil, err
	}
	transport := &http.Transport{}
	httpResponse, err := transport.RoundTrip(httpRequest)
	if err != nil {
		return nil, err
	}
	defer httpResponse.Body.Close()
	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}
	data, err := trc.request.JSONParse(body)
	if err != nil {
		return nil, err
	}
	var result AvatarResponse
	json.Unmarshal(data.Response, &result)
	return &result, nil
}

//Likes retrieves the likes of the given user.
//options can be:
//limit: the number of results to return, inclusive;
//offset: liked post number to start at.
func (trc *TumblrRestClient) Likes(options map[string]string) (*LikesResponse, error) {
	data, err := trc.request.Get("/v2/user/likes", options)
	if err != nil {
		return nil, err
	}
	var result LikesResponse
	json.Unmarshal(data.Response, &result)
	return &result, nil
}

//Following retrieves the blogs that the user is following.
//options can be:
//limit: the number of results to return;
//offset: result number to start at.
func (trc *TumblrRestClient) Following(options map[string]string) (*FollowingResponse, error) {
	data, err := trc.request.Get("/v2/user/following", options)
	if err != nil {
		return nil, err
	}
	var result FollowingResponse
	json.Unmarshal(data.Response, &result)
	return &result, nil
}

//Dashboard retrieve the dashboard of the user.
//options can be:
//limit: number of results to return;
//offset: post number to start at;
//type: the type of posts to return(text, photo, quote, link, chat, audio, video, answer);
//since_id: return posts that have apeared after this id;
//reblog_info: whether to return reblog information about the posts;
//notes_info: whether to return notes information about the posts.
func (trc *TumblrRestClient) Dashboard(options map[string]string) (*DraftsResponse, error) {
	data, err := trc.request.Get("/v2/user/dashboard", options)
	if err != nil {
		return nil, err
	}
	var result DraftsResponse
	json.Unmarshal(data.Response, &result)
	return &result, nil
}

//Tagged retrieves a list of posts with the given tag.
//tag: the tag you want to look for.
//options can be:
//before: the timestamp of when you'd like to see posts before;
//limit: the number of results to return;
//filter: the post format you want to get(e.g html, text, raw).
func (trc *TumblrRestClient) Tagged(tag string, options map[string]string) ([]json.RawMessage, error) {
	options["tag"] = tag
	options["api_key"] = trc.request.apiKey
	data, err := trc.request.Get("/v2/tagged", options)
	if err != nil {
		return nil, err
	}
	result := []json.RawMessage{}
	json.Unmarshal(data.Response, &result)
	return result, nil
}

//Posts retrieves a list of posts from a blog.
//blogname: the name of the blog you want to get posts from (e.g. mgterzieva.tumblr.com).
//postsType: the type of the posts you want to get
//(e.g. text, quote, link, answer, video, audio, photo, chat, all).
//options can be:
//id: the id of the post you are looking for;
//tag: return only posts with this tag;
//limit: the number of posts to return;
//offset: the number of the post you want to start from;
//filter: return only posts with a specific format(e.g. html, text, raw).
func (trc *TumblrRestClient) Posts(blogname, postsType string, options map[string]string) (*PostsResponse, error) {
	var requestURL string
	if postsType == "" {
		requestURL = fmt.Sprintf("/v2/blog/%s/posts", blogname)
	} else {
		requestURL = fmt.Sprintf("/v2/blog/%s/posts/%s", blogname, postsType)
	}
	options["api_key"] = trc.request.apiKey
	data, err := trc.request.Get(requestURL, options)
	if err != nil {
		return nil, err
	}
	var result PostsResponse
	json.Unmarshal(data.Response, &result)
	return &result, nil
}

//BlogInfo retrieves general information about the blog.
//blogname: name of the blog you want to get information about(e.g. mgterzieva.tumblr.com).
func (trc *TumblrRestClient) BlogInfo(blogname string) (*BlogInfoResponse, error) {
	requestURL := fmt.Sprintf("/v2/blog/%s/info", blogname)
	options := map[string]string{"api_key": trc.request.apiKey}
	data, err := trc.request.Get(requestURL, options)
	if err != nil {
		return nil, err
	}
	var result BlogInfoResponse
	json.Unmarshal(data.Response, &result)
	return &result, nil
}

//Followers retrieves the followers of the blog given.
//blogname: name of the blog whose followers you want to get.
//optons can be:
//limit: the number of results to return, inclusive;
//offset: result to start at.
func (trc *TumblrRestClient) Followers(blogname string, options map[string]string) (*FollowersResponse, error) {
	requestURL := fmt.Sprintf("/v2/blog/%s/followers", blogname)
	data, err := trc.request.Get(requestURL, options)
	if err != nil {
		return nil, err
	}
	var result FollowersResponse
	json.Unmarshal(data.Response, &result)
	return &result, nil
}

//BlogLikes retrieves the likes of blog given.
//blogname: name of the blog whose likes you want to get.
//options can be:
//limit: how many likes do you want to get;
//offset: the number of the like you want to start from.
func (trc *TumblrRestClient) BlogLikes(blogname string, options map[string]string) (*LikesResponse, error) {
	requestURL := fmt.Sprintf("/v2/blog/%s/likes", blogname)
	options["api_key"] = trc.request.apiKey
	data, err := trc.request.Get(requestURL, options)
	if err != nil {
		return nil, err
	}
	var result LikesResponse
	json.Unmarshal(data.Response, &result)
	return &result, nil
}

//Queue retrieves posts that are currently in the blog's queue.
//options can be:
//limit: the number of results to return;
//offset: post number to start at;
//filter: specify posts' format(e.g. format="html", format="text", format="raw").
func (trc *TumblrRestClient) Queue(blogname string, options map[string]string) (*DraftsResponse, error) {
	requestURL := fmt.Sprintf("/v2/blog/%s/posts/queue", blogname)
	data, err := trc.request.Get(requestURL, options)
	if err != nil {
		return nil, err
	}
	var result DraftsResponse
	json.Unmarshal(data.Response, &result)
	return &result, nil
}

//Drafts retrieves posts that are currently in the blog's drafts.
//options can be:
//filter: specify posts' format(e.g. format="html", format="text", format="raw").
func (trc *TumblrRestClient) Drafts(blogname string, options map[string]string) (*DraftsResponse, error) {
	requestURL := fmt.Sprintf("/v2/blog/%s/posts/draft", blogname)
	data, err := trc.request.Get(requestURL, options)
	if err != nil {
		return nil, err
	}
	var result DraftsResponse
	json.Unmarshal(data.Response, &result)
	return &result, nil
}

//Submission retrieves submission posts.
//options can be:
//offset: post number to start at;
//filter: specify posts' format(e.g. format="html", format="text", format="raw").
func (trc *TumblrRestClient) Submission(blogname string, options map[string]string) (*DraftsResponse, error) {
	requestURL := fmt.Sprintf("/v2/blog/%s/posts/submission", blogname)
	data, err := trc.request.Get(requestURL, options)
	if err != nil {
		return nil, err
	}
	var result DraftsResponse
	json.Unmarshal(data.Response, &result)
	return &result, nil
}

//Follow a blog via a URL
//blogname: the url of the blog to follow.
func (trc *TumblrRestClient) Follow(blogname string) (bool, error) {
	requestURL := fmt.Sprintf("/v2/user/follow")
	params := map[string]string{"url": blogname}
	data, err := trc.request.Post(requestURL, params)
	if err != nil {
		return false, err
	}
	if data.Meta.Status != 200 {
		return false, errors.New(data.Meta.Msg)
	}
	return true, nil
}

//Unfollow a blog via a URL
//blogname: the url of the blog to unfollow.
func (trc *TumblrRestClient) Unfollow(blogname string) (bool, error) {
	requestURL := fmt.Sprintf("/v2/user/unfollow")
	params := map[string]string{"url": blogname}
	data, err := trc.request.Post(requestURL, params)
	if err != nil {
		return false, err
	}
	if data.Meta.Status != 200 {
		return false, errors.New(data.Meta.Msg)
	}
	return true, nil
}

//Like post of a given blog.
//id: the id of the post you want to like.
//reblog_key: the reblog key for the post id.
func (trc *TumblrRestClient) Like(id, reblogKey string) error {
	requestUrl := fmt.Sprintf("/v2/user/like")
	params := map[string]string{"id": id, "reblog_key": reblogKey}
	data := trc.request.Post(requestUrl, params)
	if data.Meta.Status != 200 {
		return errors.New(data.Meta.Msg)
	}
	return nil
}

//Unlike a post of a given blog.
//id: the id of the post you want to unlike.
//reblog_key: the reblog key for the post id.
func (trc *TumblrRestClient) Unlike(id, reblogKey string) error {
	requestUrl := fmt.Sprintf("/v2/user/unlike")
	params := map[string]string{"id": id, "reblog_key": reblogKey}
	data := trc.request.Post(requestUrl, params)
	if data.Meta.Status != 200 {
		return errors.New(data.Meta.Msg)
	}
	return nil
}

//Create a photo post or photoset on a blog.
//blogname: the url of the blog you want to post to.
//options can be:
//(with * are marked required options)
//state: the state of the post(e.g. published, draft, queue, private);
//tags: a list of tags you want applied to the post;
//tweet: manages the autotweet for this post: set to off for no tweet
//or enter text to override the default tweet;
//date: the GMT date and time of the post as a string;
//format: sets the format type of the post(html or markdown);
//slug: add a short text summary to the end of the post url;
//caption: the caption that you want applied to the photo;
//link: the 'click-through' url for the photo;
//*source: the photo source url.
func (trc *TumblrRestClient) CreatePhoto(blogname string, options map[string]string) error {
	requestUrl := fmt.Sprintf("/v2/blog/%s/post", blogname)
	options["type"] = "photo"
	data := trc.request.Post(requestUrl, options)
	if data.Meta.Status != 201 {
		return errors.New(data.Meta.Msg)
	}
	return nil
}

//Create a text post on a blog.
//blogname: the url of the blog you want to post to.
//options can be:
//(with * are marked required options)
//state: the state of the post(e.g. published, draft, queue, private);
//tags: a list of tags you want applied to the post;
//tweet: manages the autotweet for this post: set to off for no tweet
//or enter text to override the default tweet;
//date: the GMT date and time of the post as a string;
//format: sets the format type of the post(html or markdown);
//slug: add a short text summary to the end of the post url;
//title: the optional title of the post;
//*body: the full text body.
func (trc *TumblrRestClient) CreateText(blogname string, options map[string]string) error {
	requestUrl := fmt.Sprintf("/v2/blog/%s/post", blogname)
	options["type"] = "text"
	data := trc.request.Post(requestUrl, options)
	if data.Meta.Status != 201 {
		return errors.New(data.Meta.Msg)
	}
	return nil
}

//Create a quote post on a blog.
//blogname: the url of the blog you want to post to.
//options can be:
//(with * are marked required options)
//state: the state of the post(e.g. published, draft, queue, private);
//tags: a list of tags you want applied to the post;
//tweet: manages the autotweet for this post: set to off for no tweet
//or enter text to override the default tweet;
//date: the GMT date and time of the post as a string;
//format: sets the format type of the post(html or markdown);
//slug: add a short text summary to the end of the post url;
//*quote: the full text of the quote;
//source: the cited source of the quote.
func (trc *TumblrRestClient) CreateQuote(blogname string, options map[string]string) error {
	requestUrl := fmt.Sprintf("/v2/blog/%s/post", blogname)
	options["type"] = "quote"
	data := trc.request.Post(requestUrl, options)
	if data.Meta.Status != 201 {
		return errors.New(data.Meta.Msg)
	}
	return nil
}

//Create a link post on a blog.
//blogname: the url of the blog you want to post to.
//options can be:
//(with * are marked required options)
//state: the state of the post(e.g. published, draft, queue, private);
//tags: a list of tags you want applied to the post;
//tweet: manages the autotweet for this post: set to off for no tweet
//or enter text to override the default tweet;
//date: the GMT date and time of the post as a string;
//format: sets the format type of the post(html or markdown);
//slug: add a short text summary to the end of the post url;
//title: the title of the page the link points to;
//*url: the link you are posting;
//description: the description of the link you are posting.
func (trc *TumblrRestClient) CreateLink(blogname string, options map[string]string) error {
	requestUrl := fmt.Sprintf("/v2/blog/%s/post", blogname)
	options["type"] = "link"
	data := trc.request.Post(requestUrl, options)
	if data.Meta.Status != 201 {
		return errors.New(data.Meta.Msg)
	}
	return nil
}

//Create a chat post on a blog.
//blogname: the url of the blog you want to post to.
//options can be:
//(with * are marked required options)
//state: the state of the post(e.g. published, draft, queue, private);
//tags: a list of tags you want applied to the post;
//tweet: manages the autotweet for this post: set to off for no tweet
//or enter text to override the default tweet;
//date: the GMT date and time of the post as a string;
//format: sets the format type of the post(html or markdown);
//slug: add a short text summary to the end of the post url;
//title: the title of the chat;
//*conversation: the text of the conversation/chat, with dialogue labels.
func (trc *TumblrRestClient) CreateChatPost(blogname string, options map[string]string) error {
	requestUrl := fmt.Sprintf("/v2/blog/%s/post", blogname)
	options["type"] = "chat"
	data := trc.request.Post(requestUrl, options)
	if data.Meta.Status != 201 {
		return errors.New(data.Meta.Msg)
	}
	return nil
}

//Create an audio post on a blog.
//blogname: the url of the blog you want to post to.
//options can be:
//(with * are marked required options)
//state: the state of the post(e.g. published, draft, queue, private);
//tags: a list of tags you want applied to the post;
//tweet: manages the autotweet for this post: set to off for no tweet
//or enter text to override the default tweet;
//date: the GMT date and time of the post as a string;
//format: sets the format type of the post(html or markdown);
//slug: add a short text summary to the end of the post url;
//caption: the caption of the post;
//*external_url: the url of the site that hosts the audio file.
func (trc *TumblrRestClient) CreateAudio(blogname string, options map[string]string) error {
	requestUrl := fmt.Sprintf("/v2/blog/%s/post", blogname)
	options["type"] = "audio"
	data := trc.request.Post(requestUrl, options)
	if data.Meta.Status != 201 {
		return errors.New(data.Meta.Msg)
	}
	return nil
}

//Create a video post on a blog.
//blogname: the url of the blog you want to post to.
//options can be:
//(with * are marked required options)
//state: the state of the post(e.g. published, draft, queue, private);
//tags: a list of tags you want applied to the post;
//tweet: manages the autotweet for this post: set to off for no tweet
//or enter text to override the default tweet;
//date: the GMT date and time of the post as a string;
//format: sets the format type of the post(html or markdown);
//slug: add a short text summary to the end of the post url;
//caption: the caption for the post;
//*embed: the html embed code for the video.
func (trc *TumblrRestClient) CreateVideo(blogname string, options map[string]string) error {
	requestUrl := fmt.Sprintf("/v2/blog/%s/post", blogname)
	options["type"] = "video"
	data := trc.request.Post(requestUrl, options)
	if data.Meta.Status != 201 {
		return errors.New(data.Meta.Msg)
	}
	return nil
}

//Creates a reblog on the given blog.
//blogname: the url of the blog you want to reblog to.
//options should be:
//(with * are marked required options)
//*id: the id of the reblogged post;
//*reblog_key: the reblog key of the rebloged post.
func (trc *TumblrRestClient) Reblog(blogname string, options map[string]string) error {
	requestUrl := fmt.Sprintf("/v2/blog/%s/post/reblog", blogname)
	data := trc.request.Post(requestUrl, options)
	if data.Meta.Status != 201 {
		return errors.New(data.Meta.Msg)
	}
	return nil
}

//Deletes a post with a given id.
//blogname: the url of the blog you want to delete from.
//id: the id of the post you want to delete.
func (trc *TumblrRestClient) DeletePost(blogname, id string) error {
	requestUrl := fmt.Sprintf("/v2/blog/%s/post/delete", blogname)
	params := map[string]string{"id": id}
	data := trc.request.Post(requestUrl, params)
	if data.Meta.Status != 200 {
		return errors.New(data.Meta.Msg)
	}
	return nil
}

//Edits a post with a given id.
//blogname: the url of the blog you want to post to.
//options can be:
//(with * are marked required options)
//tags: a list of tags you want applied to the post;
//tweet: manages the autotweet for this post: set to off for no tweet
//or enter text to override the default tweet;
//date: the GMT date and time of the post as a string;
//format: sets the format type of the post(html or markdown);
//slug: add a short text summary to the end of the post url;
//*id: the id of the post.
//The other options are specific to the type of post you want to edit.
func (trc *TumblrRestClient) EditPost(blogname string, options map[string]string) error {
	requestUrl := fmt.Sprintf("/v2/blog/%s/post/edit", blogname)
	data := trc.request.Post(requestUrl, options)
	if data.Meta.Status != 200 {
		return errors.New(data.Meta.Msg)
	}
	return nil
}
