package domain

import (
	"sort"
	"strconv"
)

func WordCount(wordBase [][]string) map[string]int {
	wc := make(map[string]int)
	for _, words := range wordBase {
		for _, word := range words {
			wc[word] += 1
		}
	}
	return wc
}

func BuildHeadElems(wordCount map[string]int, supportCount int) (HeadElems, map[string]*HeadElem) {
	var headElems HeadElems
	var headAddr = make(map[string]*HeadElem)
	for word, count := range wordCount {
		if count >= supportCount {
			elem := HeadElem{word: word, count: count, treeNode: nil, pattern: nil}
			headAddr[word] = &elem
			headElems = append(headElems, elem)
		}
	}
	sort.Sort(sort.Reverse(headElems))
	return headElems, headAddr
}

func FilterWordBase(headAddr map[string]*HeadElem, wordBase [][]string) [][]string {
	var filteredWordBase [][]string
	for _, words := range wordBase {
		var filteredWords []string
		var wordSupport Pairs
		for _, word := range words {
			if headAddr[word] != nil {
				wordSupport = append(wordSupport, Pair{word, headAddr[word].count})
			}
		}
		sort.Sort(sort.Reverse(wordSupport))
		for _, pair := range wordSupport {
			filteredWords = append(filteredWords, pair.key)
		}
		if len(filteredWords) > 0 {
			filteredWordBase = append(filteredWordBase, filteredWords)
		}
	}
	return filteredWordBase
}

type TwoFreqItem struct {
	BaseWord     string
	Word         string
	SupportCount int
	Confidence   float64
}

func WordConcurrence(headAddr map[string]*HeadElem, confidence float64) []TwoFreqItem {
	var freqItems []TwoFreqItem
	for baseWord, headElem := range headAddr {
		for coWord, supportCount := range headElem.pattern {
			con := float64(supportCount) / float64(headElem.count)
			if con >= confidence {
				freqItems = append(freqItems, TwoFreqItem{baseWord, coWord, supportCount, con})
			}
			con = float64(supportCount) / float64(headAddr[coWord].count)
			if con >= confidence {
				freqItems = append(freqItems, TwoFreqItem{coWord, baseWord, supportCount, con})
			}
		}
	}

	return freqItems
}

func FreqItemsToStrings(items []TwoFreqItem) [][]string {
	var ret [][]string
	for _, item := range items {
		var value []string
		value = append(value, item.BaseWord)
		value = append(value, item.Word)
		value = append(value, strconv.Itoa(item.SupportCount))
		value = append(value, strconv.FormatFloat(item.Confidence, 'f', -1, 64))
		ret = append(ret, value)
	}
	return ret
}
