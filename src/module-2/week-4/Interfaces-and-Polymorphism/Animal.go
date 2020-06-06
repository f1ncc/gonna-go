package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)



type Animal interface {
	Eat()
	Move()
	Speak()
}


type Cow struct {
	food, locomotion, noise string
}

func (a Cow) Eat() {
	fmt.Println(a.food)
}

func (a Cow) Speak() {
	fmt.Println(a.noise)
}

func (a Cow) Move() {
	fmt.Println(a.locomotion)
}


type Bird struct {
	food, locomotion, noise string
}

func (a Bird) Eat() {
	fmt.Println(a.food)
}

func (a Bird) Speak() {
	fmt.Println(a.noise)
}

func (a Bird) Move() {
	fmt.Println(a.locomotion)
}


type Snake struct {
	food, locomotion, noise string
}

func (a Snake) Eat()  {
	fmt.Println(a.food)
}

func (a Snake) Speak()  {
	fmt.Println(a.noise)
}

func (a Snake) Move()  {
	fmt.Println(a.locomotion)
}


func main() {
	var newAnimal Animal

	animals := make(map[string]Animal)

	for {

		fmt.Printf(">")

		reader := bufio.NewReader(os.Stdin)
		command, _ := reader.ReadString('\n')

		command = strings.TrimSpace(command)
		split_command := strings.Split(command, " ")

		if split_command[0] == "newanimal" { 	

			switch split_command[2] {		
				case "cow":
					newAnimal = Cow{"Grass","Walk","moo"}
				
				case "snake":
					newAnimal = Snake{"mice","slither","hsss"}
				
				case "bird":
					newAnimal = Bird{"worm","fly","peep"}
			}

			animals[split_command[1]] = newAnimal 
			fmt.Println("Created a new animal")

		} else if split_command[0] == "query" {

			switch split_command[2] {
				case "eat":
					animals[split_command[1]].Eat()

				case "move":
					animals[split_command[1]].Move()

				case "speak":
					animals[split_command[1]].Speak()
			}

		}

	}
}