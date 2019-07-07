package main

import (
	"sync"
)

func main() {
	intStrm := make(chan int, 2)
	wg := sync.WaitGroup{}

	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			intStrm <- i
		}
		close(intStrm)
	}()

	go func() {
		defer wg.Done()
		for i := range intStrm {
			println(i)
		}
	}()

	wg.Wait()
}
