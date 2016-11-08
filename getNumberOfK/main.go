package main

import "fmt"

func GetNumberOfK(data []int, k int) int {
	if len(data) == 0 {
		return 0
	}
	first := GetFirstK(data, k, 0, len(data)-1)
	last := GetLastK(data, k, 0, len(data)-1)
	return last - first + 1
}

func GetFirstK(data []int, k, start, end int) int {
	if start > end {
		return -1
	}

	indexMiddle := (start + end) / 2
	if data[indexMiddle] == k {
		if (indexMiddle > 0 && data[indexMiddle-1] != k) || indexMiddle == 0 {
			return indexMiddle
		} else {
			end = end - 1
		}
	} else if data[indexMiddle] > k {
		end = indexMiddle - 1
	} else {
		start = indexMiddle + 1
	}
	return GetFirstK(data, k, start, end)
}

func GetLastK(data []int, k, start, end int) int {
	if start > end {
		return -1
	}

	indexMiddle := (start + end) / 2
	if data[indexMiddle] == k {
		if (indexMiddle < len(data)-1 && data[indexMiddle+1] != k) || indexMiddle == len(data)-1 {
			return indexMiddle
		} else {
			start = indexMiddle + 1
		}
	} else if data[indexMiddle] < k {
		start = indexMiddle + 1
	} else {
		end = indexMiddle - 1
	}
	return GetLastK(data, k, start, end)
}

func main() {
	data := []int{1, 2, 3, 3, 3, 3, 4, 5}
	count := GetNumberOfK(data, 3)
	fmt.Println(count)
}
