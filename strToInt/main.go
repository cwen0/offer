package main

import (
	"errors"
	"fmt"
	"log"
)

func StrToInt(str string) (int64, error) {
	if len(str) <= 0 {
		return 0, errors.New("str is empty")
	}

	arr := []rune(str)
	indexArr := 0
	minus := false
	if arr[0] == '-' {
		minus = true
		indexArr++
	} else if arr[0] == '+' {
		indexArr++
	}
	num, err := StrToIntCore(arr, indexArr, minus)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func StrToIntCore(arr []rune, index int, minus bool) (int64, error) {
	var num int64
	flag := 1
	if minus {
		flag = -1
	}
	for i := index; i < len(arr); i++ {
		if arr[i] >= '0' && arr[i] <= '9' {
			num = num*int64(10) + int64(flag)*int64((arr[i]-'0'))
		} else {
			return 0, errors.New("input is wrong")
		}
	}
	return num, nil
}

func main() {
	str := "123"
	vlu, err := StrToInt(str)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Println(vlu)
}
