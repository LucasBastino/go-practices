package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type result struct {
	url     string
	err     error
	latency time.Duration
}

func main() {
	resultsChan := make(chan result)
	urls := []string{
		"https://amazon.com",
		"https://google.com",
		"https://nytimes.com",
		"https://wsj.com",
		"http://localhost:8085/",
	}

	for _, url := range urls {
		go checkURL(url, resultsChan)
	}

	for range urls {
		result := <-resultsChan
		if result.err != nil {
			log.Fatal(result.err.Error())
		} else {
			fmt.Printf("url: %s, %v\n", result.url, result.latency)
		}
	}

}

func checkURL(url string, resultsChan chan result) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		resultsChan <- result{url, err, 0}
	} else {
		latency := time.Since(start).Round(time.Millisecond)
		resultsChan <- result{url, nil, latency}
		resp.Body.Close()
	}

}
