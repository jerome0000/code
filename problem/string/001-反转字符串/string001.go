package string

func reverseString(s []byte) []byte {
	r, l := 0, len(s)-1
	for r < l {
		s[r], s[l] = s[l], s[r]
		r++
		l--
	}
	return s
}
