package challenge02

import (
	"strings"
)

func Manufacture(gifts []string, materials string) []string {
	var result []string
	for _, gift := range gifts {
		canBeManufactured := true
		for _, char := range gift {
			if !strings.Contains(materials, string(char)) {
				canBeManufactured = false
				break
			}
		}
		if canBeManufactured {
			result = append(result, gift)
		}
	}

	return result
}
func notIncludes(materials []string, char string) bool {
	for _, m := range materials {
		if m == char {
			return false
		}
	}
	return true
}
