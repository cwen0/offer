package main

import "fmt"

func Reverse(str *string, begin, end int) {
	if str == nil || begin >= end {
		return
	}
	tmp := []rune(*str)
	for begin < end {
		tmp[begin], tmp[end] = tmp[end], tmp[begin]
		begin++
		end--
	}
	*str = string(tmp)
}

func LeftRotateString(str *string, n int) {
	if str == nil || len(*str) <= 1 || n > len(*str)-1 {
		return
	}
	Reverse(str, 0, n-1)
	Reverse(str, n, len(*str)-1)
	Reverse(str, 0, len(*str)-1)
}

func main() {
	str := "abcdefg"
	LeftRotateString(&str, 2)
	fmt.Println(str)
}
