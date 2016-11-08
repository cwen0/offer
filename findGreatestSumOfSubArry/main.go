package main

import (
	"errors"
	"fmt"
	"log"
)

func FindGreatestSumOfSubArry(data []int) (int, error) {
	if len(data) == 0 {
		return 0, errors.New("array is empty")
	}

	greatestSum := data[0]
	curSum := 0
	for i := 0; i < len(data); i++ {
		if curSum <= 0 {
			curSum = data[i]
		} else {
			curSum += data[i]
		}

		if curSum > greatestSum {
			greatestSum = curSum
		}
	}
	return greatestSum, nil
}

func main() {

	data := []int{1, -2, 3, 10, -4, 7, 2, -5}
	sum, err := FindGreatestSumOfSubArry(data)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Println(sum)
}
