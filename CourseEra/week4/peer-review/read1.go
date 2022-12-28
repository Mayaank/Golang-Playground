package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func main() {
	var fileName string
	personSli := make([]Person, 0, 3)

	fmt.Printf("Enter filename: ")
	fmt.Scan(&fileName)

	data, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		larr := strings.Split(line, " ")
		personSli = append(personSli, Person{
			fname: larr[0],
			lname: larr[1],
		})
	}

	for _, v := range personSli {
		fmt.Println(v)
	}

}
