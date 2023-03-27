package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("aca"), 3)
	fmt.Println(lengthOfLongestSubstring("dvdf"), 3)
	fmt.Println(lengthOfLongestSubstring("abba"), 2)
	fmt.Println(lengthOfLongestSubstring("tmmzuxt"), 5)
}

/*
abcdacef
发现第一个重复a
abcda -> index:0 i:5

	s = s[index+1:] 等于 a｜bcdacef
	i -= index  等于 5 - 0 = 5 字符串发生裁剪，需要把i移到对应的新位置
*/
func lengthOfLongestSubstring(s string) int {
	if l := len(s); l < 2 { //直接返回
		return l
	}
	var i, res = 1, 1 //i 字符串的下标，res返回的值，赋值1 减少一次循环
	for {
		if l := len(s); res >= l || i >= l { //终止循环：res长度大于剩余长度  或者 i大于剩余长度
			if res < l {
				res = l
			}
			return res
		}
		v := s[i : i+1]
		if index := strings.Index(s[:i], v); index != -1 {
			if l := len(s[:i]); res < l { //获取当前没有重复的长度
				res = l
			}
			s = s[index+1:] //从重复处裁切
			i -= index      //重置index  字符串发生切割，需要重置index
		} else {
			i++ //index偏移
		}
	}
}

func lengthOfLongestSubstring3(s string) int {
	strMap := make(map[byte]int)
	var start, m, l int
	for i := 0; i < len(s); i++ {
		l++
		if index, ok := strMap[s[i]]; ok {
			if index >= start {
				start = index + 1
				l = i - index
			}
		}
		strMap[s[i]] = i
		if l > m {
			m = l
		}
	}
	return m
}
func lengthOfLongestSubstring1(s string) int {
	if l := len(s); l < 2 {
		return l
	}
	sB := []byte(s)
	var res int
	for {
	start:
		for i, v := range sB {
			if index := bytes.IndexByte(sB[:i], v); index != -1 {
				if res < len(sB[:i]) {
					res = len(sB[:i])
				}
				sB = sB[index+1:]
				//println(string(sB))
				goto start
			}
		}
		break
	}

	if res < len(sB) {
		res = len(sB)
	}
	return res
}
