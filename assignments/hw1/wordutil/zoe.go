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

	//Cannot iterate in a string, must make a slice first

	map1 := make(map[string]int)

	for _, word := range s {
		v := strings.ToLower(word)
		_, ok := map1[v]
		//If statement needs brackets {}
		if !ok {
			map1[v] = map1[v] + strings.Index(s, v)
		}
	}

	return map1

	// TODO: implement me
	// HINT: You may find the strings.Index and strings.ToLower functions helpful
}
