// Write a program which prompts the user to enter integers and stores the integers in a sorted slice.
// The program should be written as a loop.
// Before entering the loop, the program should create an empty integer slice of size (length) 3.
// During each pass through the loop, the program prompts the user to enter an integer to be added to the slice.
// The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order.
// The slice must grow in size to accommodate any number of integers which the user decides to enter.
// The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.

package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {

	this_slice := make([]int, 0, 3)
	var user_input string

	for true {

		fmt.Print("Insert an Integer to be added to the slice. (X to discontinue): ")
		fmt.Scan(&user_input)

		// if the user input equates to "X" stop the program.
		if strings.Compare(user_input, "X") == 0 {
			break
		}

		// convert user input to integer
		user_input_proc, err := strconv.Atoi(user_input)

		// if user input is invalid, continue
		if err != nil {
			fmt.Printf("Invalid Integer input: %s. \n", user_input)
			continue
		}

		// add the elmt to the slice and sort it
		this_slice = append(this_slice, user_input_proc)
		sort.Sort(sort.IntSlice(this_slice))

		// print the list
		fmt.Printf("[")
		for _, v := range this_slice {
			fmt.Printf("%d ", v)
		}
		fmt.Printf("]\n")
	}

	fmt.Printf("Exiting the loop.")
}
