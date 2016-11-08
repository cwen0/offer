package main

import "fmt"

func FindContinousSequence(sum int) {
	if sum < 3 {
		return
	}
	small := 1
	big := 2
	middle := (1 + sum) / 2
	curSum := small + big
	for small < middle {
		if curSum == sum {
			PrintContinousSequence(small, big)
		}

		for curSum > sum && small < middle {
			curSum -= small
			small++
			if curSum == sum {
				PrintContinousSequence(small, big)
			}
		}
		big++
		curSum += big
	}
}

func PrintContinousSequence(start, end int) {
	for i := start; i <= end; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

func main() {
	FindContinousSequence(15)
}
