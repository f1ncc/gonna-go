package main

import (
	"fmt"
)

func swap (i *int, j *int) {
	temp := *j
	*j = *i
	*i = temp
}

// using slices instead of arrays for performance reasons
func BubbleSort(sli []int) {
	for i := len(sli); i > 0; i-- {
		 for j := 1; j < i; j++ {
				if sli[j-1] > sli[j] {
					 swap(&sli[j-1], &sli[j])
				}
		}
	}
}

func main() {
	var arr [10]int
	for i := 0; i < 10; i++ {
		fmt.Scan(&arr[i])
	}
	slice := arr[:]
	BubbleSort(slice)
	for _, v := range slice {
		fmt.Printf("%d ", v)
	}
}