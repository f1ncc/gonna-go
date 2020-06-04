package main

import (
	"fmt"
	"strings"
)

func Found() {
	fmt.Printf("Found")
}

func NotFound() {
	fmt.Printf("Not Found")
}

func checkString(s string) bool {
	if ((s[0] == 'i' || s[0] == 'I') && (s[len(s)-1] == 'n' || s[len(s)-1] == 'N')){
		if(strings.ContainsRune(s,97) || strings.ContainsRune(s,65)) {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func main() {
	var str string
	fmt.Printf("give a string\n")
	fmt.Scan(&str)
	fmt.Printf("\n")
	if(checkString(str)) {
		Found()
	} else {
		NotFound()
	}
}
