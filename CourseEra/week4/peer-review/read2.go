package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	type Person struct {
		fname string
		lname string
	}
	var filePath string
	oneSlice := make([]Person, 0)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Input a file path:")
	if scanner.Scan() {
		filePath = scanner.Text()
	}

	file, _ := os.Open(filePath)
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		line := strings.TrimSpace(fileScanner.Text())
		splitLine := strings.Split(line, " ")
		oneSlice = append(oneSlice, Person{splitLine[0], splitLine[1]})
	}

	for _, person := range oneSlice {
		println(person.fname, person.lname)
	}
}
