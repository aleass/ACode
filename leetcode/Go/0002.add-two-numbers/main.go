package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//4 ms	4.2 MB 结合
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var res = new(ListNode)
	tempL := res
	var temps *ListNode
	v1, v2 := l1.Val, l2.Val
	var c1, c2 int
	var temp, v3 int
	for l1.Next != nil || l2.Next != nil {
		v3 = v1 + v2 + temp
		v1, v2, temp = 0, 0, 0
		if v3 > 9 {
			v3 -= 10
			temp = 1
		}
		if temp == 0 {
			if l1.Next == nil {
				tempL.Val = v3
				tempL.Next = l2.Next
				return res
			}
			if l2.Next == nil {
				tempL.Val = v3
				tempL.Next = l1.Next
				return res
			}
		}
		if l1.Next != nil {
			c1++
			temps = l1
			l1 = l1.Next
			v1 = l1.Val
		}
		if l2.Next != nil {
			c2++
			temps = l2
			l2 = l2.Next
			v2 = l2.Val
		}
		tempL.Val = v3
		temps.Val = 0
		temps.Next = nil
		tempL.Next = temps
		tempL = tempL.Next
	}
	if v := v2 + v1 + temp; v != 0 {
		tempL.Val = v
		if v > 9 {
			v -= 10
			tempL.Val = v
			l1.Val = 1
			l1.Next = nil
			tempL.Next = l1
		}
	}
	return res
}

//8 ms	4.2 MB 优化内存
func addTwoNumbers3(l1 *ListNode, l2 *ListNode) *ListNode {
	var res = new(ListNode)
	tempL := res
	var temps *ListNode
	v1, v2 := l1.Val, l2.Val
	var c1, c2 int
	var temp, v3 int
	for l1.Next != nil || l2.Next != nil {
		v3 = v1 + v2 + temp
		v1 = 0
		v2 = 0
		temp = 0
		if v3 > 9 {
			v3 -= 10
			temp = 1
		}
		if l1.Next != nil {
			c1++
			temps = l1
			l1 = l1.Next
			v1 = l1.Val
		}
		if l2.Next != nil {
			c2++
			temps = l2
			l2 = l2.Next
			v2 = l2.Val
		}
		tempL.Val = v3
		temps.Val = 0
		temps.Next = nil
		tempL.Next = temps
		tempL = tempL.Next
	}
	if v := v2 + v1 + temp; v != 0 {
		tempL.Val = v
		if v > 9 {
			v -= 10
			tempL.Val = v
			l1.Val = 1
			l1.Next = nil
			tempL.Next = l1
		}
	}
	return res
}

//4 ms	4.4 MB 优化循环
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	var res = new(ListNode)
	tempL := res
	v1, v2 := l1.Val, l2.Val
	var temp int
	for l1.Next != nil || l2.Next != nil {
		tempL.Val = v1 + v2 + temp
		v1 = 0
		v2 = 0
		temp = 0
		if tempL.Val > 9 {
			tempL.Val -= 10
			temp = 1
		}

		if temp == 0 {
			if l1.Next == nil {
				tempL.Next = l2.Next
				return res
			}
			if l2.Next == nil {
				tempL.Next = l1.Next
				return res
			}
		}

		if l1.Next != nil {
			l1 = l1.Next
			v1 = l1.Val
		}
		if l2.Next != nil {
			l2 = l2.Next
			v2 = l2.Val
		}
		tempL.Next = &ListNode{}
		tempL = tempL.Next
	}

	if v := v2 + v1 + temp; v != 0 {
		tempL.Val = v
		if tempL.Val > 9 {
			tempL.Val -= 10
			tempL.Next = &ListNode{1, nil}
		}
	}
	return res
}

//12 ms	4.4 MB  普通版
func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	var res = new(ListNode)
	tempL := res
	v1, v2 := l1.Val, l2.Val
	var temp int
	if l1.Next == nil && l2.Next == nil {
		tempL.Val = v1 + v2
		if tempL.Val > 9 {
			tempL.Val -= 10
			tempL.Next = &ListNode{1, nil}
		}
		return res
	}
	for l1.Next != nil || l2.Next != nil {
		tempL.Val = v1 + v2 + temp
		v1 = 0
		v2 = 0
		temp = 0
		if tempL.Val > 9 {
			tempL.Val -= 10
			temp = 1
		}
		tempL.Next = &ListNode{}
		tempL = tempL.Next
		if l1.Next != nil {
			l1 = l1.Next
			v1 = l1.Val
		}
		if l2.Next != nil {
			l2 = l2.Next
			v2 = l2.Val
		}
	}
	if v := v2 + v1 + temp; v != 0 {
		tempL.Val = v
		if tempL.Val > 9 {
			tempL.Val -= 10
			tempL.Next = &ListNode{1, nil}
		}
	}
	return res
}

func main() {
	var res, l1, l2 *ListNode
	//1,8
	l1 = &ListNode{1, &ListNode{8, nil}}
	//0
	l2 = &ListNode{0, nil}
	res = addTwoNumbers(l1, l2)
	println(18)
	for res.Next != nil {
		print(res.Val)
		res = res.Next
	}
	println(res.Val)

	//1
	l1 = &ListNode{1, nil}
	//9,9
	l2 = &ListNode{9, &ListNode{9, nil}}
	res = addTwoNumbers(l1, l2)
	println(001)
	for res.Next != nil {
		print(res.Val)
		res = res.Next
	}
	println(res.Val)

	//9,9,9,9,9,9,9
	l1 = &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}}
	//9,9,9,9
	l2 = &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}
	res = addTwoNumbers(l1, l2)
	println(89990001)
	for res.Next != nil {
		print(res.Val)
		res = res.Next
	}
	println(res.Val)
	//[2,4,3]
	l1 = &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	//[5,6,4]
	l2 = &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	res = addTwoNumbers(l1, l2)
	println(708)
	for res.Next != nil {
		print(res.Val)
		res = res.Next
	}
	println(res.Val)
	//[0]
	l1 = &ListNode{0, nil}
	//[7,3]
	l2 = &ListNode{7, &ListNode{3, nil}}
	res = addTwoNumbers(l1, l2)
	println(73)
	for res.Next != nil {
		print(res.Val)
		res = res.Next
	}
	println(res.Val)

	//[1,8]
	l1 = &ListNode{1, &ListNode{8, nil}}
	//[0]
	l2 = &ListNode{0, nil}
	res = addTwoNumbers(l1, l2)
	println(18)
	for res.Next != nil {
		print(res.Val)
		res = res.Next
	}
	println(res.Val)

	//[0]
	l1 = &ListNode{0, nil}
	//[0]
	l2 = &ListNode{0, nil}
	res = addTwoNumbers(l1, l2)
	println(0)
	for res.Next != nil {
		print(res.Val)
		res = res.Next
	}
	println(res.Val)

	//[991]
	l1 = &ListNode{9, &ListNode{9, &ListNode{1, nil}}}
	//[1]
	l2 = &ListNode{1, nil}
	res = addTwoNumbers(l1, l2)
	println(002)
	for res.Next != nil {
		print(res.Val)
		res = res.Next
	}
	println(res.Val)
}
