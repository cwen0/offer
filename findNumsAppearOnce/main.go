package main

import "fmt"

func FindNumsAppearOnce(data []int, num1, num2 *int) {
	if data == nil || len(data) < 2 {
		return
	}
	resultExclusiveOR := 0
	for i := 0; i < len(data); i++ {
		resultExclusiveOR ^= data[i]
	}

	indexOf1 := FindFirstBitIs1(resultExclusiveOR)
	*num1, *num2 = 0, 0
	for i := 0; i < len(data); i++ {
		if IsBit1(data[i], uint(indexOf1)) {
			*num1 ^= data[i]
		} else {
			*num2 ^= data[i]
		}
	}
}

func FindFirstBitIs1(num int) int {
	indexBit := 0
	for num&1 == 0 {
		num = num >> 1
		indexBit++
	}
	return indexBit
}

func IsBit1(num int, index uint) bool {
	num = num >> index
	if num&1 == 1 {
		return true
	}
	return false
}

func main() {
	data := []int{2, 4, 3, 6, 3, 2, 5, 5}
	num1, num2 := 0, 0
	FindNumsAppearOnce(data, &num1, &num2)
	fmt.Println(num1)
	fmt.Println(num2)
}
