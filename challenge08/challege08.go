package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func organizeGifts(gifts string) string {
	const boxSize = 10
	const palletSize = 5 * boxSize

	var organizedGifts string
	remainings := gifts

	for len(remainings) > 0 {
		re := regexp.MustCompile("[A-Za-z]")
		nextTypeIndex := re.FindStringIndex(remainings)
		nextTypeSymbol := remainings[nextTypeIndex[0]:nextTypeIndex[1]]
		nextTypeAmountStr := remainings[:nextTypeIndex[0]]
		remainings = remainings[nextTypeIndex[1]:]

		nextTypeAmount, _ := strconv.Atoi(nextTypeAmountStr)

		numberOfPallets := nextTypeAmount / palletSize
		giftOfPallets := numberOfPallets * palletSize
		giftCount := float64(nextTypeAmount-giftOfPallets) / float64(boxSize)
		numberOfBoxes := int(giftCount)
		giftOfBoxes := numberOfBoxes * boxSize

		numberOfSpares := nextTypeAmount - giftOfBoxes - giftOfPallets

		if numberOfPallets > 0 {
			organizedGifts += repeatString(fmt.Sprintf("[%s]", nextTypeSymbol), numberOfPallets)
		}
		if numberOfBoxes > 0 {
			organizedGifts += repeatString(fmt.Sprintf("{%s}", nextTypeSymbol), numberOfBoxes)
		}
		if numberOfSpares > 0 {
			organizedGifts += fmt.Sprintf("(%s)", repeatString(nextTypeSymbol, numberOfSpares))
		}
	}

	return organizedGifts
}

func repeatString(s string, n int) string {
	var result string
	for i := 0; i < n; i++ {
		result += s
	}
	return result
}
func main() {
	result := organizeGifts(`76a11b`)
	fmt.Println(result) // "[a]{a}{a}(aaaaaa){b}(b)"

}
