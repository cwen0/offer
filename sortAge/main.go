package main

import (
	"errors"
	"fmt"
)

func SortAge(age []int) error {
	if len(age) == 0 {
		return errors.New("data is empty")
	}
	const oldestAge = 99
	var timeOfAge [oldestAge + 1]int
	for i := 0; i < len(age); i++ {
		if age[i] < 0 || age[i] > oldestAge {
			return errors.New("age is wrong")
		}
		timeOfAge[age[i]]++
	}
	index := 0
	for i := 0; i <= oldestAge; i++ {
		for j := 0; j < timeOfAge[i]; j++ {
			age[index] = i
			index++
		}
	}
	return nil
}

func main() {
	age := []int{34, 5, 13, 23, 5}
	SortAge(age)
	for i := 0; i < len(age); i++ {
		fmt.Printf("%d ", age[i])
	}
}
