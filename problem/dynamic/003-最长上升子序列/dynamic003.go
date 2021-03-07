package dynamic

func lengthOfLIS(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	// dp[i] 表示包含第i个元素(i=0,1,...n-1)的最长上升子序列的长度, 至少为1
	dp := make([]int, length)
	for i, _ := range dp {
		dp[i] = 1
	}

	for i := 0; i < length; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
	}

	var result int
	for _, v := range dp {
		if v > result {
			result = v
		}
	}
	return result
}
