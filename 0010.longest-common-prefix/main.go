package main

import "fmt"

func main() {
	fmt.Println("'" + longestCommonPrefix([]string{"abab", "aba", ""}) + "'")
	fmt.Println("'" + longestCommonPrefix([]string{"cir", "car"}) + "'")
}
func longestCommonPrefix(strs []string) string {
	var l = len(strs)
	if l < 2 {
		return strs[0]
	}

	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < l; j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

func longestCommonPrefix1(strs []string) string {
	if len(strs) < 2 {
		return strs[0]
	}
	var l = len(strs[0])
	for j := 1; j < len(strs); j++ {
		if _l := len(strs[j]); _l < l {
			l = _l
		}
	}
	var res = make([]byte, l)
	i := 0
	for ; i < l; i++ {
		var tmp = strs[0][i]
		for j := 1; j < len(strs); j++ {
			if len(strs[j]) == 0 {
				return string(res[:i])
			}
			if strs[j][i] != tmp {
				return string(res[:i])
			}
		}
		res[i] = tmp
	}
	return string(res[:i])
}
