package common

import "strconv"

func ToInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}
