// Write a program which prompts the user to enter a floating point number
// prints the integer which is a truncated version of the floating point number that was entered.
// Truncation is the process of removing the digits to the right of the decimal place.

package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Insert a floating point number: ")

	var user_prompt float32
	fmt.Scan(&user_prompt)
	fmt.Printf("%d", int(user_prompt))

}
