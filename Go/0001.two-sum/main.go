package main

import (
	"fmt"
	"sort"
)

func main() {
	var params = []int{2, 7, 11, 15}
	var taget = 9
	res := twoSum(params, taget)
	fmt.Println(params[res[0]], params[res[1]], taget)
	res = twoSum2(params, taget)
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
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	sort.Ints(sorted)
	i, j := 0, len(sorted)-1
	for i < j {
		sum := sorted[i] + sorted[j]
		if sum < target {
			i++
		} else if sum > target {
			j--
		} else {
			break
		}
	}
	indices := make([]int, 2)
	var flag = false
	for k, num := range nums {
		if num == sorted[i] && flag == false {
			indices[0] = k
			flag = true
		} else if num == sorted[j] {
			indices[1] = k
		}
		if i == -1 && j == -1 {
			break
		}
	}
	return indices
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
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
