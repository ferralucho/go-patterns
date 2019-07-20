package main

import (
	"fmt"
	"net/http"
	"sync"
)

type Result struct {
	StatusCode int
	Url        string
}

func PingUrl(url string, group *sync.WaitGroup) (Result) {
	res, err := http.Get(url)
	if err != nil {
		print(err.Error())
	}
	result := Result{
		Url:        url,
		StatusCode: res.StatusCode,
	}

	//si me da error forzo un panic

	fmt.Printf("url: %s - status_code: %d\n",
		result.Url, result.StatusCode)
	group.Done()
	return result
}

func GetResults(urls []string) *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(len(urls))

	for _, urls := range urls {
		go PingUrl(urls, &wg)
	}
	return &wg
}

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.lavoz.com.ar",
	}

	wg := GetResults(urls)
	wg.Wait()
	fmt.Println("Proceso terminado")
}
