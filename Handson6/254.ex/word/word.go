// Package word is a customized package to split string in the Word
package word

import "strings"

// Function Wordcount is a customized function to count the occurence of a word in statement
//using string.Fields method
func Wordcount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Function Wrdcount is a customized function to count the occurence of a word in statement
//using string.Split method
func Wrdcount(s string) map[string]int {
	xs := strings.Split(s, " ")
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Function Count is a customized function to count the number of a word in statement
func Count(s string) int {
	xs := strings.Fields(s)
	return len(xs)
}
