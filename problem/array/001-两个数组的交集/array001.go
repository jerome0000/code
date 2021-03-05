package array

import "sort"

func intersectV1(num1 []int, num2 []int) []int {
	result := make([]int, 0)

	m := make(map[int]int)
	for _, item := range num1 {
		m[item] += 1
	}

	for _, item := range num2 {
		if m[item] != 0 {
			result = append(result, item)
		}
	}

	return result
}

func intersectV2(num1 []int, num2 []int) []int {
	result := make([]int, 0)

	sort.Ints(num1)
	sort.Ints(num2)

	i, j := 0, 0
	for i < len(num1) && j < len(num2) {
		if num1[i] > num2[j] {
			j++
		} else if num1[i] < num2[j] {
			i++
		} else {
			result = append(result, num1[i])
			i++
			j++
		}
	}

	return result
}
