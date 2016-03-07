package gotumblr

import (
	"encoding/json"
	"testing"
)

func TestNewTumblrRequest(t *testing.T) {
	c := NewTumblrRequest("", "", "", "", "", "http://api.tumblr.com")
	if c.host != "http://api.tumblr.com" {
		t.Errorf("New Client host = %v, want http://api.tumblr.com", c.host)
	}
	if c.apiKey != "" {
		t.Errorf("New Client host = %v, want the empty string", c.apiKey)
	}
}

func TestGet(t *testing.T) {
	setup()
	defer teardown()

	handleFunc("/v2/user/info", "GET", `{"response": {"user": {"name": "mgterzieva"}}}`, map[string]string{}, t)
	data := client.request.Get("/v2/user/info", map[string]string{})

	expectedMeta := MetaInfo{Msg: "", Status: 0}
	var response UserInfoResponse
	json.Unmarshal(data.Response, &response)
	expectedname := "mgterzieva"

	if data.Meta != expected_meta {
		t.Errorf("Get returned %+v, want %+v", data.Meta, expectedMeta)
	}
	if response.User.Name != expected_name {
		t.Errorf("Get returned %v, want %v", response.User.Name, expectedName)
	}
}

func TestPost(t *testing.T) {
	setup()
	defer teardown()

	response := `{"meta": {"status":404, "msg": "Not Found"}}`

	handleFunc("/v2/user/follow", "POST", response, map[string]string{"url": "thehungergames"}, t)
	data := client.request.Post("/v2/user/follow", map[string]string{"url": "thehungergames"})

	expectedMeta := MetaInfo{Msg: "Not Found", Status: 404}
	if data.Meta != expectedMeta {
		t.Errorf("Post returned %+v, want %+v", data.Meta, expectedMeta)
	}
}

func TestJSONParse(t *testing.T) {
	data := client.request.JSONParse([]byte(`{"meta": {"msg": "OK", "status": 200}, "response": {"user": {"name": "mgterzieva"}}}`))

	expectedMeta := MetaInfo{Msg: "OK", Status: 200}
	var response UserInfoResponse
	json.Unmarshal(data.Response, &response)
	expectedName := "mgterzieva"

	if data.Meta != expected_meta {
		t.Errorf("Get returned %+v, want %+v", data.Meta, expected_meta)
	}
	if response.User.Name != expected_name {
		t.Errorf("Get returned %v, want %v", response.User.Name, expected_name)
	}
}
