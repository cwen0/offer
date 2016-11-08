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

type StackWithMin struct {
	Data Stack
	min  Stack
}

func (s *StackWithMin) Push(num int) {
	s.Data.Push(num)
	if s.min.Size() == 0 || num < s.min.Top() {
		s.min.Push(num)
	} else {
		s.min.Push(s.min.Top())
	}
}
func (s *StackWithMin) Pop() error {
	if s.Data.Size() > 0 && s.min.Size() > 0 {
		s.Data.Top()
		s.min.Pop()
		return nil
	}
	return errors.New("wrong")
}

func (s *StackWithMin) Min() (int, error) {
	if s.Data.Size() > 0 && s.min.Size() > 0 {
		return s.min.Top(), nil
	}
	return 0, errors.New("wrong")

}
func main() {
	var t StackWithMin
	t.Push(3)
	m, _ := t.Min()
	fmt.Println(m)
	t.Push(2)
	m, _ = t.Min()
	fmt.Println(m)
	t.Push(6)
	m, _ = t.Min()
	fmt.Println(m)
	t.Push(1)
	m, _ = t.Min()
	fmt.Println(m)
	t.Pop()
	m, _ = t.Min()
	fmt.Println(m)
}
