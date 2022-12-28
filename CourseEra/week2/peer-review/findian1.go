package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
	"bufio"
	"os"
)

func main() {
	var stringToCheck string
	fmt.Printf("Enter a string: ")
	for {
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			stringToCheck = scanner.Text()
			if strings.TrimSpace(stringToCheck) == "" { 
				fmt.Printf("You entered an empty string. Try again: ") 
				continue
			} else { break }
		}
	}

	var lowerString string = strings.ToLower(stringToCheck)
	if strings.Contains(lowerString, "a") && strings.IndexAny(lowerString, "i") == 0 && strings.LastIndex(lowerString, "n") == (utf8.RuneCountInString(lowerString)-1) {
		fmt.Printf("\nFound!")
	} else {
		fmt.Printf("\nNot Found!")
	}
}