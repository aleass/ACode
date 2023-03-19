package main

import (
	"sort"
	"testing"
)

func BenchmarkTwo(b *testing.B) {
	var params = []int{3, 2, 4, 234, 234, 23, 24, 54, 6, 75, 7, 54, 2, 323, 6, 87, 34, 23, 1, 2, 7, 11, 15}
	//var taget = 9
	b.Run("twoSum3", func(b *testing.B) {
		sorted := make([]int, len(params))
		copy(sorted, params)
		for i := 0; i < b.N; i++ {
			sort.Ints(sorted)
			//twoSum3(params, taget)
		}
	})
	b.Run("twoSum2", func(b *testing.B) {
		//twoSum2(params, taget)
		sorted := make([][2]int, len(params))
		for j, v := range params {
			sorted[j] = [2]int{v, j}
		}
		for i := 0; i < b.N; i++ {
			sort.Slice(sorted, func(i, j int) bool {
				return sorted[i][0] < sorted[j][0]
			})
		}
	})
}
