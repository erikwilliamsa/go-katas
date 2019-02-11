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
	WordCount  int
	FileName   string
	CharCounts map[string]int
}

type DirTxtStats struct {
	TotalCount      int
	TotalCharCounts map[string]int
	TextFileStats   TextFileStats
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
			result = append(result, TextFileStats{
				count,
				name,
				map[string]int{},
			})
		}
	}
	return result
}

//AlphaNumericCount returns the count of all alpha numeric
// characters in the string as a map
func AlphaNumericCount(s string) (chars map[string]int) {
	chars = map[string]int{}
	eval := alphNumericOnly(s)
	for _, c := range eval {
		if string(c) != " " {
			chars[string(c)]++
		}
	}
	return
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
		wc += CountWords(line)
	}
	return wc
}

/*
CountWords counts words in  a string
where a word is any group of alpha numeric characters seperated
by spaces
*/
func CountWords(s string) int {
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
