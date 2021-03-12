package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request failed")

type requestResult struct {
	url    string
	status string
}

func main() {
	// map은 {}나 make 없이는 nil을 나타냄.
	// results := map[string]string{} // emptymap 생성 방법1
	c := make(chan requestResult)

	urls := []string{
		"https://www.airbnb.com",
		"https://www.google.com",
		"https://www.amazon.com",
		"https://www.reddit.com",
		"https://soundcloud.com",
		"https://www.facebook.com",
		"https://www.instagram.com",
		"https://academy.nomadcoders.co",
		"https://www.daum.net",
		"https://www.naver.com",
	}

	for _, url := range urls {
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)
	}
}

func hitURL(url string, c chan<- requestResult) {
	fmt.Println("Checking : ", url)
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err, resp.StatusCode)
		status = "FAILED"
	}

	c <- requestResult{url: url, status: status}
}
