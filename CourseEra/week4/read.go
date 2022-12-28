// Write a program which reads information from a file and represents it in a slice of structs.
// Assume that there is a text file which contains a series of names.
// Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.

// Your program will define a name struct which has two fields,
// 		fname for the first name, and
//		lname for the last name.
// Each field will be a string of size 20 (characters).

// Your program should prompt the user for the name of the text file.
// Your program will successively read each line of the text file and
// create a struct which contains the first and last names found in the file.
// Each struct created will be added to a slice, and after all lines have been read from the file,
// your program will have a slice containing one struct for each line in the file.
// After reading all lines from the file,
// your program should iterate through your slice of structs and print the first and last names found in each struct.

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	// Define a struct for the user
	type Person struct {
		fname string
		lname string
	}

	// Initialize an array of such structs
	person_list := make([]Person, 0, 10)

	// Prompt the user for the filename
	fmt.Printf("Filename: ")
	var filename string
	fmt.Scan(&filename)

	// Read the content of the entire file
	file_content, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Printf("Error reading file: %s", filename)
	} else {
		// Split the data based on line breaks and iterate over each line
		for _, _line := range strings.Split(string(file_content), "\n") {
			_line := strings.Trim(_line, " ")

			// Continue to the next line if there is no data in the current line
			if _line == "" {
				continue
			}

			// Add person's data in the struct
			this_person_name := strings.Split(_line, " ")
			person_list = append(
				person_list,
				Person{
					fname: strings.Trim(this_person_name[0], " "),
					lname: strings.Trim(this_person_name[1], " "),
				})
		}

		// Iterate over each line and display each entry in the array
		for _, this_person := range person_list {
			fmt.Printf("%s %s\n", this_person.fname, this_person.lname)
		}
	}
}
