package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntHeap) Top() int {
	temp := *h
	return temp[len(temp)-1]
}

func GetLeastNumbers(data []int, h *IntHeap, k int) {
	heap.Init(h)
	if k < 1 || len(data) < k {
		return
	}

	for i := 0; i < len(data); i++ {
		if h.Len() < k {
			h.Push(data[i])
		} else {
			if data[i] < h.Top() {
				h.Pop()
				h.Push(data[i])
			}
		}
	}
}

func main() {
	data := []int{2, 4, 6, 7, 4, 7, 8, 3, 222}
	h := &IntHeap{}
	GetLeastNumbers(data, h, 4)
	for i := 0; i < h.Len(); i++ {
		fmt.Println((*h)[i])
	}
}
