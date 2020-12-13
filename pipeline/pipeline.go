package main

import "fmt"

func generador(nums ...int) <-chan int {
	outCh := make(chan int)
	go func() {
		for _, n := range nums {
			outCh <- n
		}
		close(outCh)
	}()
	return outCh
}

func cuadrado(in <-chan int) <-chan int {
	outCh := make(chan int)

	go func() {
		for n := range in {
			outCh <- n * n
		}
		close(outCh)
	}()
	return outCh
}

func main() {
	c1 := generador(1, 2, 3, 4, 5)
	c2 := cuadrado(c1)

	for n := range c2 {
		fmt.Println(n)
	}
}
