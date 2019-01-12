package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	searchTerm  = flag.String("s", "", "Search term")
	ignoreCase  = flag.Bool("i", false, "Ignore the case")
	invertMatch = flag.Bool("v", false, "Inverts the match to include lines that don't that the search term.")
)

func main() {
	flag.Parse()
	if len(*searchTerm) == 0 {
		fmt.Println("No search term passed in!")
		os.Exit(1)
	}
	if *ignoreCase {
		*searchTerm = strings.ToLower(*searchTerm)
	}

	fmt.Println(
		strings.Join(
			search(
				*searchTerm,
				lines(book, *ignoreCase),
				*invertMatch,
			),
			"\n",
		),
	)
}

// lines returns a slice of the text broken apart on the new lines.
func lines(text string, toLower bool) []string {
	if toLower {
		text = strings.ToLower(text)
	}
	return strings.Split(text, "\n")
}

func search(term string, lines []string, invert bool) []string {
	re := regexp.MustCompile(term)
	returnLines := make([]string, 0)
	for _, line := range lines {
		if matcher(re.Match([]byte(line)), invert) {
			returnLines = append(returnLines, line)
		}
	}
	return returnLines
}

func matcher(match, invert bool) bool {
	if invert {
		return !match
	}
	return match
}
