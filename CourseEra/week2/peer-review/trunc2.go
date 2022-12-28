package main

import (
	"fmt"
	"log"
)

func main(){
	var input float64
	var output int64
	for {
		fmt.Printf("Enter a floating point number (Ctrl+C to Quit):")
		_, err := fmt.Scan(&input)
		if err != nil {
			log.Fatalln(err)
			break
		}
		fmt.Printf("The input float is %f\n", input)
		output = int64(input)
		fmt.Printf("The truncated int is %d\n", output)
		input = 0.0
		output = 0
	}
}