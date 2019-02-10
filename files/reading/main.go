package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	slurp("./main.go")
	readLines("./main.go")

}

func slurp(path string) {
	if data, err := ioutil.ReadFile(path); err == nil {
		fmt.Println(string(data))

	}

}

func readLines(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	i, wordcount := 1, 0

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("%d:  %s\n", i, line)
		i++
		wordcount += countWords(line)
	}
	fmt.Printf("Word count: %d", wordcount)
}

func countWords(s string) int {
	fields := strings.Fields(s)
	words := []string{}
	for _, field := range fields {

		if notspecial := alphNumericOnly(field); notspecial != "" {
			words = append(words, notspecial)
		}

	}
	return len(words)
}

func alphNumericOnly(s string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(s)
		return ""
	}
	return reg.ReplaceAllString(s, "")
}
