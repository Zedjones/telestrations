package telestrationsLib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

const nounURL = "https://api.datamuse.com/words?rel_spc=%s&md=fp"
const sugURL = "https://api.datamuse.com/sug?s=%s&max=%d"

//Word is a single word returned from the /words endpoint in DataMuse
type Word struct {
	Word  string   `json:"word"`
	Score int      `json:"score"`
	Tags  []string `json:"tags"`
}

func contains(lst []string, item string) bool {
	for _, n := range lst {
		if item == n {
			return true
		}
	}
	return false
}

func WordFilter(vs []Word, f func(Word) bool) []Word {
	vsf := make([]Word, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func getNounsForWord(word string, wordPipe chan<- []Word) {
	formattedURL := fmt.Sprintf(nounURL, word)
	resp, err := http.Get(formattedURL)
	if err != nil {
		fmt.Println(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	var result []Word
	json.Unmarshal(body, &result)
	wordPipe <- result
}

func GetNouns(words []string) []Word {
	var wg sync.WaitGroup
	wg.Add(len(words))
	var allWords []Word
	allChan := make(chan []Word, len(words))
	for _, word := range words {
		go getNounsForWord(word, allChan)
	}
	go func() {
		for wordSlice := range allChan {
			allWords = append(allWords, wordSlice...)
			wg.Done()
		}
	}()
	wg.Wait()
	return allWords
}

func GetRandomWords(sug string, limit int) []Word {
	formattedURL := fmt.Sprintf(sugURL, sug, limit)
	resp, err := http.Get(formattedURL)
	if err != nil {
		fmt.Println(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	var result []Word
	var wordList []string
	json.Unmarshal(body, &result)
	for _, item := range result {
		wordList = append(wordList, item.Word)
	}
	nouns := GetNouns(wordList)
	nouns = WordFilter(nouns, func(word Word) bool {
		return contains(word.Tags, "n")
	})
	fmt.Println(nouns)
	fmt.Println(len(nouns))
	return nouns
}
