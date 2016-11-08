package main

import (
	"errors"
	"fmt"
	"log"
)

func Duplicate(data []int) (int, bool, error) {
	if len(data) == 0 {
		return 0, false, errors.New("data is empty")
	}
	for i := 0; i < len(data); i++ {
		if data[i] < 0 || data[i] > len(data)-1 {
			return 0, false, errors.New("input is wrong")
		}
	}
	for i := 0; i < len(data); i++ {
		for data[i] != i {
			if data[i] == data[data[i]] {
				return data[i], true, nil
			}
			data[i], data[data[i]] = data[data[i]], data[i]
		}
	}
	return 0, false, nil
}

func main() {
	data := []int{2, 3, 1, 0, 2, 5, 3}
	re, isFound, err := Duplicate(data)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	if isFound {
		fmt.Println(re)
	} else {
		fmt.Println("no found")
	}
}
