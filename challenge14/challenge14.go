package challenge14

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaxGifts(houses []int) int {
	if len(houses) == 0 {
		return 0
	}
	if len(houses) == 1 {
		return houses[0]
	}

	dp := make([]int, len(houses))
	dp[0] = houses[0]
	dp[1] = max(houses[0], houses[1])

	for i := 2; i < len(houses); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+houses[i])
	}

	return dp[len(houses)-1]
}
