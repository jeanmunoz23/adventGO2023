package challenge10

import (
	"strings"
)

func CreateChristmasTree(ornaments string, height int) string {
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
