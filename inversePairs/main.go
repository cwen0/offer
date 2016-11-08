package main

import "fmt"

func InverseParse(data []int) int {
	if len(data) == 0 {
		return 0
	}
	copy := make([]int, len(data))
	count := InverseParseCore(data, copy, 0, len(data)-1)
	return count
}

func InverseParseCore(data, copy []int, start, end int) int {
	if start == end {
		copy[start] = data[start]
		return 0
	}

	length := (end - start) / 2
	left := InverseParseCore(copy, data, start, start+length)
	right := InverseParseCore(copy, data, start+length+1, end)
	i := start + length
	j := end
	count := 0
	copyIndex := end
	for i >= start && j >= start+length+1 {
		if data[i] > data[j] {
			copy[copyIndex] = data[i]
			copyIndex--
			i--
			count += j - start - length
		} else {
			copy[copyIndex] = data[j]
			copyIndex--
			j--
		}
	}

	for i >= start {
		copy[copyIndex] = data[i]
		copyIndex--
		i--
	}
	for j >= start+length+1 {
		copy[copyIndex] = data[j]
		copyIndex--
		j--
	}
	return left + right + count
}

func main() {
	data := []int{7, 5, 6, 4}
	sum := InverseParse(data)
	fmt.Println(sum)
}
