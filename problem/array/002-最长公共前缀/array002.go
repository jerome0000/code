package array

func longestCommonPrefix(arr []string) string {
	if len(arr) == 0 {
		return ""
	}

	if len(arr) == 1 {
		return arr[0]
	}

	tmp := arr[0]
	for i := 1; i <= len(tmp); i++ {
		str := tmp[:i]
		for _, item := range arr {
			if len(item) >= i && str == item[:i] {
				continue
			} else {
				return tmp[:i-1]
			}
		}
	}
	return tmp
}
