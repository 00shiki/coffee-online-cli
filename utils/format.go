package utils

import "strconv"

func PriceFormat(price float64) string {
	s := strconv.FormatFloat(price, 'f', 2, 64)
	for i := len(s) - 6; i > 0; i -= 3 {
		s = s[:i] + "," + s[i:]
	}
	return s
}
