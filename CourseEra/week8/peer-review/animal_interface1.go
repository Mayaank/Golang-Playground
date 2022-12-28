package main

import (
	"fmt"
)

const eatRequest = "eat"
const moveRequest = "move"
const speakRequest = "speak"

const cow = "cow"
const bird = "bird"
const snake = "snake"

const newAnimalCommand = "newanimal"
const queryCommand = "query"

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}

func (c Cow) Eat() {
	fmt.Println("grass")
}

func (c Cow) Move() {
	fmt.Println("walk")
}

func (c Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct{}

func (b Bird) Eat() {
	fmt.Println("worms")
}

func (b Bird) Move() {
	fmt.Println("fly")
}

func (b Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct{}

func (s Snake) Eat() {
	fmt.Println("mice")
}

func (s Snake) Move() {
	fmt.Println("slither")
}

func (s Snake) Speak() {
	fmt.Println("hsss")
}

func addAnimal(animals map[string]Animal, name, kind string) {
	switch kind {
	case cow:
		animals[name] = Cow{}
	case bird:
		animals[name] = Bird{}
	case snake:
		animals[name] = Snake{}
	default:
		fmt.Println("the animal is not exists")
		return
	}

	fmt.Println("Created it!")
}

func executeAction(animals map[string]Animal, name, action string) {
	selectedAnimal, ok := animals[name]

	if ok == false {
		fmt.Println("we don't have animal with this name ")
		return
	}

	switch action {
	case eatRequest:
		selectedAnimal.Eat()
	case speakRequest:
		selectedAnimal.Speak()
	case moveRequest:
		selectedAnimal.Move()
	default:
		fmt.Println("this kind of action we are not support")
	}
}

func main() {
	animals := map[string]Animal{}

	for true {

		fmt.Print("> ")

		var cmd, name, kind string

		_, err := fmt.Scanln(&cmd, &name, &kind)

		if err != nil {
			fmt.Println("something wrong, please try again")
			continue
		}

		if cmd == newAnimalCommand {
			addAnimal(animals, name, kind)
		} else if cmd == queryCommand {
			executeAction(animals, name, kind)
		} else {
			fmt.Println("the command is not supported")
		}
	}
}
