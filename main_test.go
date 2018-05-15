package main

import (
	"net/http"
	"testing"

	"gopkg.in/jarcoal/httpmock.v1"
)

func TestPOSTrequest(t *testing.T) {
	posts := mockPostStruct()
	httpmock.Activate()
	httpmock.RegisterResponder("POST", "https://example.com/resource",
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, posts)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)
	response := POSTrequest()
	defer response.Body.Close()
	if response.StatusCode != 200 {
		t.Errorf("expected status code %d, got %d", 200, response.StatusCode)
	}
}

func TestGETrequest(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()

	httpmock.RegisterResponder("GET", "https://example.com/resource",
		httpmock.NewStringResponder(200, `[{"name": "Hello world!", "author": "someone"}]`))

	response := GETrequest()
	defer response.Body.Close()
	if response.StatusCode != 200 {
		t.Errorf("expected status code %d, got %d", 200, response.StatusCode)
	}
}

func TestMultipleGetRequests(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()
	httpmock.RegisterResponder("GET", "https://example.com/resource",
		httpmock.NewStringResponder(200, `[{"name": "Hello world!", "author": "someone"}]`))
	httpmock.RegisterResponder("GET", "https://example.com/post",
		httpmock.NewStringResponder(200, `[{"name": "Hello world #2", "author": "someone else"}]`))
	resp, resp2 := MultipleGetRequests()
	if resp.StatusCode != 200 && resp2.StatusCode != 200 {
		t.Errorf("expected status code %d and %d, got %d", 200, resp.StatusCode, resp2.StatusCode)
	}
}

func TestMultipleGetRequestsUsingMockFunction(t *testing.T) {
	mockURL([]string{"https://example.com/resource", "https://example.com/post"})
	defer httpmock.Deactivate()
	resp, resp2 := MultipleGetRequests()
	if resp.StatusCode != 200 && resp2.StatusCode != 200 {

		t.Errorf("expected status code %d and %d, got %d", 200, resp.StatusCode, resp2.StatusCode)
	}
}

func mockPostStruct() []Post {
	posts := []Post{}
	for i := 0; i <= 2; i++ {
		post := Post{
			Name:   "hello world #" + string(i),
			Author: "author #" + string(i),
		}
		posts = append(posts, post)
	}
	return posts
}

func mockURL(urls []string) {
	httpmock.Activate()
	for _, url := range urls {
		httpmock.RegisterResponder("GET", url,
			httpmock.NewStringResponder(200, `[{"name": "Post name", "author": "someone"}]`))
	}
}
