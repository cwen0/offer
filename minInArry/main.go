package main

import (
	"errors"
	"fmt"
	"log"
)

func Min(data []int) (int, error) {
	if len(data) == 0 {
		return 0, errors.New("data is empty")
	}
	index1 := 0
	index2 := len(data) - 1
	indexMid := index1
	for data[index1] >= data[index2] {
		if index2-index1 == 1 {
			indexMid = index2
			break
		}

		indexMid = (index1 + index2) / 2
		if data[index1] == data[indexMid] && data[indexMid] == data[index2] {
			return MinInOrder(data, index1, index2), nil
		}

		if data[index1] <= data[indexMid] {
			index1 = indexMid
		} else if data[index2] >= data[indexMid] {
			index2 = indexMid
		}
	}
	return data[indexMid], nil
}

func MinInOrder(data []int, index1, index2 int) int {
	min := data[index1]
	for i := index1; i <= index2; i++ {
		if data[i] < min {
			min = data[i]
		}
	}
	return min
}

func main() {
	data := []int{3, 4, 5, 1, 2}
	min, err := Min(data)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Println(min)
}
