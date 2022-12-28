package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var ints = make([]int, 0, 3)

	for {
		fmt.Print("Enter an integer: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Errorf("Error: %v", err)
			os.Exit(1)
		}
		input = strings.Replace(input, "\n", "", -1)
		if input == "X" {
			break
		}
		if num, err2 := strconv.Atoi(strings.SplitN(input, " ", 1)[0]); err2 == nil {
			ints = append(ints, num)
			sort.Ints(ints)
			fmt.Println("Sorted Ints slice: ", ints)
		}
	}
}
