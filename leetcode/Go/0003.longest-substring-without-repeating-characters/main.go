package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"), 3)
	fmt.Println(lengthOfLongestSubstring("dvdf"), 3)
}

func lengthOfLongestSubstring(s string) int {
	var l = len(s)
	if l <= 1 {
		return l
	}
	var h, t, max int
	var m = make(map[byte]int, l)
	for i := 0; i < l; i++ {
		if res, ok := m[s[i]]; ok {
			if r := t - h; r > max {
				max = r
			}
			h = res + 1
			t = res + 1
		} else {
			t++
		}
		m[s[i]] = i
	}
	if r := t - h; r > max {
		max = r
	}
	return max
}
