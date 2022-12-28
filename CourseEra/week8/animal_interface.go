package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}

func (cow Cow) Eat() {
	fmt.Println("grass")
}

func (cow Cow) Move() {
	fmt.Println("walk")
}

func (cow Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct{}

func (bird Bird) Eat() {
	fmt.Println("worms")
}

func (bird Bird) Move() {
	fmt.Println("fly")
}

func (bird Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct{}

func (snake Snake) Eat() {
	fmt.Println("mice")
}

func (snake Snake) Move() {
	fmt.Println("slither")
}

func (snake Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	var command, animalName, thirdArgument string
	animalMap := map[string]Animal{
		"cow":   Cow{},
		"bird":  Bird{},
		"snake": Snake{},
	}

	informationFunction := map[string]func(Animal){
		"eat":   func(animal Animal) { animal.Eat() },
		"move":  func(animal Animal) { animal.Move() },
		"speak": func(animal Animal) { animal.Speak() },
	}

	thisRecord := make(map[string]Animal)

	for true {
		fmt.Print("> ")
		fmt.Scan(&command, &animalName, &thirdArgument)

		if command == "newanimal" {
			anotherAnimal, animalFound := animalMap[thirdArgument]
			if !animalFound {
				fmt.Println("Invalid animal name.")
				continue
			}
			thisRecord[animalName] = anotherAnimal
			fmt.Println("Created it!")
		} else if command == "query" {
			anotherAnimal, animalFound := thisRecord[animalName]
			action, actionFound := informationFunction[thirdArgument]
			if !animalFound || !actionFound {
				fmt.Println("Invalid animalName or action name.")
				continue
			}
			action(anotherAnimal)

		} else {
			fmt.Printf("Invalid command: %s. \n", command)
		}
	}
}
