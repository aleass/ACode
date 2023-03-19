package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("babad"))
}
func longestPalindrome(s string) string {
	var l = len(s)
	if l <= 1 {
		return s
	}
	// abcdbac
	var max, res = 0, s[:1]
	for i := 1; i < l; i++ {
		if s[i] == s[i-1] {
			max = 2
			res = s[i-1 : i+1]
			continue
		}
		var ok bool = true
		j, k := i-1, i+1
		for ; j > 0 && k < l-2; j, k = j+1, k+1 {
			if s[j] != s[k] {
				ok = false
				break
			}
		}
		if (k-j) > max && ok {
			res = s[j : k+1]
		}
	}
	return res
}
