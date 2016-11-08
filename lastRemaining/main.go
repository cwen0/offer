package main

import "fmt"

const MAXN = 10000

func LastRemainingUseArr(n, m int) int {
	if n < 1 || m < 1 {
		return -1
	}
	var Jose [MAXN]bool
	curCount := 0
	outCount := 0
	index := 0
	for outCount < n-1 {
		index = index % n
		if !Jose[index] {
			curCount++
		}
		if curCount == m {
			curCount = 0
			Jose[index] = true
			outCount++
		}
		index++
	}
	for i := 0; i < n; i++ {
		if !Jose[i] {
			return i
		}
	}
	return -1
}

func LastRemaining(n, m int) int {
	if n < 1 || m < 1 {
		return -1
	}

	last := 0
	for i := 2; i <= n; i++ {
		last = (last + m) % i
	}
	return last
}

func main() {
	fmt.Println(LastRemaining(10, 4))
	fmt.Println(LastRemainingUseArr(10, 4))
}
