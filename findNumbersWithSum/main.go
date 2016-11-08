package main

import "fmt"

func FindNumbersWithSum(data []int, sum int64) int {
	if len(data) < 1 {
		return 0
	}
	ahead := len(data) - 1
	behind := 0
	count := 0
	for behind < ahead {
		var curSum int64
		curSum = int64(data[ahead]) + int64(data[behind])
		if curSum == sum {
			count++
			behind++
		} else if curSum > sum {
			ahead--
		} else {
			behind++
		}
	}
	return count
}

func main() {
	data := []int{1, 2, 3, 4, 7, 11, 12, 15}
	var sum int64 = 15
	count := FindNumbersWithSum(data, sum)
	fmt.Println(count)
}
