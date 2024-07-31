package utils

import (
	"strconv"
	"strings"
)

func PriceFormat(price float64) string {
	s := strconv.FormatFloat(price, 'f', 2, 64)
	s = strings.Replace(s, ".", ",", 1)
	for i := len(s) - 6; i > 0; i -= 3 {
		s = s[:i] + "." + s[i:]
	}
	return s
}
