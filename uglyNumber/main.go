// 方法一:逐个判断每个整数是不是丑数的解法
// 方法二:空间换时间，保存已有的丑数
package main

import "fmt"

// 方法一

func IsUgly(num int) bool {
	for num%2 == 0 {
		num = num / 2
	}
	for num%3 == 0 {
		num = num / 3
	}
	for num%5 == 0 {
		num = num / 5
	}
	if num == 1 {
		return true
	}
	return false
}

func GetUglyNumber(index int) int {
	if index <= 0 {
		return 0
	}

	num := 0
	uglyFound := 0
	for uglyFound < index {
		num++
		if IsUgly(num) {
			uglyFound++
		}
	}
	return num
}

// 方法二

func GetUglyNumber2(index int) int {
	if index <= 0 {
		return 0
	}
	uglyArry := make([]int, index)
	uglyArry[0] = 1
	nextUglyIndex := 1
	index2 := 0
	index3 := 0
	index5 := 0
	for nextUglyIndex < index {
		min := Min(uglyArry[index2]*2, uglyArry[index3]*3, uglyArry[index5]*5)
		uglyArry[nextUglyIndex] = min
		for uglyArry[index2]*2 <= uglyArry[nextUglyIndex] {
			index2++
		}
		for uglyArry[index3]*3 <= uglyArry[nextUglyIndex] {
			index3++
		}
		for uglyArry[index5]*5 <= uglyArry[nextUglyIndex] {
			index5++
		}
		nextUglyIndex++
	}

	return uglyArry[nextUglyIndex-1]
}
func Min(num1, num2, num3 int) int {
	min := num1
	if num2 < min {
		min = num2
	}
	if num3 < min {
		min = num3
	}
	return min
}
func main() {
	fmt.Println(GetUglyNumber(1500))
	fmt.Println(GetUglyNumber2(1500))
}
