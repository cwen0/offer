package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Qsort(data []int, start, end int) {
	if start == end {
		return
	}
	index := Partition(data, start, end)
	if index > start {
		Qsort(data, start, index-1)
	}
	if index < end {
		Qsort(data, index+1, end)
	}
}

func Partition(data []int, start, end int) int {
	index := RandIndex(start, end)
	Swap(data, index, end)
	small := start - 1
	for index = start; index < end; index++ {
		if data[index] < data[end] {
			small++
			if small != index {
				Swap(data, index, small)
			}
		}
	}
	small++
	Swap(data, small, end)
	return small
}

func Swap(data []int, i, j int) {
	data[j], data[i] = data[i], data[j]
}

func RandIndex(start, end int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := r.Intn((end - start)) + start
	return num
}

func main() {
	data := []int{6, 3, 6, 3, 2, 77, 4}
	Qsort(data, 0, len(data)-1)
	for i := 0; i < len(data); i++ {
		fmt.Printf("%d ", data[i])
	}
}
