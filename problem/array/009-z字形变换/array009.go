package array

import (
	"strings"
)

func z(s string, rows int) string {
	if rows == 1 {
		return s
	}
	b := rows - 1
	arr := make([]string, rows)
	for k, v := range s {
		if (k/b)%2 == 0 {
			arr[k%b] += string(v)
		} else {
			arr[b-(k%b)] += string(v)
		}
	}

	return strings.Join(arr, "")
}
