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
	getName() string
}

type Cow struct{ name string }
type Bird struct{ name string }
type Snake struct{ name string }

func (c Cow) getName() string {
	return c.name
}
func (c Cow) Eat() {
	fmt.Println("grass")
}
func (c Cow) Move() {
	fmt.Println("walk")
}
func (c Cow) Speak() {
	fmt.Println("moo")
}

func (b Bird) getName() string {
	return b.name
}
func (b Bird) Eat() {
	fmt.Println("worms")
}
func (b Bird) Move() {
	fmt.Println("fly")
}
func (b Bird) Speak() {
	fmt.Println("peep")
}

func (s Snake) getName() string {
	return s.name
}
func (s Snake) Eat() {
	fmt.Println("mice")
}
func (s Snake) Move() {
	fmt.Println("slither")
}
func (s Snake) Speak() {
	fmt.Println("hiss")
}

func main() {

	animalLog := []Animal{}

	for {
		fmt.Print(">")
		input := bufio.NewReader(os.Stdin)
		input_req, _ := input.ReadString('\n')
		animal_req := strings.Split(input_req, " ")

		task := strings.TrimSpace(animal_req[0])
		task = strings.ToLower(task)

		first_param := strings.TrimSpace(animal_req[1])
		first_param = strings.ToLower(first_param)

		sec_param := strings.TrimSpace(animal_req[2])
		sec_param = strings.ToLower(sec_param)

		if task == "newanimal" {
			switch sec_param {
			case "cow":
				animalLog = append(animalLog, Cow{name: first_param})
				fmt.Println("Created it!!!")
				break
			case "bird":
				animalLog = append(animalLog, Bird{name: first_param})
				fmt.Println("Created it!!!")
				break
			case "snake":
				animalLog = append(animalLog, Snake{name: first_param})
				fmt.Println("Created it!!!")
				break
			default:
				fmt.Println("Invalid request")

			}
			fmt.Println(animalLog)

		} else if task == "query" {
			for i := 0; i < len(animalLog); i++ {
				if animalLog[i].getName() == first_param {
					switch sec_param {
					case "eat":
						animalLog[i].Eat()
						break
					case "move":
						animalLog[i].Move()
						break
					case "speak":
						animalLog[i].Speak()
						break
					default:
						fmt.Println("Invalid query")
					}
				} else {
					fmt.Println(("No such animal exist!!!"))
				}
			}
		}

	}
}
