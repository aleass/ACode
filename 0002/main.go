package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var res, l1, l2 *ListNode
	l1 = &ListNode{5, nil}
	l2 = &ListNode{5, nil}
	res = addTwoNumbers2(l1, l2)
	println("1")
	for res.Next != nil {
		print(res.Val)
		res = res.Next
	}
	println(res.Val)
	l1 = &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 = &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	res = addTwoNumbers2(l1, l2)
	println("708")
	for res.Next != nil {
		print(res.Val)
		res = res.Next
	}
	println(res.Val)
	return

	//1,8
	l1 = &ListNode{1, &ListNode{8, nil}}
	//0
	l2 = &ListNode{0, nil}
	res = addTwoNumbers(l1, l2)
	println("18")
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
	println("001")
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
	println("89990001")
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
	println("708")
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
	println("73")
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
	println("18")
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
	println("0")
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
	println("002")
	for res.Next != nil {
		print(res.Val)
		res = res.Next
	}
	println(res.Val)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//创建头
	var res = new(ListNode)
	//复制信息
	temp := res
	var reTemp *ListNode //复用
	v1, v2 := l1.Val, l2.Val
	//
	var power, v3 int
	for l1.Next != nil || l2.Next != nil {
		v3 = v1 + v2 + power
		v1, v2, power = 0, 0, 0
		if v3 > 9 {
			v3 -= 10
			power = 1
		}
		if power == 0 {
			if l1.Next == nil {
				temp.Val = v3
				temp.Next = l2.Next
				return res
			}
			if l2.Next == nil {
				temp.Val = v3
				temp.Next = l1.Next
				return res
			}
		}
		if l1.Next != nil {
			reTemp = l1
			l1 = l1.Next
			v1 = l1.Val
		}
		if l2.Next != nil {
			reTemp = l2
			l2 = l2.Next
			v2 = l2.Val
		}
		temp.Val = v3
		reTemp.Val = 0    //清空
		reTemp.Next = nil //清空
		temp.Next = reTemp
		temp = temp.Next
	}
	if v := v2 + v1 + power; v != 0 {
		temp.Val = v
		if v > 9 {
			v -= 10
			temp.Val = v
			l1.Val = 1
			l1.Next = nil
			temp.Next = l1
		}
	}
	return res
}

func addTwoNumbers3(l1, l2 *ListNode) *ListNode {
	var _res = new(ListNode)
	var res = _res
	var power int
	val := l1.Val + l2.Val + power
	for l1.Next != nil && l2.Next != nil {
		power = 0
		for val > 9 {
			val -= 10
			power++
		}
		_res.Val = val
		_res.Next = new(ListNode)
		_res = _res.Next
		l1 = l1.Next
		l2 = l2.Next
		val = l1.Val + l2.Val + power
	}

	if l1.Next == nil {
		for l2.Next != nil {
			power = 0
			if val > 9 {
				for val > 9 {
					val -= 10
					power++
				}
			} else {
				_res.Val = val
				_res.Next = l2.Next
				return res
			}
			_res.Val = val
			_res.Next = new(ListNode)
			_res = _res.Next
			l2 = l2.Next
			val = l2.Val + power
		}
		power = 0
		if val > 9 {
			for val > 9 {
				val -= 10
				power++
			}
		}
		_res.Val = val
		if power > 0 {
			_res.Next = new(ListNode)
			_res = _res.Next
			_res.Val = power
		}
		return res
	}

	if l2.Next == nil {
		for l1.Next != nil {
			power = 0
			if val > 9 {
				for val > 9 {
					val -= 10
					power++
				}
			} else {
				_res.Val = val
				_res.Next = l1.Next
				return res
			}
			_res.Val = val
			_res.Next = new(ListNode)
			_res = _res.Next
			l1 = l1.Next
			val = l1.Val + power
		}
		power = 0
		if val > 9 {
			for val > 9 {
				val -= 10
				power++
			}
		}
		_res.Val = val
		if power > 0 {
			_res.Next = new(ListNode)
			_res = _res.Next
			_res.Val = power
		}
	}

	return res
}

func addTwoNumbers2(l1, l2 *ListNode) *ListNode {
	var _res = new(ListNode)
	var res = _res
	var power int
	val := l1.Val + l2.Val + power
	for l1.Next != nil && l2.Next != nil {
		power = 0
		for val > 9 {
			val -= 10
			power++
		}
		_res.Val = val
		_res.Next = new(ListNode)
		_res = _res.Next
		l1 = l1.Next
		l2 = l2.Next
		val = l1.Val + l2.Val + power
	}

	if l2.Next == nil && l1.Next == nil {
		power = 0
		if val > 9 {
			for val > 9 {
				val -= 10
				power++
			}
			_res.Val = val
			_res.Next = &ListNode{}
		}
		if power > 0 {
			_res.Val = power
			_res.Next = new(ListNode)
		}
		return res
	}

	var notEmpty = l2
	if l2.Next == nil {
		notEmpty = l1
	}

	for notEmpty.Next != nil {
		power = 0
		if val > 9 {
			for val > 9 {
				val -= 10
				power++
			}
		} else {
			_res.Val = val
			_res.Next = notEmpty.Next
			return res
		}
		_res.Val = val
		_res.Next = new(ListNode)
		_res = _res.Next
		notEmpty = notEmpty.Next
		val = notEmpty.Val + power
	}
	power = 0
	if val > 9 {
		for val > 9 {
			val -= 10
			power++
		}
	}
	_res.Val = val
	if power > 0 {
		_res.Next = new(ListNode)
		_res = _res.Next
		_res.Val = power
	}
	return res
}
