package main

import (
	"fmt"
	"math"
)

func isFloatInt(floatValue float64) bool {
	return math.Mod(floatValue, 1.0) == 0
}

func main() {
	var fNumber float64

	fmt.Printf("Enter a floating number:\t")

	for {
		fmt.Scan(&fNumber)
		if isFloatInt(fNumber) {
			fmt.Printf("You did not enter a floating number.\nTry again:\t")
		}
		if !isFloatInt(fNumber) {
			break
		}
	}

	fmt.Printf("\nTruncated number: %.0f", math.Trunc(fNumber)) //or int64(fNumber)
}
