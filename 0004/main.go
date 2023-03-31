package main

import "fmt"

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
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
	for i := 0; l1 > 0 || l2 > 0; i++ {
		if l1 == 0 {
			res = float64(nums2[0])
			nums2 = nums2[1:]
			l2--

		} else if l2 == 0 {
			res = float64(nums1[0])
			nums1 = nums1[1:]
			l1--

		} else if nums1[0] < nums2[0] {
			res = float64(nums1[0])
			nums1 = nums1[1:]
			l1--
		} else {
			res = float64(nums2[0])
			nums2 = nums2[1:]
			l2--
		}
		if i >= medianNum {
			if !double {
				return res
			}
			if res2 != 0 {
				return (res + res2) / 2
			}
			res2 = res
		}
	}
	return 0
}
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	if l1 == 1 && l2 == 1 {
		return float64(nums1[0]+nums2[0]) / 2
	}
	var num = (l1+l2)/2 + 1
	var isSingle bool
	if (l1+l2)&1 == 1 {
		isSingle = true
	}
	var v1, v2 int
	for ; num != 0; num-- {
		if len(nums1) == 0 {
			if num < 2 {
				v2 = v1
			} else {
				v2 = nums2[num-2]
			}
			v1 = nums2[num-1]
			break
		}
		if len(nums2) == 0 {
			if num < 2 {
				v2 = v1
			} else {
				v2 = nums1[num-2]
			}
			v1 = nums1[num-1]
			break
		}
		v2 = v1
		if nums1[0] < nums2[0] {
			v1 = nums1[0]
			nums1 = nums1[1:]
		} else {
			v1 = nums2[0]
			nums2 = nums2[1:]
		}
	}
	if !isSingle {
		return float64(v1+v2) / 2
	}
	return float64(v1)
}
