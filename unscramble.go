package main

// unscramble.go reads a dictionary from file dict, register all words
// in a length indexed map with the sorted version of each workd as
// key and the word as value. Func unscramble is called with input
// sentence and dictionary reference as parameters, returning an
// unscrambled sentence and a true/false (false meaning we failed to
// unscrable, true meaning we managed to unscramble to a legal output
// - but not necessarily the intended sentence given a large
// dictionary)
//
// Example: elhloothtedrowl => hello to the world

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"sort"
	"strings"
)

var maxWordLen = 25 // limits the dictByLen array, we drop insanely long dict words

type wordByLen struct {
	len   int
	words map[string]string // one sortedWord => word map per word length
}

func main() {
	var inputSentence, dictFileName string
	if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "") {
		fmt.Println(path.Base(os.Args[0]), "[input_sentence] [dict_file]")
		os.Exit(1)
	} else if len(os.Args) > 1 {
		inputSentence = os.Args[1]
	} else {
		inputSentence = "elhloothtedrowl" // The default test case
	}

	if len(os.Args) > 2 {
		dictFileName = os.Args[2]
	} else {
		dictFileName = "dict"
	}

	// Read the dictionary file
	dictByLen := *readDict(dictFileName)

	// Unscramble sentense
	resultSentence, success := unscramble(inputSentence, &dictByLen)
	if success {
		fmt.Printf("%s => %s\n", inputSentence, resultSentence)
	} else {
		log.Fatalf("Not able to unscramble: '%s', incorrect result: '%s'\n", inputSentence, resultSentence)
	}
}

func readDict(dictFileName string) *[]wordByLen {
	// Read words from the dictionary
	file, err := os.Open(dictFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Give the dictByLen slice a skeleton
	var dictByLen []wordByLen
	for len := 0; len < maxWordLen; len++ {
		dictByLen = append(dictByLen, wordByLen{len: len, words: make(map[string]string)})
	}

	// Register all words in a length indexed map where sorted-word and word is key and value
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		charSlice := strings.Split(word, "")
		sort.Strings(charSlice)
		wordSorted := strings.Join(charSlice, "")
		if len(word) < maxWordLen {
			dictByLen[len(word)].words[wordSorted] = word
		}
	}

	return &dictByLen
}

// unscramble returns unscrambled resultSentence and true/false based on
// longest word first from the wordByLen dict
func unscramble(inputSentence string, dictByLen *[]wordByLen) (string, bool) {
	// foreach dict word length, check all sorted words against
	// the sorted part of the remaining string
	resultSentenceArr := []string{} // The result sentence
	prevLen := -1
	inputLen := len(inputSentence)
	for len(inputSentence) > 0 && len(inputSentence) != prevLen {
		prevLen = len(inputSentence)

		wlen := len(*dictByLen) - 1
		if wlen > maxWordLen {
			wlen = maxWordLen
		}
		if wlen > len(inputSentence) {
			wlen = len(inputSentence)
		}

		for ; wlen > 0; wlen-- {
			for wordSorted, word := range (*dictByLen)[wlen].words {
				inputChars := inputSentence[:wlen]
				inputString := strings.Split(inputChars, "")
				sort.Strings(inputString)
				inputSorted := strings.Join(inputString, "")

				if inputSorted == wordSorted {
					resultSentenceArr = append(resultSentenceArr, word)
					inputSentence = inputSentence[wlen:]
					goto nextWord
				}
			}
		}
	nextWord:
	}
	return strings.Join(resultSentenceArr, " "), inputLen == len(strings.Join(resultSentenceArr, ""))
}
