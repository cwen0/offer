package main

import "fmt"

func PrintMatrix(data [][]int, colums, rows int) error {
	if len(data) == 0 {
		return nil
	}
	start := 0
	for colums > start*2 && rows > start*2 {
		PrintMatrixInCircle(data, colums, rows, start)
		start++
	}
	return nil
}

func PrintMatrixInCircle(data [][]int, colums, rows, start int) {
	endX := colums - 1 - start
	endY := rows - 1 - start
	for i := start; i <= endX; i++ {
		fmt.Printf("%d ", data[start][i])
	}

	if start < endY {
		for i := start + 1; i <= endY; i++ {
			fmt.Printf("%d ", data[i][endX])
		}
	}

	if start < endX && start < endY {
		for i := endY - 1; i >= start; i-- {
			fmt.Printf("%d ", data[endY][i])
		}
	}

	if start < endX && start < endY-1 {
		for i := endY - 1; i > start; i-- {
			fmt.Printf("%d ", data[i][start])
		}
	}
}

func main() {
	data := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	PrintMatrix(data, 4, 4)
}
