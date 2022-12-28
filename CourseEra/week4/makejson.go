// Write a program which prompts the user to first enter a name, and then enter an address.
// Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively.
// Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	// Prompt the user to provide the name
	fmt.Printf("Name: ")
	scanner.Scan()
	user_name := scanner.Text()

	// Prompt the user to provide the address
	fmt.Printf("Address: ")
	scanner.Scan()
	user_addr := scanner.Text()

	// Construct the map
	user_info := map[string]string{
		"name":    user_name,
		"address": user_addr,
	}

	// Convert the map to JSON []byte
	json_bytes, _ := json.Marshal(user_info)

	// Print the JSON []byte and it's string representation
	fmt.Printf("JSON object as string: %s. \n", string(json_bytes))
	fmt.Printf("JSON object as []byte: %v. \n", json_bytes)

}
