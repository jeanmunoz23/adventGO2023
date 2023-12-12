package challenge11

import (
	"fmt"
	"strings"
)

// En go el ZeroValue de un array es [] entonces cuando de null/nil devolvera un []
func getIndexsForPalindromeVGO(word string) []int {
	if word == reverseString(word) {
		return []int{}
	}

	for i := 0; i < len(word)-1; i++ {
		for j := i + 1; j < len(word); j++ {
			changed := []rune(word)
			changed[i], changed[j] = changed[j], changed[i]
			if isPalindrome(string(changed)) {
				return []int{i, j}
			}
		}
	}

	return nil
}

// simulando lo mismo que da en nodejs
func GetIndexsForPalindrome(word string) string {
	if word == reverseString(word) {
		return "[]"
	}

	for i := 0; i < len(word)-1; i++ {
		for j := i + 1; j < len(word); j++ {
			changed := []rune(word)
			changed[i], changed[j] = changed[j], changed[i]
			if isPalindrome(string(changed)) {
				return fmt.Sprintf("[%d,%d]", i, j)
			}
		}
	}
	return "null"
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	runes := []rune(s)
	for i := 0; i < len(runes)/2; i++ {
		if runes[i] != runes[len(runes)-1-i] {
			return false
		}
	}
	return true
}
