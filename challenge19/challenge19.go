package challenge19

import (
	"fmt"
)

func revealSabotage(store [][]string) [][]string {
	rows := len(store)
	cols := len(store[0])

	result := make([][]string, rows)
	for i := 0; i < rows; i++ {
		result[i] = make([]string, cols)
		copy(result[i], store[i])
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			celda := store[i][j]
			var prev, actual, next []string

			if i > 0 {
				prev = store[i-1]
			}
			actual = store[i]
			if i < rows-1 {
				next = store[i+1]
			}

			c1 := boolToInt(i > 0 && j > 0 && prev[j-1] == "*") + boolToInt(i > 0 && prev[j] == "*")
			c2 := boolToInt(i > 0 && j < cols-1 && prev[j+1] == "*") + boolToInt(j > 0 && actual[j-1] == "*")
			c3 := boolToInt(j < cols-1 && actual[j+1] == "*") + boolToInt(i < rows-1 && j > 0 && next[j-1] == "*")
			c4 := boolToInt(i < rows-1 && next[j] == "*") + boolToInt(i < rows-1 && j < cols-1 && next[j+1] == "*")
			calc := c1 + c2 + c3 + c4

			result[i][j] = []string{celda, fmt.Sprintf("%d", calc)}[boolToInt(calc > 0)*boolToInt(celda == " ")]
		}
	}

	return result
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func main() {
	store := [][]string{
		{"*", " ", " ", " "},
		{" ", " ", "*", " "},
		{" ", " ", " ", " "},
		{"*", " ", " ", " "},
	}

	result := revealSabotage(store)

	for _, row := range result {
		fmt.Println(row)
	}
}

/*
const store = [
  ["*", " ", " ", " "],
  [" ", " ", "*", " "],
  [" ", " ", " ", " "],
  ["*", " ", " ", " "]
]

console.log(revealSabotage(store))
/* Debería mostrar:
[
    ["*", "2", "1", "1"],
    ["1", "2", "*", "1"],
    ["1", "2", "1", "1"],
    ["*", "1", " ", " "]
]*/
