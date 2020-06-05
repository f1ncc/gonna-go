package main

import (
	"fmt"
	"strings"
	"strconv"
)

func truncate(f float64) int {
	s := strconv.FormatFloat(f, 'f', -1, 64)
	trunc := strings.Split(s,".")[0]
	i, err := strconv.Atoi(trunc)
	if(err != nil) {
		return -1
	}
	return i
}

func main() {
	var ogfloat64 float64
	fmt.Printf("Enter a floating point number -> ")
	fmt.Scan(&ogfloat64)
	TruncInt := truncate(ogfloat64)
	fmt.Printf("%d", TruncInt)
}