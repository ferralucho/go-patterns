package main

import (
	"fmt"
	"time"
)

func main(){
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)
	limiter := time.Tick(1000 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	rafaga := make(chan time.Time, 3)

	go func() {
		for t := range time.Tick(1000 * time.Millisecond){
			for i := 1; i <= 3; i++ {
				rafaga <- t
			}
		}
	}()

	rafagaRequest := make(chan int, 15)
	for i := 1; i <= 15; i++ {
		rafagaRequest <- i
	}
	close(rafagaRequest)
	for req := range rafagaRequest {
		<- rafaga
		fmt.Println("request", req, time.Now())
	}
}
