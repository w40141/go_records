package main

import (
	"fmt"
)

func receive(ch <-chan int) {
	for {
		i := <-ch
		j := i * i
		fmt.Println(j)
	}
}

func main() {
	ch := make(chan int)

	go receive(ch)

	i := 0
	for i < 10 {
		ch <- i
		i++
	}
}
