// Find the top K most common words in a text document.
// Input path: location of the document, K top words
// Output: Slice of top K words
// For this excercise, word is defined as characters separated by a whitespace
// Note: You should use `checkError` to handle potential errors.

package topwords

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func topWords(path string, K int) []WordCount {
	//open file using os
	f, err := os.Open(path)

	//check if path is valid
	checkError(err)

	//close file before exiting the program
	defer f.Close()

	//create scanner object to read contents of the file
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	//create map to keep count of words
	m := make(map[string]*WordCount)

	//read file line by line and append words by separating spaces
	for scanner.Scan() {
		_, pres := m[scanner.Text()]
		if pres == false {
			m[scanner.Text()] = &WordCount{Word: scanner.Text(), Count: 1}
		} else {
			m[scanner.Text()].Count++
		}

	}

	//create a slice of wordcount objects
	slce := make([]WordCount, 0, len(m))

	for k := range m {
		slce = append(slce, *m[k])
	}

	//fmt.Println(slce)
	//sort the slice
	sortWordCounts(slce)
	//fmt.Println(slce)
	return slce[:K]

}

//--------------- DO NOT MODIFY----------------!

// A struct that represents how many times a word is observed in a document
type WordCount struct {
	Word  string
	Count int
}

// Method to convert struct to string format
func (wc WordCount) String() string {
	return fmt.Sprintf("%v: %v", wc.Word, wc.Count)
}

// Helper function to sort a list of word counts in place.
// This sorts by the count in decreasing order, breaking ties using the word.
func sortWordCounts(wordCounts []WordCount) {
	sort.Slice(wordCounts, func(i, j int) bool {
		wc1 := wordCounts[i]
		wc2 := wordCounts[j]
		if wc1.Count == wc2.Count {
			return wc1.Word < wc2.Word
		}
		return wc1.Count > wc2.Count
	})
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
