package main

import "fmt"

func main() {
	fmt.Println(len(letterCombinations("9")))
}

//144
//48
//16
//4

var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {
	var l = len(digits)
	if l == 0 {
		return []string{}
	}
	size := 4
	if l == 2 {
		size = 16
	} else if l == 3 {
		size = 48
	} else if l == 4 {
		size = 144
	}
	combinations := make([]string, 0, size)
	backtrack(digits, 0, "", combinations)
	return combinations
}

func backtrack(digits string, index int, combination string, combinations []string) {
	if index == len(digits) {
		combinations = append(combinations, combination)
	} else {
		digit := string(digits[index])
		letters := phoneMap[digit]
		lettersCount := len(letters)
		for i := 0; i < lettersCount; i++ {
			backtrack(digits, index+1, combination+string(letters[i]), combinations)
		}
	}
}
