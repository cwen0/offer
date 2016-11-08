package main

import (
	"fmt"
	"sort"
	"strconv"
)

type Data []string

func (d Data) Len() int {
	return len(d)
}

func (d Data) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (d Data) Less(i, j int) bool {
	t1 := d[i] + d[j]
	t2 := d[j] + d[i]
	return t1 < t2
}

func PrintMinNumber(data []int) {
	if len(data) == 0 {
		return
	}
	var byteData []string
	for i := 0; i < len(data); i++ {
		byteData = append(byteData, strconv.Itoa(data[i]))
	}
	sort.Sort(Data(byteData))
	for i := 0; i < len(byteData); i++ {
		fmt.Printf("%v", byteData[i])
	}
}

func main() {
	data := []int{3, 32, 321}
	PrintMinNumber(data)
}
