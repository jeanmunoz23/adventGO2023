package main

import (
	"fmt"
	"strings"
)

func cyberReindeer(road string, time int) []string {
	moves := []string{road}
	sledIndex := 0
	barrierSymbol := "."

	for currentTime := 1; currentTime < time; currentTime++ {
		if currentTime == 5 {
			road = strings.ReplaceAll(road, "|", "*")
		}

		newRoad := strings.Replace(road, "S.", barrierSymbol+"S", 1)
		newRoad = strings.Replace(newRoad, "S*", barrierSymbol+"S", 1)

		if newRoad != road {
			sledIndex++
			barrierSymbol = string(road[sledIndex])
		}

		road = newRoad
		moves = append(moves, road)
	}

	return moves
}

func main() {
	road := "S..|...|.."
	time := 10 // unidades de tiempo

	for _, result := range cyberReindeer(road, time) {
		fmt.Println(result)
	}
}
