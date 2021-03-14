package string

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	str := ""
	// to string and number
	for i := 0; i < len(s); i++ {
		if unicode.IsNumber(rune(s[i])) || unicode.IsDigit(rune(s[i])) {
			str += string(s[i])
		}
	}
	// 遍历
	n := len(str)
	str = strings.ToLower(str)
	for i := 0; i < n/2; i++ {
		if str[i] != str[n-1-i] {
			return false
		}
	}
	return true
}
