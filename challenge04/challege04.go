package main

import (
	"fmt"
	"regexp"
)

func decode(message string) string {
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

func main() {
	result := decode("hola (odnum)")
	fmt.Println(result) // hola mundo

	result = decode("(olleh) (dlrow)!")
	fmt.Println(result) // hello world!

	result = decode("sa(u(cla)atn)s")
	fmt.Println(result) // santaclaus

	// Paso a paso:
	// 1. Invertimos el anidado -> sa(ualcatn)s
	// 2. Invertimos el que queda -> santaclaus
}
