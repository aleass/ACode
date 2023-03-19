package main

import "fmt"

func main() {
	//fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}), 49)
	fmt.Println(maxArea([]int{1, 2, 3, 4}), 4)
}

func maxArea(height []int) int {
	var l = len(height)
	if l < 3 {
		if height[0] >= height[1] {
			return height[1]
		}
		return height[0]
	}
	//l == width
	l--
	var max, _max int
	h, t := 0, l
	//减少一次循环
	if vh, vt := height[h], height[t]; vh >= vt {
		max = vt * l
		t--
	} else {
		max = vh * l
		h++
	}

	//因为少了一次循环，需要长度-1
	for l--; ; l-- {
		if vh, vt := height[h], height[t]; vh >= vt {
			_max = vt * l
			t--
			if height[t] < vt {
				t--
				l--
			}
		} else {
			_max = vh * l
			h++
			if height[h] < vh {
				h++
				l--
			}
		}
		if _max > max {
			max = _max
		}
		if l == 1 && t >= h {
			return max
		}
	}
}

func maxArea2(height []int) int {
	var l = len(height)
	if l < 3 {
		if height[0] >= height[1] {
			return height[1]
		}
		return height[0]
	}

	l--
	var max, _max int
	for h, t := 0, l; ; l = t - h {
		if vh, vt := height[h], height[t]; vh >= vt {
			_max = vt * l
			t--
			if height[t] < vt {
				t--
				l--
			}
		} else {
			_max = vh * l
			h++
			if height[h] < vh {
				h++
				l--
			}
		}
		if _max > max {
			max = _max
		}
		if l == 1 && t >= h {
			return max
		}
	}
}
