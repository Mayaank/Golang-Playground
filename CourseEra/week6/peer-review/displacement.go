package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(acceleration, initialVelocity, initialDisplacement float64) func(float64) float64 {

	fn := func(t float64) float64 {
		return ((0.5 * acceleration * math.Pow(t, 2)) + (initialVelocity * t) + initialDisplacement)
	}

	return fn
}

func main() {

	var a, v, s, t float64
	fmt.Println("Enter the acceleration:")
	fmt.Scanln(&a)
	fmt.Println("Enter the initial velocity:")
	fmt.Scanln(&v)
	fmt.Println("Enter the initial displacement:")
	fmt.Scanln(&s)

	displacement := GenDisplaceFn(a, v, s)

	fmt.Println("Enter the time:")
	fmt.Scanln(&t)

	fmt.Println("displacement:", displacement(t))

}
