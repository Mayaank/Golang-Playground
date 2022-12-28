// Write a program which prompts the user to enter a string.
// The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’.
// The program should print
//      “Found!” if the entered string starts with the character ‘i’, ends with the character ‘n’, and contains the character ‘a’.
// The program should print “Not Found!” otherwise.
// The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.

// Examples: The program should print “Found!” for the following example entered strings, “ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”.
// The program should print “Not Found!” for the following strings, “ihhhhhn”, “ina”, “xian”.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Enter the input string: ")

	// Read input from the user. Using bufio, since the input can contain whitespaces as well.
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input_str := scanner.Text()

	// Convert the input to lower-case
	input_str = strings.ToLower(input_str)

	// Construct the output message based on the condition
	var output_message string
	if strings.HasPrefix(input_str, "i") && strings.HasSuffix(input_str, "n") && strings.Contains(input_str, "a") {
		output_message = "Found!"
	} else {
		output_message = "Not Found!"
	}
	fmt.Printf(output_message)
}
