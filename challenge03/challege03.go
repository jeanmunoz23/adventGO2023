package main

import (
	"fmt"
)

func findNaughtyStep(original string, modified string) string {
	var larger, shorter string
	if len(original) >= len(modified) {
		larger, shorter = original, modified
	} else {
		larger, shorter = modified, original
	}
	for i, char := range larger {
		if i < len(shorter) && char != rune(shorter[i]) {
			return string(char)
		}
	}

	if len(larger) > len(shorter) {
		return string(larger[len(shorter)])
	}

	return ""
}

func main() {
	original := "abcd"
	modified := "abcde"
	result := findNaughtyStep(original, modified) // "e"
	fmt.Println(result)

	original = "stepfor"
	modified = "stepor"
	result = findNaughtyStep(original, modified) // "f"
	fmt.Println(result)

	original = "abcde"
	modified = "abcde"
	result = findNaughtyStep(original, modified) // ""
	fmt.Println(result)
}
