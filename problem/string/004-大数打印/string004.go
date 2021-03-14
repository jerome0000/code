package string

func printNum(num int) []int {
	if num == 0 {
		return []int{0}
	}
	max := 10 ^ num
	result := make([]int, 0)
	for i := 1; i < max-1; i++ {
		result = append(result, i)
	}
	return result
}
