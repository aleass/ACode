package main

import "fmt"

func main() {
	var params = []int{2, 7, 11, 15}
	var taget = 9
	res := twoSum(params, taget)
	fmt.Println(params[res[0]], params[res[1]], taget)

	params = []int{3, 2, 4}
	taget = 6
	res = twoSum2(params, taget)
	fmt.Println(params[res[0]], params[res[1]], taget)

	params = []int{3, 2, 4}
	taget = 6
	res = twoSum2(params, taget)
	fmt.Println(params[res[0]], params[res[1]], taget, res)
}

func twoSum2(nums []int, target int) []int {
	for i, num := range nums {
		for j := i + 1; j < len(nums); j++ {
			if num+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, 100)
	m[nums[0]] = 0
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if i2, ok := m[target-num]; ok {
			return []int{i, i2}
		}
		m[num] = i
	}
	return nil
}
