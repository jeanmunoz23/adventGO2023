package main

import (
	"fmt"
	"strings"
)

func createChristmasTree(ornaments string, height int) string {
	tree := ""
	ornament := 0
	for i := 0; i < height; i++ {
		ornamentsInRow := i + 1
		row := strings.Repeat(" ", height-ornamentsInRow)
		for j := 0; j < ornamentsInRow; j++ {
			row += string(ornaments[ornament]) + " "
			ornament++
			if ornament == len(ornaments) {
				ornament = 0
			}
		}
		tree += strings.TrimRight(row, " ") + "\n"
	}
	tree += strings.Repeat(" ", height-1) + "|\n"
	return tree
}

func main() {
	// Ejemplos de uso:
	tree := createChristmasTree("123", 4)
	fmt.Println(tree)

	tree = createChristmasTree("*@o", 3)
	fmt.Println(tree)
	// Ejemplos de uso:
	tree = createChristmasTree("xo", 4)
	fmt.Println(tree)
}
