package challenge13

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func CalculateTime(deliveries []string) string {
	const limit = 7 * 60 * 60 // 7 horas en segundos
	var totalDuration int

	for _, delivery := range deliveries {
		parts := strings.Split(delivery, ":")
		hours, _ := strconv.Atoi(parts[0])
		minutes, _ := strconv.Atoi(parts[1])
		seconds, _ := strconv.Atoi(parts[2])

		totalDuration += hours*3600 + minutes*60 + seconds
	}

	remainingTime := totalDuration - limit
	sign := ""
	if remainingTime < 0 {
		sign = "-"
	}
	remainingTime = int(math.Abs(float64(remainingTime)))

	hours := remainingTime / 3600
	minutes := (remainingTime % 3600) / 60
	seconds := remainingTime % 60

	hh := fmt.Sprintf("%s%02d", sign, hours)
	mm := fmt.Sprintf("%02d", minutes)
	ss := fmt.Sprintf("%02d", seconds)

	return fmt.Sprintf("%s:%s:%s", hh, mm, ss)
}
