package main

import (
	"fmt"
	"strings"
)

func notIncludes(materials []string, char string) bool {
	for _, m := range materials {
		if m == char {
			return false
		}
	}
	return true
}
func manufacture(gifts []string, materials string) []string {
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

func main() {
	gifts := []string{"tren", "oso", "pelota"}
	materials := "tronesa"
	result := manufacture(gifts, materials)
	fmt.Println(result) // ["tren", "oso"]
	// "tren" SÍ porque sus letras están en "tronesa"
	// "oso" SÍ porque sus letras están en "tronesa"
	// "pelota" NO porque sus letras NO están en "tronesa"

	gifts = []string{"juego", "puzzle"}
	materials = "jlepuz"
	result = manufacture(gifts, materials) // ["puzzle"]
	fmt.Println(result)                    //

	gifts = []string{"libro", "ps5"}
	materials = "psli"
	result = manufacture(gifts, materials) // []
	fmt.Println(result)                    //
}
