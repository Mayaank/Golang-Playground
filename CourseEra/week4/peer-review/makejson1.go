package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name string
	var addr string
	fmt.Printf("Enter name: ")
	fmt.Scan(&name)
	fmt.Printf("Enter address: ")
	fmt.Scan(&addr)

	idMap := map[string]string{
		"name":    name,
		"address": addr,
	}

	barr, _ := json.Marshal(idMap)

	fmt.Print(string(barr))

}
