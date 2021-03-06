package array

func removeElement(arr []int, val int) []int {
	for i := 0; i < len(arr); {
		if arr[i] == val {
			arr = append(arr[:i], arr[i+1:]...)
		} else {
			i++
		}
	}
	return arr
}
