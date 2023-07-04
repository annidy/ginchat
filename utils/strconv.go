package utils

import "strconv"

func Atou(s string) uint {
	i, _ := strconv.Atoi(s)
	return uint(i)
}
