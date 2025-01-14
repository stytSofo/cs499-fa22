package wordutil

import (
	"strings"
)

// Finds first occurrence of each word in a string.
//
// Returns a map that stores each unique word in the string s as the key and
// the index of the first occurence of the word in the input string as the
// corresponding value.
// Matching is case insensitive, e.g. "Orange" and "orange" is considered the
// same word.
func WordIndex(s string) map[string]int {
	// TODO: implement me
	// HINT: You may find the `strings.Index` and `strings.ToLower` functions helpful

	var s_split = strings.Fields(s)

	IndexMap := make(map[string]int)

	for _, substring := range s_split {

		if IndexMap[substring] == 0 && substring == strings.ToLower(substring) {
			IndexMap[substring] = strings.Index(s, substring)
		}

	}

	return IndexMap

}
