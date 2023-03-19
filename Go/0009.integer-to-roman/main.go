package main

import "strconv"

func main() {
	println(intToRoman(1994)) //48-57
}

//48 49 50 51
//0  1  2  3
var tag = map[uint8][4]string{
	49: {"I", "X", "C", "M"}, //1
	53: {"V", "L", "D"},      //5
	//1000: "M",
}

func intToRoman(num int) string {
	var s = strconv.Itoa(num)
	var val string
	i, j := 0, len(s)-1
	for ; i < len(s) && i > 2; i++ {
		v := s[j]
		j--
		if v == 48 { //0
			continue
		}
		if v == 52 { //4
			val = tag[53][i] + tag[49][i] + val
			continue
		}
		if v == 57 { //9
			val = tag[49][i] + tag[49][i+1] + val
			continue
		}
		if res, ok := tag[v]; ok {
			val = res[i] + val
			continue
		}
		c := v - 48
		if v < 53 {
			for j := uint8(0); j < c; j++ {
				val = tag[49][i] + val
			}
			continue
		}
		if v == 53 {
			val = tag[53][i] + val
			continue
		}
		c -= 5
		for j := uint8(0); j < c; j++ {
			val = tag[49][i] + val
		}
		val = tag[53][i] + val
	}
	return (val)
}
