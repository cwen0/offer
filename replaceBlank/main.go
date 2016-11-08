package main

import (
	"fmt"
	"strings"
)

func ReplaceBlank(str string) string {
	str = strings.Replace(str, " ", "%20", -1)
	return str
}

func main() {
	str := "we are strudent"
	re := ReplaceBlank(str)
	fmt.Println(re)
}
