package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"
)

type Animal struct {
	food string
	locomotion string
	noise string
}


func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}


func main() {
	var cow = Animal {
		 "grass",
		 "walk",
		 "moo",
	}
	var bird = Animal {
		 "worms",
		 "fly",
		 "peep",
	}
	var snake = Animal {
		 "mice",
		 "slither",
		 "hsss",
	}
	
	for {
		
		fmt.Print(">")
		
		reader := bufio.NewReader(os.Stdin)
		request, _ := reader.ReadString('\n')

		request = strings.TrimSpace(request) //cutting out leading and trailing spaces inlcuding \n, \t etc

		if request == "cow Eat" {
			cow.Eat()
		}
		if (request == "cow Move") {
			cow.Move()
		}
		if (request == "cow Speak") {
			cow.Speak()
		}
		if (request == "bird Eat") {
			bird.Eat()
		}
		if (request == "bird Move") {
			bird.Move()
		}
		if (request == "bird Speak") {
			bird.Speak()
		}
		if (request == "snake Eat") {
			snake.Eat()
		}
		if (request == "snake Move") {
			snake.Move()
		}
		if (request == "snake Speak") {
			snake.Speak()
		}
	}
}