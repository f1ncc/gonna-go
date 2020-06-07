package main

import (
	"fmt"
	"bufio"
	"os"
	"sync"
	"strings"
	"strconv"
	// "time"
	"sort"
)


func sort_slice(sli []int, channel chan []int, wg *sync.WaitGroup) {
	fmt.Printf("this is the slice to be sorted in this goroutine: %v\n", sli)
	sort.Ints(sli)
	channel <- sli
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	inputArray := make([]int,0,10)
	fmt.Println("give a slice")
	reader := bufio.NewReader(os.Stdin)
	inputString, _ := reader.ReadString('\n')
	inputString = strings.TrimSpace(inputString)
	StringArray := strings.Split(inputString, " ")
	for _, v := range StringArray {
		int, err := strconv.Atoi(v)
		if err == nil {
			inputArray = append(inputArray,int)
		}
	}
	// fmt.Printf("%v\n", inputArray)
	
	size := len(inputArray)
	channel := make(chan []int)
	div := size / 4
	wg.Add(4)
	go sort_slice(inputArray[0:div], channel, &wg)
	go sort_slice(inputArray[div:2*div], channel, &wg)
	go sort_slice(inputArray[2*div:3*div], channel, &wg)
	go sort_slice(inputArray[3*div:size], channel, &wg)
	// time.Sleep(10000 * time.Millisecond)
	var sortedSlice []int

	slice1 := <- channel
	slice2 := <- channel
	slice3 := <- channel
	slice4 := <- channel

	wg.Wait()
	
	appendSlice(&slice1, &sortedSlice)
	appendSlice(&slice2, &sortedSlice)
	appendSlice(&slice3, &sortedSlice)
	appendSlice(&slice4, &sortedSlice)
	// sortedSlice = append(sortedSlice, slice1...)
	// sortedSlice = append(sortedSlice, slice2...)
	// sortedSlice = append(sortedSlice, slice3...)
	// sortedSlice = append(sortedSlice, slice4...)

	sort.Ints(sortedSlice)
	fmt.Printf("%v", sortedSlice)
}

func appendSlice (ss *[]int, s *[]int) {
	for _, v := range *ss {
		*s = append(*s,v)
	}
}