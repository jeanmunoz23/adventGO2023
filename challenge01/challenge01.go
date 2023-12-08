package main

import "fmt"

func findFirstRepeated(gifts []int) int {
	var seenGifts = make(map[int]bool)

	for _, gift := range gifts {
		if seenGifts[gift] {
			return gift
		}
		seenGifts[gift] = true
	}

	return -1
}
func main() {
	giftIds := []int{2, 1, 3, 5, 3, 2}
	result := findFirstRepeated(giftIds)
	fmt.Println(result) // 3
	// Aunque el 2 y el 3 se repiten
	// el 3 aparece primero por segunda vez

	giftIds = []int{1, 2, 3, 4}
	result = findFirstRepeated(giftIds)
	fmt.Println(result) // -1
	// Es -1 ya que no se repite ningún número

	giftIds = []int{5, 1, 5, 1}
	result = findFirstRepeated(giftIds)
	fmt.Println(result) // 5
}
