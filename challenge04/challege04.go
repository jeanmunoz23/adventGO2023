package challenge04

import (
	"regexp"
)

func Decode(message string) string {
	regex := regexp.MustCompile(`\(([^()]*)\)`)

	for regex.MatchString(message) {
		message = regex.ReplaceAllStringFunc(message, func(match string) string {
			text := regex.FindStringSubmatch(match)[1]
			reversed := reverseString(text)
			return reversed
		})
	}

	return message
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
