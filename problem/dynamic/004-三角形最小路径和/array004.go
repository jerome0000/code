package array

func minimumTotal(triangle [][]int) int {
	h := len(triangle)

	dp := make([][]int, h)
	for i := range dp {
		dp[i] = make([]int, len(triangle[i]))
	}

	for i := h - 1; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			if i == h-1 {
				dp[i][j] = triangle[i][j]
			} else {
				dp[i][j] = min(dp[i+1][j], dp[i+1][j+1]) + triangle[i][j]
			}
		}
	}
	return dp[0][0]
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
