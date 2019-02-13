package parsing

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

func PipedGetAllFileStats(dir string) {

	files, err := ioutil.ReadDir(dir)

	if err != nil {
		log.Fatal("ERROR:", err)
	}
	collectors := []<-chan TextFileStats{}
	wg := sync.WaitGroup{}
	for _, file := range files {

		if file.IsDir() {
			continue
		}
		if of, err := os.Open(dir + "/" + file.Name()); err == nil {
			//Ensure everyhing is created before continuing on.
			wg.Add(1)

			// For each file, create a starting point that will read the file.
			// Create 2 channels that will handle out of of a string intended
			// to be given to both the character counter and word counter.
			// The inputStart will not return until the waitgroup is done to
			// ensure all readers have been started as well.
			wcLineCh, charCtLineCh := inputStart(of, &wg)

			//Create the go routine that will count words in each line.
			//This returns a channel that the routine will output counts for
			//every input it is given.
			wctCh := pipeCountWords(wcLineCh)

			//Same as piped countwords, except the count is a map of all characters
			// in the string.
			charctCh := pipedCountChars(charCtLineCh)

			// For each file we want to begin combining the steps. A stat colloctor will begin
			// To "fan-in" the channels coming from the count routines.
			collectors = append(collectors, statCollector(of.Name(), wctCh, charctCh))
			wg.Done()

		}
	}

	// Create a channel that will return the combined stats from all files.
	dtsch := mergeStats(collectors...)
	dts := <-dtsch
	_ = dts.TotalCount //
}

func statCollector(filename string, wordcount <-chan int, charcount <-chan map[string]int) <-chan TextFileStats {
	tfsch := make(chan TextFileStats)
	go func() {
		wg := sync.WaitGroup{}
		wg.Add(2)
		tfs := TextFileStats{0, filename, map[string]int{}}

		go func() {
			for count := range wordcount {
				tfs.WordCount += count
			}
			wg.Done()
		}()
		go func() {
			for counts := range charcount {
				for k, v := range counts {
					tfs.CharCounts[k] += v
				}
			}
			wg.Done()
		}()
		wg.Wait()
		tfsch <- tfs

	}()
	return tfsch
}

func mergeStats(statcollection ...<-chan TextFileStats) <-chan DirTxtStats {
	dtsch := make(chan DirTxtStats)
	go func() {
		dts := DirTxtStats{TotalCharCounts: map[string]int{}, TextFileStats: []TextFileStats{}}

		for _, stats := range statcollection {
			fs := <-stats
			dts.TotalCount += fs.WordCount
			dts.TotalCharCounts = mergCharCountMap(dts.TotalCharCounts, fs.CharCounts)
			dts.TextFileStats = append(dts.TextFileStats, fs)
		}

		dtsch <- dts
		close(dtsch)
	}()
	return dtsch

}

func inputStart(of *os.File, wg *sync.WaitGroup) (wct, charct chan string) {
	wct, charct = make(chan string, 10), make(chan string, 10)
	wg.Add(1)

	go func() {
		scanner := bufio.NewScanner(of)
		defer close(wct)
		defer close(charct)
		defer of.Close()
		for scanner.Scan() {
			l := scanner.Text()
			wct <- l
			charct <- l
		}
		wg.Done()
		wg.Wait()

	}()
	return wct, charct
}

func pipeCountWords(lines <-chan string) <-chan int {
	count := make(chan int)
	go func() {
		for line := range lines {
			count <- CountWords(line)
		}
		close(count)
	}()
	return count

}

func pipedCountChars(lines <-chan string) chan map[string]int {
	counts := make(chan map[string]int)
	go func() {
		for line := range lines {
			counts <- AlphaNumericCount(line)
		}
		close(counts)
	}()
	return counts

}
