package main

import "fmt"

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var l1, l2 = len(nums1), len(nums2)
	l := l1 + l2
	if l == 0 {
		return 0
	}
	var medianNum = l / 2
	var double bool
	if l%2 == 0 {
		medianNum--
		double = true
	}
	var res, res2 float64
	for i := 0; l1 != 0 && l2 != 0; i++ {
		if i == medianNum+1 {
			if !double {
				return res
			}
			res2 = res
		}
		if i == medianNum+2 {
			return (res + res2) / 2
		}
		if l1 == 0 {
			nums2 = nums2[1:]
			res = float64(nums2[0])
			l2--
			continue
		}
		if l2 == 0 {
			nums1 = nums1[1:]
			res = float64(nums1[0])
			l1--
			continue
		}
		if nums1[0] < nums2[0] {
			res = float64(nums1[0])
			nums1 = nums1[1:]
			l1--
		} else {
			res = float64(nums1[0])
			nums2 = nums2[1:]
			l2--
		}
	}
	return -1
}
