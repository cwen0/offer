package main

import "fmt"

func NumberOf1Between1AndN(n int) int {
	base := 1
	passed := 0
	sum := 0
	for n != 0 {
		current := n % 10
		n = n / 10
		if current > 1 {
			sum += (n + 1) * base
		} else if current == 1 {
			sum += n*base + passed + 1
		} else {
			sum += n * base
		}

		passed = current * base
		base *= 10
	}
	return sum
}

func main() {
	sum := NumberOf1Between1AndN(12)
	fmt.Println(sum)
}
