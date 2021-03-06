package array

func reverse(arr []int, k int) []int {
	result := make([]int, 0)
	result = append(result, arr[len(arr)-k:]...)
	result = append(result, arr[:len(arr)-k]...)

	return result
}

