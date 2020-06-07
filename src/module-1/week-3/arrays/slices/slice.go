package main

import (
    "fmt"
    "sort"
    "strconv"
   )

func main() {
    slice := make([]int, 0, 5)

    for {
       fmt.Println("give a number")
       var userInput string
       _, err := fmt.Scanln(&userInput)
       if err == nil {
	   if userInput == "X" {
	       break
	    } else {
	       Int, _ := strconv.Atoi(userInput)
	       slice = append(slice, Int)
	       sort.Ints(slice)
         fmt.Printf("New slice: %v\n", slice)
	   }
       }
   }
}
