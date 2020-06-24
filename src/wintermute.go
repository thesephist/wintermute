package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// TODO: we should pull this markov chain into a separate package?

const maxSequenceLen = 200

// we intern all words into an array, and reference to it with an int index
var words = []string{"the"}
var keys = map[string]int{
	"the": 0,
}

// int -> map[int]int
type frequencyChain map[int](map[int]int)

func (fc *frequencyChain) addWord(last, key int) {
	keyFrequencyMap := (*fc)[last]
	if keyFrequencyMap == nil {
		keyFrequencyMap = make(map[int]int)
		(*fc)[last] = keyFrequencyMap
	}

	if freqs, prs := keyFrequencyMap[key]; prs {
		keyFrequencyMap[key] = freqs + 1
	} else {
		keyFrequencyMap[key] = 1
	}
}

func (fc *frequencyChain) predict(last int) int {
	keyFrequencyMap := (*fc)[last]
	if keyFrequencyMap == nil {
		log.Println("Hit dead-end, generating random word.")
		return rand.Intn(len(words))
	}

	freqSum := 0
	for _, freq := range keyFrequencyMap {
		freqSum += freq
	}

	pick := rand.Intn(freqSum)
	for key, freq := range keyFrequencyMap {
		pick -= freq
		if pick <= 0 {
			return key
		}
	}

	panic("unreachable!")
}

func normalize(chunk string) []int {
	lines := strings.Split(chunk, "\n")
	normedLines := []string{}
	for _, line := range lines {
		if strings.HasPrefix(line, "!") ||
			strings.HasPrefix(line, "#") ||
			strings.HasPrefix(line, "<") ||
			line == "---" {
			continue
		}

		normedLines = append(normedLines, line)
	}

	chunk = strings.Join(normedLines, " ")
	chunk = strings.ToLower(chunk)

	reg := regexp.MustCompile("\\s+")
	chunk = reg.ReplaceAllString(chunk, " ")

	reg = regexp.MustCompile(`[\[.+\]]`)
	chunk = reg.ReplaceAllString(chunk, "")

	reg = regexp.MustCompile(`[\(\)\{\}#_*\"]+`)
	chunk = reg.ReplaceAllString(chunk, "")

	reg = regexp.MustCompile(`([\.,:;?!]) `)
	chunk = reg.ReplaceAllString(chunk, " $1 ")

	words := strings.Split(chunk, " ")

	keys := make([]int, len(words))
	for i, word := range words {
		keys[i] = makeWord(word)
	}

	return keys
}

func render(words []string) string {
	last := "."
	result := ""
	for _, word := range words {
		if !strings.Contains(".,:;?!", word) {
			result += " "
		}

		if last == "." ||
			word == "i" ||
			strings.HasPrefix(word, "i'") ||
			strings.HasPrefix(word, "i’") {
			result += strings.ToTitle(word[0:1]) + word[1:]
		} else {
			result += word
		}

		last = word
	}

	return result[1:]
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
	keys := normalize(chunk)

	if len(keys) < 2 {
		fmt.Println("WARNING NOT HERE")
		return
	}

	last := keys[0]
	for _, key := range keys[1:] {
		chain.addWord(last, key)
		last = key
	}
}

func makeChain(corpusDir string) (frequencyChain, error) {
	chain := make(frequencyChain)

	entries, err := ioutil.ReadDir(corpusDir)
	if err != nil {
		return nil, err
	}

	// TODO: only [0:5]
	for _, entry := range entries[0:5] {
		if entry.IsDir() {
			continue
		}

		file, err := os.Open(filepath.Join(corpusDir, entry.Name()))
		if err != nil {
			return nil, err
		}
		defer file.Close()

		data, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}

		processChunk(chain, string(data))
	}

	return chain, nil
}

func generate(chain frequencyChain, seed string) []string {
	last := 0
	sequence := make([]int, maxSequenceLen)
	for i := range sequence {
		last = chain.predict(last)
		sequence[i] = last
	}

	unkeyedSequence := make([]string, maxSequenceLen)
	for i, key := range sequence {
		unkeyedSequence[i] = getWord(key)
	}

	return append([]string{seed}, unkeyedSequence...)
}

func start() {
	chain, _ := makeChain("./corpus")

	sequence := generate(chain, "the")
	fmt.Println(render(sequence))
}

func main() {
	start()
}
