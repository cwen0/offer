package main

import (
	"errors"
	"fmt"
)

type Stack []int

func (s *Stack) Push(num int) {
	*s = append(*s, num)
}

func (s *Stack) Pop() error {
	if s.Size() == 0 {
		return errors.New("stack is empty")
	}
	*s = (*s)[:len(*s)-1]
	return nil
}

func (s *Stack) Top() int {
	return (*s)[len(*s)-1]
}

func (s *Stack) Size() int {
	return len(*s)
}

func IsPopOrder(push, pop []int) bool {
	if len(push) != len(pop) {
		return false
	}
	nextPush := 0
	nextPop := 0
	var pushStack Stack
	for nextPop < len(pop) {
		for pushStack.Size() == 0 || pushStack.Top() != pop[nextPop] {
			if nextPush == len(push) {
				break
			}
			pushStack.Push(push[nextPush])
			nextPush++
		}
		if pushStack.Top() != pop[nextPop] {
			break
		}
		pushStack.Pop()
		nextPop++
		if nextPop == len(pop) && nextPush == len(push) {
			return true
		}
	}
	return false
}

func main() {
	push := []int{1, 4, 6, 3}
	pop := []int{6, 3, 4, 1}
	re := IsPopOrder(push, pop)
	fmt.Println(re)
}
