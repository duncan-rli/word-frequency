package main

import (
	"sort"
)

type Data struct {
	word  []byte
	count int
}

var wordStore map[string]int

func CheckTreeContainsAndUpdate(word []byte) {

	if wordStore == nil {
		wordStore = make(map[string]int)
	}

	count, present := wordStore[string(word)]
	if present {
		wordStore[string(word)] = count + 1
	} else {
		wordStore[string(word)] = 1
	}
}

func FindTwentyMostCommon(freqList *[20]Data) {

	for w, c := range wordStore {
		addWordToList(freqList, Data{[]byte(w), c})
	}
}

func addWordToList(freqList *[20]Data, val Data) {

	// as list is sorted,  the bottom of the list should be the current min
	if (*freqList)[19].count < val.count {
		(*freqList)[19] = val
		sort.Slice(freqList[:], func(i, j int) bool {
			return freqList[i].count > freqList[j].count
		})
	}
}
