package main

import (
	"errors"
	"fmt"
	"log"
)

func Find(data []int, rows, columns, number int) (bool, error) {
	found := false
	if len(data) == 0 || rows <= 0 || columns <= 0 {
		return found, errors.New("input is wrong")
	}

	row := 0
	column := columns - 1
	for row < rows && column >= 0 {
		if data[row*columns+column] == number {
			found = true
			return found, nil
		} else if data[row*columns+column] > number {
			column--
		} else {
			row++
		}
	}
	return found, nil
}

func main() {
	data := []int{1, 2, 8, 9, 2, 4, 9, 12, 4, 7, 10, 13, 6, 8, 11, 15}
	re, err := Find(data, 4, 4, 5)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Println(re)
}
