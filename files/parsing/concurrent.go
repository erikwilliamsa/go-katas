package parsing

import (
	"io/ioutil"
	"log"
	"os"
)

func ConcurrentAllFileStats(dir string) {

	dts := DirTxtStats{TotalCharCounts: map[string]int{}, TextFileStats: []TextFileStats{}}
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		log.Fatal("ERROR:", err)
	}

	chans := []chan TextFileStats{}

	for _, file := range files {

		if file.IsDir() {
			continue
		}
		if of, err := os.Open(dir + "/" + file.Name()); err == nil {

			chans = append(chans, ConcurrentFileStats(of))

		}
	}

	for _, ch := range chans {
		fs := <-ch
		dts.TotalCount += fs.WordCount
		dts.TotalCharCounts = mergCharCountMap(dts.TotalCharCounts, fs.CharCounts)
		dts.TextFileStats = append(dts.TextFileStats, fs)
	}

	//fmt.Println("Concurrent Count", dts.TotalCount)

}

func ConcurrentFileStats(of *os.File) chan TextFileStats {
	tfsCh := make(chan TextFileStats)
	go func() {
		defer of.Close()
		tfsCh <- GetTextFileStats(of)

	}()
	return tfsCh
}
