package main

import (
	"os"
	"fmt"
	"log"
	"bufio"
	"strings"
	// "io"
)

type name struct {
	fname string 
	lname string 
}

func main() {
	var filename string
	fmt.Printf("Enter the name of the file\n")
	fmt.Scan(&filename)
	totalData := make([]name, 100)
	file, err := os.OpenFile(filename + ".txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		f, l := strings.Split(line, " ")[0], strings.Split(line, " ")[1]
		newName := name{
			fname: f,
			lname: l,
		}
		totalData = append(totalData, newName)

		if err != nil {
				break
		}
	}

	for _, v := range totalData {
		os.Stdout.WriteString(v.fname + " " + v.lname)
	}
	// if err != io.EOF {
	// 	fmt.Printf("Failed to read because %v\n", err)
	// }
}


