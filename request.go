//A Go Tumblr API v2 Client.

package gotumblr

import (
	"encoding/json"
	"github.com/kurrik/oauth1a"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

//TumblrRequest a structure to connect to Tumblr
type TumblrRequest struct {
	service    *oauth1a.Service
	userConfig *oauth1a.UserConfig
	host       string
	apiKey     string
}

//NewTumblrRequest initializes the TumblrRequest.
//consumerKey is the consumer key of your Tumblr Application.
//consumerSecret is the consumer secret of your Tumblr Application.
//callbackUrl is the callback URL of your Tumblr Application.
//oauthToken is the user specific token, received from the /access_token endpoint.
//oauthSecret is the user specific secret, received from the /access_token endpoint.
//host is the host that you are tryng to send information to (e.g. http://api.tumblr.com).
func NewTumblrRequest(consumerKey, consumerSecret, oauthToken, oauthSecret, callbackURL, host string) *TumblrRequest {
	service := &oauth1a.Service{
		RequestURL:   "http://www.tumblr.com/oauth/request_token",
		AuthorizeURL: "http://www.tumblr.com/oauth/authorize",
		AccessURL:    "http://www.tumblr.com/oauth/access_token",
		ClientConfig: &oauth1a.ClientConfig{
			ConsumerKey:    consumerKey,
			ConsumerSecret: consumerSecret,
			CallbackURL:    callbackURL,
		},
		Signer: new(oauth1a.HmacSha1Signer),
	}
	userConfig := oauth1a.NewAuthorizedConfig(oauthToken, oauthSecret)
	return &TumblrRequest{service, userConfig, host, consumerKey}
}

//Get makes a GET request to the API with properly formatted parameters.
//requestURL: the url you are making the request to.
//params: the parameters needed for the request.
func (tr *TumblrRequest) Get(requestURL string, params map[string]string) (CompleteResponse, error) {
	fullURL := tr.host + requestURL
	if len(params) != 0 {
		values := url.Values{}
		for key, value := range params {
			values.Set(key, value)
		}
		fullURL = fullURL + "?" + values.Encode()
	}
	httpRequest, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}
	tr.service.Sign(httpRequest, tr.userConfig)
	var httpResponse *http.Response
	httpClient := new(http.Client)
	httpResponse, err2 := httpClient.Do(httpRequest)
	if err2 != nil {
		return nil, err2
	}
	defer httpResponse.Body.Close()
	body, err3 := ioutil.ReadAll(httpResponse.Body)
	if err3 != nil {
		return nil, err3
	}
	return tr.JSONParse(body), nil
}

//Post makes a POST request to the API, allows for multipart data uploads.
//requestURL: the url you are making the request to.
//params: all the parameters needed for the request.
func (tr *TumblrRequest) Post(requestURL string, params map[string]string) (CompleteResponse, error) {
	fullURL := tr.host + requestURL
	values := url.Values{}
	for key, value := range params {
		values.Set(key, value)
	}
	httpRequest, err := http.NewRequest("POST", fullURL, strings.NewReader(values.Encode()))
	if err != nil {
		return nil, err
	}
	httpRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	tr.service.Sign(httpRequest, tr.userConfig)
	var httpResponse *http.Response
	httpClient := new(http.Client)
	httpResponse, err2 := httpClient.Do(httpRequest)
	if err2 != nil {
		return nil, err2
	}
	defer httpResponse.Body.Close()
	body, err3 := ioutil.ReadAll(httpResponse.Body)
	if err3 != nil {
		return nil, err3
	}
	return tr.JSONParse(body)
}

//JSONParse is a convenience function to parse JSON response.
//content: the content returned from the web request to be parsed as JSON.
func (tr *TumblrRequest) JSONParse(content []byte) (CompleteResponse, error) {
	var data CompleteResponse
	err := json.Unmarshal(content, &data)
	if err != nil {
		return nil, err
	}
	return data
}
