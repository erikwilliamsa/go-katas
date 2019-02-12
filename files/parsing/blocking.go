package parsing

import (
	"io/ioutil"
	"log"
	"os"
)

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
	//fmt.Println("Blocking Total:", dts.TotalCount)
	return dts
}

func mergCharCountMap(m1, m2 map[string]int) map[string]int {
	for k, v := range m2 {
		m1[k] += v
	}
	return m1
}
