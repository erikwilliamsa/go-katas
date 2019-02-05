package main

import (
	"fmt"
	"sync"
)

/*Work in progress*/
func main() {
	const count = 10
	var chans [10]chan string
	for i := 0; i < count; i++ {
		chans[i] = make(chan string)
	}

	result := merge(chans[:]...)
	go func() {
		for i := 0; i < count; i++ {
			chans[i] <- fmt.Sprintf("This is message #%d", i+1)
		}
	}()

	fmt.Println("Hi")
	for msg := range result {
		fmt.Println(msg)
	}
}

func merge(chans ...chan string) chan string {

	var wg sync.WaitGroup

	out := make(chan string)
	wg.Add(len(chans))

	output := func(c chan string) {
		for data := range c {
			out <- data
		}
		wg.Done()
	}

	for _, ch := range chans {
		go output(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
