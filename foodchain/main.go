package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
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

func processRequest(animal_req []string) {

	animalMap := map[string]Animal{
		"cow":   Animal{"grass", "walk", "moo"},
		"bird":  Animal{"worms", "fly", "peep"},
		"snake": Animal{"mice", "slither", "hsss"},
	}

	animal := strings.TrimSpace(animal_req[0])
	animal = strings.ToLower(animal)

	action := strings.TrimSpace(animal_req[1])
	action = strings.ToLower(action)

	switch action {
	case "eat":
		animalMap[animal].Eat()
		break
	case "move":
		animalMap[animal].Move()
		break
	case "speak":
		animalMap[animal].Speak()
		break
	default:
		fmt.Println("No such action....please try again")
	}
}

func main() {

	for {
		fmt.Print(">")
		input := bufio.NewReader(os.Stdin)
		animal_input, _ := input.ReadString('\n')
		animal_req := strings.Split(animal_input, " ")

		processRequest(animal_req)

	}
}
