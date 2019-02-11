package parsing

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

type TextFileStats struct {
	WordCount int
	FileName  string
}

/*
BlockingFileStats Reads all files in a directory
and returns the word counts.
*/
func BlockingFileStats(dir string) []TextFileStats {

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal("ERROR:", err)
	}
	result := []TextFileStats{}

	for _, f := range files {
		if !f.IsDir() {
			name := dir + "/" + f.Name()
			count := wordCount(name)
			result = append(result, TextFileStats{count, name})
		}
	}
	return result
}

func wordCount(path string) int {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	i, wc := 1, 0

	for scanner.Scan() {
		line := scanner.Text()
		i++
		wc += countWords(line)
	}
	return wc
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
