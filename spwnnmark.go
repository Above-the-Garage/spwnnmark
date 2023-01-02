/*
Package main - spwnncli
*/
package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/above-the-garage/spwnn"
)

//
// Parallel version of benchmarking code
//

var (
	mu    sync.Mutex
	dicts []*spwnn.SpwnnDictionary
)

// Manage a set of dictionaries, growing the list on demand
func getDict() *spwnn.SpwnnDictionary {
	mu.Lock()
	defer mu.Unlock()
	for i, dict := range dicts {
		if dict != nil {
			res := dict
			// Unlink to make it inaccessible
			// Go routine will add it back when done
			dicts[i] = nil
			return res
		}
	}
	newDict := spwnn.ReadDictionary(false)
	return newDict
}

func releaseDict(dict *spwnn.SpwnnDictionary) {
	mu.Lock()
	defer mu.Unlock()
	for i, d := range dicts {
		if d == nil {
			dicts[i] = dict
			return
		}
	}
	// All full, must have created a new dictionary, save for reuse
	dicts = append(dicts, dict)
}

var tokens = make(chan int, runtime.GOMAXPROCS(0))

func goCorrectSpelling(word string, noisy bool, strictLen bool) {
	//start := time.Now()

	dict := getDict()
	correctedWords, _ := spwnn.CorrectSpelling(dict, word, strictLen)

	if len(correctedWords) != 1 {
		fmt.Printf("Parallel validation:  '%s' could be '%v'\n", word, correctedWords)
	}

	// Let next go routine run
	releaseDict(dict)
	<-tokens
}

func benchmarkParallel(words []string, input string, noisy bool) {
	letters := spwnn.RemoveSpaces(input)
	if len(letters) == 0 {
		letters = "abcdefghijklmnopqrstuvwxyz"
	}
	fmt.Printf("Letters = '%s'; noisy = %v\n", letters, noisy)

	start := time.Now()

	for _, word := range words {
		testLetterIndex := 0
		if word[0] == '_' {
			testLetterIndex++
		}
		testLetter := fmt.Sprintf("%c", word[testLetterIndex])
		if strings.ContainsAny(testLetter, letters) {
			tokens <- len(tokens)
			go goCorrectSpelling(word, noisy, true /* strictLen */)
		}
	}

	// Wait for all go routines to complete (this may be considered bad form!)
	for len(tokens) > 0 {
		time.Sleep(10 * time.Millisecond)
	}
	elapsed := time.Since(start)

	fmt.Printf("Elapsed time %s; GOMAXPROCS %d\n", elapsed, runtime.GOMAXPROCS(0))
}

func main() {

	dict := spwnn.ReadDictionary(true)
	benchmarkParallel(spwnn.GetWords(dict), "" /* entire dictionary */, true)
	os.Exit(0)

}
