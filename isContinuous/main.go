package main

import (
	"errors"
	"fmt"
	"log"
	"sort"
)

func IsContinuous(data []int) (bool, error) {
	if len(data) <= 0 {
		return false, errors.New("data is empty")
	}

	sort.Ints(data)
	numZero := 0
	numGap := 0
	for i := 0; i < len(data); i++ {
		if data[i] == 0 {
			numZero++
		}
	}
	small := numZero
	big := small + 1
	for big < len(data) {
		if data[small] == data[big] {
			return false, nil
		}
		numGap += data[big] - data[small] - 1
		small = big
		big++
	}

	if numGap > numZero {
		return false, nil
	}
	return true, nil
}

func main() {
	data := []int{6, 3, 1, 4, 5}
	re, err := IsContinuous(data)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Println(re)
}
