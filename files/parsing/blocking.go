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
	TextFileStats   []TextFileStats
}

/*
BlockingAllFileStats Reads all files in a directory
and returns the word counts, alpha numeric character count
and individual TextFileStats.
*/
func BlockingAllFileStats(dir string) DirTxtStats {

	dts := DirTxtStats{TotalCharCounts: map[string]int{}, TextFileStats: []TextFileStats{}}
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		log.Fatal("ERROR:", err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if of, err := os.Open(dir + "/" + file.Name()); err == nil {
			fs := GetTextFileStats(of)
			of.Close()

			dts.TotalCount += fs.WordCount
			dts.TotalCharCounts = mergCharCountMap(dts.TotalCharCounts, fs.CharCounts)
			dts.TextFileStats = append(dts.TextFileStats, fs)

		}
	}
	return dts
}

func mergCharCountMap(m1, m2 map[string]int) map[string]int {
	for k, v := range m2 {
		m1[k] += v
	}
	return m1
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
