package main

import (
	"fmt"
	"net/http"
	"time"
)

type HttpResponse struct {
	url      string
	response *http.Response
	err      error
}

func asyncGets(urls []string) []*HttpResponse {

	ch := make(chan *HttpResponse)
	responses := []*HttpResponse{}

	for _, url := range urls {
		go func(url string) {
			fmt.Printf("Fetching url %s\n", url)
			resp, err := http.Get(url)
			ch <- &HttpResponse{url, resp, err}
		}(url)
	}

	for {
		select {
		case r := <-ch:
			fmt.Printf("%s was fetched \n", r.url)
			responses = append(responses, r)
			if len(responses) == len(urls) {
				return responses
			}
		case <-time.After(50 * time.Millisecond):
			fmt.Printf(".")
		}
	}

	return responses
}

func main() {
	fmt.Printf("Async requests")
	urls := []string{
		"http://google.com",
		"http://golang.org/",
		"http://matt.aimonetti.net/",
	}
	results := asyncGets(urls)
	for _, result := range results {
		fmt.Printf("%s status %s \n", result.url, result.response.Status)
	}
}
