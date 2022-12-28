package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func GenDisplaceFn(acc, vel, disp float64) func(float64) float64 {
	return func(time float64) float64 {
		return 0.5*acc*math.Pow(time, 2) + vel*time + disp
	}
}

func PromptUserForPrimaryInputs() (float64, float64, float64) {
	var acc, vel, disp float64
	fmt.Print("Insert acceleration: ")
	fmt.Scan(&acc)

	fmt.Print("Insert initial velocity: ")
	fmt.Scan(&vel)

	fmt.Print("Insert initial displacement: ")
	fmt.Scan(&disp)

	return acc, vel, disp
}

func PromptUserForSecondaryInput() (float64, bool) {
	var time string
	fmt.Print("Insert time. (Press X to discontinue): ")
	fmt.Scan(&time)

	to_break := false
	time_float := 0.0
	if strings.Compare(strings.Trim(time, " "), "X") == 0 {
		to_break = true
	} else {
		var err error
		time_float, err = strconv.ParseFloat(time, 64)
		if err != nil {
			fmt.Printf("Invalid value detected: %s.\n", time)
			time_float, to_break = PromptUserForSecondaryInput()
		}
	}
	return time_float, to_break

}

func main() {
	acc, vel, disp := PromptUserForPrimaryInputs()
	DisplaceFn := GenDisplaceFn(acc, vel, disp)

	for true {
		time, to_break := PromptUserForSecondaryInput()
		if to_break {
			break
		}
		fmt.Printf("Displacement after %f seconds: %f. \n", time, DisplaceFn(time))
	}
}
