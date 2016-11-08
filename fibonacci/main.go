package main

import "fmt"

func Fibonacci1(n int) int64 {
	if n <= 0 {
		return int64(0)
	}
	if n == 1 {
		return int64(1)
	}
	return Fibonacci1(n-1) + Fibonacci1(n-2)
}

func Fibonacci2(n int) int64 {
	result := [2]int64{0, 1}
	if n < 2 {
		return result[n]
	}
	var fibOne int64 = 1
	var fibTwo int64 = 0
	var fibN int64
	for i := 2; i <= n; i++ {
		fibN = fibOne + fibTwo
		fibTwo = fibOne
		fibOne = fibN
	}
	return fibN
}

func main() {
	fmt.Println(Fibonacci1(4))
	fmt.Println(Fibonacci2(4))
}
