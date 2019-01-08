package main

import "testing"

func TestMain(t *testing.T) {
	var test = `
one
two two
three three three
four four four four
isn't isn't
Africa
sun-shine
abc
abd
abe
a
b
c
`

	expectedResults := map[string]uint64{
		"one":       1,
		"two":       2,
		"three":     3,
		"four":      4,
		"abc":       1,
		"abd":       1,
		"abe":       1,
		"a":         1,
		"b":         1,
		"c":         1,
		"isn't":     2,
		"africa":    1,
		"sun-shine": 1,
	}
	for i := 0; i < 100; i++ {
		sortedwords := sortWords(countWords(findWords(test)))
		for _, wd := range sortedwords {
			if expectedResults[wd.word] != wd.count {
				t.Logf("On run %d: %s did not match. Expected: %d, Got: %d",
					i,
					wd.word,
					expectedResults[wd.word],
					wd.count)
				t.FailNow()
			}
		}
	}
}
