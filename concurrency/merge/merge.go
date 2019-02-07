package main

import (
	"fmt"
	"sync"
	"time"
)

/*Work in progress*/
func main() {
	const count = 10
	var chans [count]chan string

	for i := 0; i < count; i++ {
		chans[i] = make(chan string)
	}
	go func() {
		for i := 0; i < count; i++ {
			chans[i] <- fmt.Sprintf("This is message #%d", i+1)

		}
	}()

	done := make(chan struct{}, 2)
	result := merge(done, chans[:]...)
	time.AfterFunc(time.Second, func() {
		for _, c := range chans {
			close(c)
		}
	})
	x := 2
	for m := range result {
		fmt.Println(m, x)
		if x == 1 {
			fmt.Println("Early Close")
			done <- struct{}{}
		}
		x++
	}

}

func merge(done chan struct{}, chans ...chan string) chan string {

	var wg sync.WaitGroup

	out := make(chan string)

	output := func(c chan string) {
		defer func() {
			wg.Done()
		}()
		for data := range c {
			select {
			case out <- data:
			case <-done:
				fmt.Println("Close early", done)
				return
			}
		}
	}
	wg.Add(len(chans))
	for _, ch := range chans {
		go output(ch)
	}

	go func() {
		defer close(out)
		wg.Wait()
	}()

	return out
}
