package main

import "fmt"

func main() {
	// var sat = new(int)
	// var i = *sat
	// var addr *int = &i
	// fmt.Printf("%v\n%v", sat, addr)
	for {
		var r int
		fmt.Scan(&r)
		fmt.Printf("%d\n",r/4)
	}
}