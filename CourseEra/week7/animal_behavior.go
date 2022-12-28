package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	eat   string
	move  string
	speak string
}

func (animal Animal) Eat() {
	fmt.Println(animal.eat)
}

func (animal Animal) Move() {
	fmt.Println(animal.move)
}

func (animal Animal) Speak() {
	fmt.Println(animal.speak)
}

func main() {
	cow := Animal{eat: "grass", move: "walk", speak: "moo"}
	bird := Animal{eat: "worms", move: "fly", speak: "peep"}
	snake := Animal{eat: "mice", move: "slither", speak: "hsss"}

	this_scanner := bufio.NewScanner(os.Stdin)
	for true {
		fmt.Print("> ")
		this_scanner.Scan()
		input_str := this_scanner.Text()

		input_str_array := strings.Split(input_str, " ")
		var this_animal Animal
		switch input_str_array[0] {
		case "cow":
			this_animal = cow
		case "bird":
			this_animal = bird
		case "snake":
			this_animal = snake
		default:
			fmt.Printf("Unrecognized animal: %s.\n", input_str_array[0])
			continue
		}

		switch input_str_array[1] {
		case "eat":
			this_animal.Eat()
		case "move":
			this_animal.Move()
		case "speak":
			this_animal.Speak()
		default:
			fmt.Printf("Undefined action: %s for animal: %s.\n", input_str_array[1], input_str_array[0])
		}
	}
}
