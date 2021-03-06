package array

func thirdSum(arr []int, target int) []int {
	if len(arr) < 3 {
		return []int{}
	}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			for k := j + 1; k < len(arr); k++ {
				if arr[i]+arr[j]+arr[k] == target {
					return []int{arr[i], arr[j], arr[k]}
				}
			}
		}
	}
	return []int{}
}
