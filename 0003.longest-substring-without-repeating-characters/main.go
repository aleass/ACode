package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("aca"), 3)
	//fmt.Println(lengthOfLongestSubstring("dvdf"), 3)
	//fmt.Println(lengthOfLongestSubstring("abba"), 2)
	//fmt.Println(lengthOfLongestSubstring("tmmzuxt"), 5)
}

func lengthOfLongestSubstring(s string) int {
	aa := make(map[byte]int)
	start := 0
	m := 0
	l := 0
	size := len(s)
	for i := 0; i < size; i++ {
		l++
		if value, ok := aa[s[i]]; ok {
			if value >= start {
				start = value + 1
				l = i - value
			}
		}
		aa[s[i]] = i
		if l > m {
			m = l
		}
	}
	return m
}

func lengthOfLongestSubstring1(s string) int {
	var l = len(s)
	if l <= 1 {
		return l
	}
	var max int
	var rb = make(map[byte]struct{}, l)
	for i := 0; i < l; i++ {
		rb = map[byte]struct{}{}
		rb[s[i]] = struct{}{}
		j := i + 1
		if i == 2 {
			println()
		}
		for ; j < l; j++ {
			if _, ok := rb[s[j]]; ok {
				break
			}
			rb[s[j]] = struct{}{}
		}

		if _max := j - i; max < _max {
			max = _max
		}
	}
	return max
}
