package gotumblr

import "github.com/garyburd/go-oauth/oauth"

//Make queries to the Tumblr API through TumblrRequest
type TumblrRequest struct {
	consumer *oauth.Client
	token *oauth.Credentials
	host string
}

//Initializes the TumblrRequest.
//consumerKey is the consumer key of your Tumblr Application
//consumerSecret is the consumer secret of your Tumblr Application
//oauthToken is the user specific token, received from the /access_token endpoint
//oauthSecret is the user specific secret, received from the /access_token endpoint
//host is the host that you are tryng to send information to (e.g. http://api.tumblr.com)
func NewTumblrRequest(consumerKey string, consumerSecret string, oauthToken string, oauthSecret string, host string) *TumblrRequest {
	consumer := oauth.Client{
		Credentials: oauth.Credentials{
			Token: consumerKey
			Secret: consumerSecret
		},
		TemporaryCredentialRequestURI: "http://www.tumblr.com/oauth/request_token",
		ResourceOwnerAuthorisationURI: "http://www.tumblr.com/oauth/authorize",
		TokenRequestURI: "http://www.tumblr.com/oauth/access_token",
	}
	token := oauth.Credentials{
		Token: oauthToken
		Secret: oauthSecret
	}
	return &TumblrRequest{&consumer, &token, host}
}

//Make a GET request to the API with properly formatted parameters
//url: the url you are making the request to
//params: the parameters needed for the request 
func (tr TumblrRequest) Get(url string, params map[string]string) map[string]interface{}{} {

}

//Makes a POST request to the API, allows for multipart data uploads
//url: the url you are making the request to
//params: all the parameters needed for the request
//files: list of files
func (tr TumblrRequest) Post(url string, params map[string]string, files []string) map[string]interface{}{} {

}

//Parse JSON response.
//content: the content returned from the web request to be pares as JSON
func (tr TumblrRequest) JsonParse(content) map[string]interface{}{} {

}

//Generates and makes a multipart request for data files
//url: the url you are making the request to
//params: all parameters needed for the request
//files: a list of files
func (tr TumblrRequest) PostMultipart(url string, params map[string]string, files []string) map[string]interface{}{} {

}

//Properly encodes the multipart body of the request
//fields: the parameters used in the request
//files: a list of lists containing information about the files
func (tr TumblrRequest) EncodeMultipartFormdata(fields map[string]string, files []string) []string {

}