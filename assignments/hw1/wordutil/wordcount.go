package wordutil

import (
	"strings"
)

// Counts occurrences of each word in a string.
//
// Returns a map that stores each unique word in the string s as the key and
// its count as the corresponding value.
// Matching is case insensitive, e.g. "Orange" and "orange" is considered the
// same word.
func WordCount(s string) map[string]int {
	// TODO: implement me
	// HINT: You may find the `strings.Fields` and `strings.ToLower` functions helpful

	var s_split = strings.Fields(s)

	WordCountMap := make(map[string]int)

	for i := 0; i < len(s_split); i++ {
		s_split[i] = strings.ToLower(s_split[i])

		WordCountMap[s_split[i]] = WordCountMap[s_split[i]] + 1

	}

	return WordCountMap

}
