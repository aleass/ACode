package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(reverse(1234))
}
func reverse(x int) int {
	s := strconv.Itoa(x)
	isO := false
	if s[0] == '-' {
		s = s[1:]
		isO = true
	}
	s1 := make([]byte, len(s))
	for i, v := range s {
		s1[len(s)-i-1] = byte(v)
	}
	var s2, _ = strconv.Atoi(string(s1))
	if isO {
		s2 = -s2
	}
	if s2 > (1<<32/2-1) || s2 < (-1<<32/2) {
		return 0
	}
	return s2
}
