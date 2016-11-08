package main

import (
	"fmt"
	"math"
)

const FACE_NUM = 6

func PrintProbability(number int) {
	if number <= 0 {
		return
	}

	sum := make([]int, number*FACE_NUM+1)
	total := math.Pow(6, float64(number))
	size := number * FACE_NUM
	sum[0] = 0
	for i := 1; i <= FACE_NUM; i++ {
		sum[i] = 1
	}

	for i := FACE_NUM + 1; i <= size; i++ {
		sum[i] = 0
	}

	for i := 2; i <= number; i++ {
		j := 0
		for j := i * FACE_NUM; j >= i; j-- {
			sum[j] = 0
			for k := 1; k <= 6 && j >= k; k++ {
				sum[j] += sum[j-k]
			}
		}
		for j = i - 1; j >= 0; j-- {
			sum[j] = 0
		}
	}
	for i := number; i <= size; i++ {
		fmt.Printf("sum= %v, p= %v\n", i, float64(sum[i])/total)
	}
}

func main() {
	PrintProbability(4)
}
