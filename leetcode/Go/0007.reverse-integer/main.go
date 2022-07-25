package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(reverse(123456))
}
func reverse(x int) int {
	var minus bool
	if x < 0 {
		minus = true
		x = -x
	}
	var ids = 0
	var m2 = 0
	var count = len(strconv.Itoa(x)) - 1
	if count > 10 {
		return 0
	}
	for x > 0 {
		var m = math.Pow10(count)
		var y = x / int(m)
		x -= y * int(m)
		if m2 == 0 {
			ids += y
		} else {
			ids += y * int(math.Pow10(m2))
		}
		m2++
		count--
	}
	if minus {
		ids = -ids
	}
	if ids > (1<<32/2-1) || ids < (-1<<32/2) {
		return 0
	}
	return ids
}
