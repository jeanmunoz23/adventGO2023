package challenge01

func FindFirstRepeated(gifts []int) int {
	var seenGifts = make(map[int]bool)

	for _, gift := range gifts {
		if seenGifts[gift] {
			return gift
		}
		seenGifts[gift] = true
	}

	return -1
}
