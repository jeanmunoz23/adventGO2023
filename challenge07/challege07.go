package challenge07

import (
	"strings"
)

func DrawGift(size int, symbol string) string {
	if size < 2 {
		return strings.Repeat("#", size) + "\n"
	}

	var lines []string

	// Primero, dibujar la parte superior de la caja de regalo.
	lines = append(lines, strings.Repeat(" ", size-1)+strings.Repeat("#", size))

	// Dibujar el resto de la parte superior de la caja de regalo.
	for i := 1; i < size-1; i++ {
		line := strings.Repeat(" ", size-i-1) + "#" + strings.Repeat(symbol, size-2) + "#" + strings.Repeat(symbol, i-1) + "#"
		lines = append(lines, line)
	}

	// Dibujar la parte media de la caja de regalo.
	gift := strings.Repeat("#", size) + strings.Repeat(symbol, size-2) + "#"
	lines = append(lines, gift)

	// Invertir la parte superior para crear la otra mitad de la caja de regalo.
	topPart := make([]string, len(lines)-1)
	copy(topPart, lines[:len(lines)-1])
	reverseTopPart := make([]string, len(topPart))
	for i := 0; i < len(topPart); i++ {
		reverseTopPart[i] = strings.ReplaceAll(topPart[len(topPart)-1-i], " ", "")
	}
	lines = append(lines, reverseTopPart...)

	return strings.Join(lines, "\n") + "\n"
}

func reverse(lines []string) []string {
	newLines := make([]string, 0, len(lines))
	for i := len(lines) - 1; i >= 0; i-- {
		newLines = append(newLines, lines[i])
	}
	return newLines
}
