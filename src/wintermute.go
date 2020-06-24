package main

import (
	"fmt"
	"strings"
)

// we intern all words into an array, and reference to it with an int index
var words = []string{"this"}
var keys = make(map[string]int)

// int -> map[int]int
type frequencyChain map[int](map[int]int)

func normalize(chunk string) []string {
	// all bits of the form <thing>"." -> <thing>[ space ]"."
	// same with commas, same with semicolons, question marks, explanations.
	return []string{}
}

func getWord(key int) string {
	return words[key]
}

func makeWord(word string) int {
	if key, prs := keys[word]; prs {
		return key
	}

	key := len(words)

	keys[word] = key
	words = append(words, word)
	return key
}

func processChunk(chain frequencyChain, chunk string) {

}

func makeChain(corpusDir string) (frequencyChain, error) {
	return make(frequencyChain), nil
}

func predict(chain frequencyChain, previous int) int {
	return 0
}

func generate(chain frequencyChain, seed string) []string {
	last := 0                    // TODO: get a random word? or input?
	sequence := make([]int, 100) // whatever length we want this to be.
	for i := range words {
		last = predict(chain, last)
		sequence[i] = last
	}

	unkeyedSequence := make([]string, 100)
	for i, key := range sequence {
		unkeyedSequence[i] = getWord(key)
	}

	return append([]string{seed}, unkeyedSequence...)
}

func start() {
	chain, _ := makeChain("./corpus")

	sequence := generate(chain, "the")
	fmt.Println(strings.Join(sequence, " "))
}

func main() {
	start()
}
