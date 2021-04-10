package sliding_window

func maxSlidingWindow(arr []int, k int) []int {
	l := len(arr)

	if l*k == 0 {
		return []int{0}
	}

	result := make([]int, 0)

	for i := 0; i < l-k+1; i++ {
		max := 0 << 31
		for j := i; j < i+k; j++ {
			max = getMAX(max, arr[j])
		}
		result = append(result, max)
	}

	return result

}

func getMAX(a, b int) int {
	if a > b {
		return a
	}
	return b
}
