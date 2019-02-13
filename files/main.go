package main

import (
	"fmt"
	"time"

	"github.com/erikwilliamsa/go-katas/files/parsing"
)

func main() {
	start := time.Now()
	parsing.PipedGetAllFileStats("/Users/exw5373/Documents/files")
	fmt.Println(time.Since(start))

	start = time.Now()
	parsing.BlockingAllFileStats("/Users/exw5373/Documents/files")
	fmt.Println(time.Since(start))

	start = time.Now()
	parsing.ConcurrentAllFileStats("/Users/exw5373/Documents/files")
	fmt.Println(time.Since(start))

}
