package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	wordmap := countWords(findWords(test1))
	//for word, count := range wordmap {
	//	fmt.Printf("MAP: %s, %d\n", word, count)
	//}
	sortedwords := sortWords(wordmap)
	for _, word := range sortedwords {
		fmt.Printf("%s: %d\n", word.word, word.count)
	}
}

func findWords(text string) []string {
	// Use a regex to find strings that start with a letter,
	//have letters or single quotes or a - and end with a letter.
	wordMatcher := regexp.MustCompile(`[a-z]+(?:\-[a-z]+|[a-z'])?[a-z]?`)

	words := wordMatcher.FindAllString(strings.ToLower(text), -1)
	return words
}

func countWords(wordslice []string) map[string]uint64 {
	wordMap := make(map[string]uint64)

	for _, value := range wordslice {
		wordMap[value]++
	}

	return wordMap
}

type worddata struct {
	word  string
	count uint64
}

func sortWords(wordMap map[string]uint64) []worddata {
	sortedSlice := make([]worddata, 0)
	firstWord := true

	for word, count := range wordMap {
		wd := worddata{
			word:  word,
			count: count,
		}
		// if its the first time we need to skip processing and add it in.
		if firstWord {
			fmt.Println("First word", word, count)
			sortedSlice = append(sortedSlice, wd)
			firstWord = false
			continue
		}
		// if count is lower than point in slice place before index then break
		// if count is equal place it after index then break
		// if its larger move on to the next index, if index + 1 == length then place at the end
		for index, sValue := range sortedSlice {
			if count < sValue.count {
				fmt.Println("Inserting <:", word, index)
				fmt.Println("Current:", sortedSlice)

				firstHalf := sortedSlice[:index]
				secondHalf := sortedSlice[index:]
				fmt.Println("firstHalf:", firstHalf)
				fmt.Println("secondHalf:", secondHalf)

				sortedSlice = append(firstHalf, wd)
				fmt.Println("sortedSlice - firstHalf:", sortedSlice)

				sortedSlice = append(sortedSlice, secondHalf...)
				fmt.Println("Adding:", secondHalf)
				fmt.Println("secondHalf:", secondHalf)
				fmt.Println("sortedSlice - SecondHalf:", sortedSlice)

				break
			} else if count == sValue.count {
				//fmt.Println("Inserting ==:", word, index)
				//fmt.Println("Current:", sortedSlice)
				var firstHalf []worddata
				var secondHalf []worddata
				if index == 0 {
					firstHalf = make([]worddata, 0)
					secondHalf = sortedSlice
				} else if index+1 == len(sortedSlice) {
					firstHalf = sortedSlice
					secondHalf = make([]worddata, 0)
				} else {
					firstHalf = sortedSlice[:index]
					secondHalf = sortedSlice[index:]
				}
				//fmt.Println("firstHalf:", firstHalf)
				//fmt.Println("secondHalf:", secondHalf)
				sortedSlice = append(firstHalf, wd)
				//fmt.Println("sortedSlice - firstHalf:", sortedSlice)
				sortedSlice = append(sortedSlice, secondHalf...)
				//fmt.Println("sortedSlice - SecondHalf:", sortedSlice)
				break
			} else if index+1 == len(sortedSlice) {
				sortedSlice = append(sortedSlice, wd)
				break
			}
		}
	}
	return sortedSlice
}
