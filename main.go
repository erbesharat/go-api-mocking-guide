package main

import (
	"bytes"
	"fmt"
	"net/http"
)

type Post struct {
	Name   string
	Author string
}

func main() {
	fmt.Println("Running...")
}

func POSTrequest() *http.Response {
	content := []byte(`{"title":"mocking apis in go."}`)
	resp, _ := http.Post("https://example.com/resource", "application/json", bytes.NewBuffer(content))
	return resp
}

func GETrequest() *http.Response {
	resp, _ := http.Get("https://example.com/resource")
	return resp
}

func MultipleGetRequests() (*http.Response, *http.Response) {
	resp, _ := http.Get("https://example.com/resource")
	resp2, _ := http.Get("https://example.com/post")
	return resp, resp2
}
