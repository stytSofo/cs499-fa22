package wordutil

import (
	"strings"

	"golang.org/x/exp/slices"
)

// Finds all occurrences of each word in a string.
//
// Returns a map that stores each unique word in the string s as the key and
// a slice that contains the index of each occurence of the word in the input
// string as the corresponding value.
// Matching is case insensitive, e.g. "Orange" and "orange" is considered the
// same word.
func WordIndexAll(s string) map[string][]int {
	// TODO: implement me
	// HINT: You may find the `strings.Index` and `strings.ToLower` functions helpful

	var s_split = strings.Fields(s)

	IndexMap := make(map[string][]int)

	for _, substring := range s_split {

		if IndexMap[substring] == 0 && substring == strings.ToLower(substring) {
			IndexMap[substring] = append(IndexMap[substring], strings.Index(s, substring)) 
		}
		
	}

	return IndexMap

}
