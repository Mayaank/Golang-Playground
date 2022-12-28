package main

import "fmt"

const (
	EnterAccelerationPrompt string = "Enter Acceleration:"
	EnterVelocityPrompt     string = "Enter Velocity:"
	EnterDisplacementPrompt string = "Enter initial displacement:"
	EnterTimePrompt         string = "Enter time to calculate displace:"
	DisplacePrompt          string = "The displace is %v.\r\n"
	ExitPrompt              string = "Enter X to exit, otherwise to continue:"

	ExitValue string = "X"
)

type Animal struct {
	food, locomotion, noise string
}

// func (animal *Animal) Init(food, locomotion, noise string) {
// 	animal.food = food
// 	animal.locomotion = locomotion
// 	animal.noise = noise
// }

func (animal *Animal) Eat() {
	fmt.Println(animal.food)
}
func (animal *Animal) Move() {
	fmt.Println(animal.locomotion)
}
func (animal *Animal) Speak() {
	fmt.Println(animal.noise)
}

func main() {
	var (
		animals = map[string]*Animal{
			"cow":   {food: "grass", locomotion: "walk", noise: "moo"},
			"bird":  {food: "worms", locomotion: "fly", noise: "peep"},
			"snake": {food: "mice", locomotion: "slither", noise: "hsss"},
		}
		informationFunction = map[string]func(*Animal){
			"eat":   func(animal *Animal) { animal.Eat() },
			"move":  func(animal *Animal) { animal.Move() },
			"speak": func(animal *Animal) { animal.Speak() },
		}
		animalName string
		animalInfo string
	)

	for {
		fmt.Println("Enter the animal name (cow,bird,snake) and its information (eat,move,speak):")
		fmt.Print(">")
		fmt.Scan(&animalName, &animalInfo)
		animal, animalFound := animals[animalName]
		infoFunc, infoFund := informationFunction[animalInfo]

		if !animalFound || !infoFund {
			fmt.Println("Invalid animal or information name")
			continue
		}

		fmt.Printf("The result is:")
		infoFunc(animal)
	}
}
