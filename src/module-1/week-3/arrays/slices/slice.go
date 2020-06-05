package main

import (
	"fmt"
	// "strconv"
	"sort"
	"log"
	"os"
	"bufio"
)

// func sort(a []int, n int) bool {
// 	if (a[n-1] > a[n-2]) {
// 		return true
// 	} else {
// 		return false
// 	}
// }

// func CheckInput(s string) bool {
// 	if (s[0] == 'X') {
// 		return true
// 	} else {
// 		return false
// 	}
// }

func main() {
	sli := make([]int, 0, 3)
	for {
		// var a int 
		// fmt.Scan(&a)
		// if(a == 'X') {
			
		// }
		input := bufio.NewReader(os.Stdin)
		a, _, err := input.ReadRune()
		if err != nil {
			log.Fatal(err)
		}
		if a == 'X' {
			break
		}
		// if(CheckInput(strconv.Itoa(a))) {
		// 	return nil
		// }
		sli = append(sli, int(a)-48)
		// if(len(sli) >= 2) {
		// 	if(!sort(sli, len(sli))) {
		// 		temp := sli[len(sli)-2]
		// 		sli[len(sli)-2] = sli[len(sli)-1] 
		// 		sli[len(sli)-1] = temp
		// 	}
		// }
		// sli = sort.Ints(sli)
		sort.Ints(sli)
		fmt.Printf("%d\n",sli)
	}
}