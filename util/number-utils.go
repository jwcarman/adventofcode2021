package util

import "strconv"

func StringsToInts(strings []string) []int {
	var integers []int
	for _, s := range strings {
		i, _ := strconv.Atoi(s)
		integers = append(integers, i)
	}
	return integers
}
