package main

import "fmt"

func Print1ToMaxOfNDigits(n int) {
	if n <= 0 {
		return
	}
	number := make([]int, n)
	for i := 0; i < 10; i++ {
		number[0] = i
		Print1ToMaxOfNDigitsRecursively(number, n, 0)
	}
}

func Print1ToMaxOfNDigitsRecursively(num []int, length, index int) {
	if index == length-1 {
		PrintNumber(num)
		return
	}
	for i := 0; i < 10; i++ {
		num[index+1] = i
		Print1ToMaxOfNDigitsRecursively(num, length, index+1)
	}
}

func PrintNumber(num []int) {
	isBegin := true
	for i := 0; i < len(num); i++ {
		if isBegin && num[i] != 0 {
			isBegin = false
		}
		if !isBegin {
			fmt.Printf("%d", num[i])
		}
	}
	fmt.Println()
}

func main() {
	Print1ToMaxOfNDigits(3)
}
