package sort001

func quickSort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	f := arr[0]

	b := make([]int, 0)
	s := make([]int, 0)
	m := make([]int, 0)

	for i := 0; i < len(arr); i++ {
		if arr[i] > f {
			b = append(b, arr[i])
		} else if arr[i] == f {
			m = append(m, arr[i])
		} else {
			s = append(s, arr[i])
		}
	}

	// fmt.Println("big", b)
	// fmt.Println("small", s)
	// fmt.Println("middle", m)

	b, s = quickSort(b), quickSort(s)

	return append(append(s, m...), b...)
}
