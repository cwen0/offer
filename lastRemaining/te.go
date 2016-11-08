package main

import (
	"errors"
	"fmt"
)

func reverse(str *string, begin, end int) {
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

func shift(str string, n int) (string, error) {
	if len(str) <= 1 || n > len(str)-1 {
		return str, errors.New("err")
	}
	if n == 0 {
		return str, nil
	}
	reverse(&str, 0, n-1)
	reverse(&str, n, len(str)-1)
	reverse(&str, 0, len(str)-1)
	return str, nil
}

func main() {

	str := ""
	fmt.Scanln(&str)
	count := 0
	for i := 0; i < len(str); i++ {
		tmp, _ := shift(str, i)
		if tmp == str {
			count++
		}
	}
	fmt.Println(count)

}
