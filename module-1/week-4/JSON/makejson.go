package main

import (
	"fmt"
	"encoding/json"
	"log"
	"os"
)

func main() {
	
	var (
		n string 
		addr string
	)
	fmt.Printf("enter a name: ")
	fmt.Scan(&n)
	fmt.Printf("enter address: ")
	fmt.Scan(&addr)
	
	sample := map[string]string {
		"name": n,
		"address": addr,
	}

	jsonObj, err := json.Marshal(sample)
	
	if(err != nil) {
		log.Fatal(err)
	}
	os.Stdout.Write(jsonObj)
	
}
