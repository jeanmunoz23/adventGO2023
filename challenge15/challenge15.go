package challenge15

import "strings"

func autonomousDrive(store []string, movements []string) []string {
	posY := 0
	posX := 0
	// Buscar la primera posición
	for i := 0; i < len(store); i++ {
		robotPos := strings.Index(store[i], "!")
		if robotPos != -1 {
			posX = robotPos
			posY = i
			store[posY] = strings.Replace(store[posY], "!", ".", 1)
		}
	}

	for _, move := range movements {
		nextX := posX + map[bool]int{move == "R": 1, move == "L": -1}[true]
		nextY := posY - map[bool]int{move == "U": 1, move == "D": -1}[true]
		if nextY >= 0 && nextY < len(store) && nextX >= 0 && nextX < len(store[nextY]) && string(store[nextY][nextX]) == "." {
			posX = nextX
			posY = nextY
		}
	}

	store[posY] = store[posY][:posX] + "!" + store[posY][posX+1:]
	return store
}
