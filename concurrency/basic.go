package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)

	go func() {
		defer close(ch)
		for i := 0; i < 4; i++ {
			ch <- fmt.Sprintf("%d", i)
		}
	}()

	for m := range ch {
		fmt.Println(m)

	}
}
