package parsing

import (
	"bufio"
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
	TextFileStats   []TextFileStats
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

func GetTextFileStats(file *os.File) (tfs TextFileStats) {

	scanner := bufio.NewScanner(file)
	tfs = TextFileStats{FileName: file.Name(), WordCount: 0, CharCounts: map[string]int{}}

	for scanner.Scan() {
		line := scanner.Text()
		tfs.WordCount += CountWords(line)
		for k, v := range AlphaNumericCount(line) {
			tfs.CharCounts[k] += v
		}
	}
	return
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
