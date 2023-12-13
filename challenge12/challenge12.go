package challenge12

import (
	"regexp"
	"strings"
	"unicode"
)

func CheckIsValidCopy(original, copy string) bool {
	originalRunes := []rune(original)
	copyRunes := []rune(copy)
	originalArray := strings.Split(original, " ")
	copyArray := strings.Split(copy, " ")
	i := -1
	res := true

	for _, o := range originalRunes {
		i++

		if len(originalArray[0]) != len(copyArray[0]) {
			return false
		}
		if unicode.ToLower(o) == copyRunes[i] {
			continue
		}

		symbols := []rune{o, unicode.ToLower(o), '#', '+', ':', '.', ' '}
		regx := regexp.MustCompile(`([A-Z])|([a-z])|(\#)|(\+)|(\:)|(\.)|(\s)`)
		matches := regx.FindStringSubmatch(string(o))
		var symbolIndex int
		if len(matches) > 0 {
			symbolIndex = strings.LastIndexAny(string(symbols), matches[0])
		} else {
			symbolIndex = 1
		}

		arr := strings.ContainsAny(string(symbols[symbolIndex:]), string(copyRunes[i]))

		return arr
	}
	return res
}
