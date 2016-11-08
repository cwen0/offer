package main

import (
	"errors"
	"fmt"
	"log"
)

func FirstNotRepeatingChar(data string) (rune, error) {
	if len(data) == 0 {
		return ' ', errors.New("data is empry")
	}
	var hashTable [256]int
	for i := 0; i < len(hashTable); i++ {
		hashTable[i] = 0
	}

	for i := 0; i < len(data); i++ {
		hashTable[rune(data[i])]++
	}

	for i := 0; i < len(data); i++ {
		if hashTable[rune(data[i])] == 1 {
			return rune(data[i]), nil
		}
	}
	return ' ', nil
}

func main() {
	data := "ababccdeff"
	s, err := FirstNotRepeatingChar(data)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Printf("%c\n", s)
}
