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

func ReverseSequence(str *string) {
	if str == nil || len(*str) == 0 {
		return
	}

	Reverse(str, 0, len(*str)-1)
	begin := 0
	end := 0
	for begin != len(*str)-1 {
		if (*str)[begin] == ' ' {
			begin++
			end++
		} else if (*str)[end] == ' ' || end == len(*str)-1 {
			end--
			Reverse(str, begin, end)
			end++
			begin = end
		} else {
			end++
		}
	}
}

func main() {
	str := "I am a student."
	ReverseSequence(&str)
	fmt.Println(str)
}
