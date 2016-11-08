package main

import "fmt"

type Judge func(n int) bool

func isEvent(n int) bool {
	if n&1 == 0 {
		return true
	}
	return false
}

func Reorder(data []int, f Judge) {
	if len(data) == 0 {
		return
	}
	indexBegin := 0
	indexEnd := len(data) - 1
	for indexBegin < indexEnd {
		for indexBegin < indexEnd && !f(data[indexBegin]) {
			indexBegin++
		}
		for indexBegin < indexEnd && f(data[indexEnd]) {
			indexEnd--
		}
		if indexBegin < indexEnd {
			data[indexBegin], data[indexEnd] = data[indexEnd], data[indexBegin]
		}
	}
}

func main() {
	data := []int{1, 2, 3, 4, 5}
	Reorder(data, isEvent)
	fmt.Println(data)
}
