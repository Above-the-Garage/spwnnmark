/*
Package main - spwnncli
*/
package main

import (
	"flag"
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
	mu           sync.Mutex
	dicts        []*spwnn.SpwnnDictionary
	dictFilename string
	scoring      spwnn.PairScoring
	lengthMatch  spwnn.LengthMatch
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
	newDict := spwnn.ReadDictionary(dictFilename)
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

func goCorrectSpelling(word string, noisy bool) {
	//start := time.Now()

	dict := getDict()
	correctedWords, _ := spwnn.CorrectSpelling(dict, word, lengthMatch, scoring)

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
			go goCorrectSpelling(word, noisy)
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
	dictFlag := flag.String("dict", "knownWords.txt", "dictionary file to use")
	lengthFlag := flag.String("length", "same", "which winners to keep: any length, or the same length as the input")
	scoringFlag := flag.String("scoring", "capped", "how repeated letter pairs are scored: capped or original")
	flag.Parse()

	var err error
	scoring, err = parseScoring(*scoringFlag)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	lengthMatch, err = parseLength(*lengthFlag)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	fmt.Printf("Pair scoring: %s; length match: %s\n", *scoringFlag, *lengthFlag)

	dictFilename = *dictFlag
	loadStart := time.Now()
	dict := spwnn.ReadDictionary(dictFilename)
	fmt.Printf("%d words in dictionary, loaded in %v\n", dict.WordCount(), time.Since(loadStart).Round(time.Millisecond))
	benchmarkParallel(dict.Words(), "" /* entire dictionary */, true)
	os.Exit(0)
}

// parseScoring converts the -scoring flag text to a spwnn.PairScoring
func parseScoring(name string) (spwnn.PairScoring, error) {
	switch name {
	case "capped":
		return spwnn.CappedPairs, nil
	case "original":
		return spwnn.OriginalPairs, nil
	}
	return spwnn.CappedPairs, fmt.Errorf("unknown scoring %q (want capped or original)", name)
}

// parseLength converts the -length flag text to a spwnn.LengthMatch
func parseLength(name string) (spwnn.LengthMatch, error) {
	switch name {
	case "any":
		return spwnn.AnyLength, nil
	case "same":
		return spwnn.SameLength, nil
	}
	return spwnn.AnyLength, fmt.Errorf("unknown length %q (want any or same)", name)
}
