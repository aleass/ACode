package main

func main() {

}
func twoSum(nums []int, target int) []int {
	var list = make(map[int]int)
	for i, num := range nums {
		if k, ok := list[num]; ok {
			return []int{i, k}
		}
		list[target-num] = i
	}
	return nil
}
