package main

import (
	"fmt"
	"net/http"
)

type Result struct {
	StatusCode int
	Url        string
}

func PingUrl(url string, ch chan Result) {
	res, err := http.Get(url)
	if err != nil {
		print(err.Error())
	}
	result := Result{
		Url:        url,
		StatusCode: res.StatusCode,
	}
	ch <- result
}

func GetResults(urls []string) <-chan Result {
	ch := make(chan Result, len(urls))
	for _, urls := range urls {
		go PingUrl(urls, ch)
	}
	return ch
}

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.facebook.com",
	}
	ch := GetResults(urls)

	for i := 0; i < len(urls); i++ {
		result := <-ch
		fmt.Printf("url: %s - status_code: %d\n",
			result.Url, result.StatusCode)
	}
	fmt.Println("Proceso terminado")
}
