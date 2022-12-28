package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// inputIntegerList() will prompt the user to input a list space separated integers.
// The return will be a slice pointing to the list of integers.
func inputIntegerList() []int {
	fmt.Println("Enter a list of integers on a single line, space seperated.")
	theList := make([]int, 0, 5)
	cmdline := bufio.NewScanner(os.Stdin)
	if !cmdline.Scan() {
		fmt.Println("error reading integers from stdin: ", cmdline.Err())
		return theList
	}

	elems := strings.Split(cmdline.Text(), " ")
	if len(elems) == 1 && elems[0] == "" {
		return theList // empty input/list
	}
	for _, e := range elems {
		i, err := strconv.Atoi(e)
		if err != nil {
			fmt.Printf("failed to convert str to int: %s: %s\n", e, err)
			continue
		}
		theList = append(theList, i)
	}
	return theList
}

// Swap element list[i] with list[i+1]
func Swap(list []int, index int) {
	tmp := list[index]
	list[index] = list[index+1]
	list[index+1] = tmp
	return
}

// BubbleSort will sort a slice pointing at a 'list' of integers.  Upon return the slice will
// sorted in ascending order.
func BubbleSort(list []int) {
	// if len(list) == 0 {
	// 	return
	// }
	for tail := len(list) - 1; tail > 0; tail = tail - 1 {
		for i := 0; i < tail; i++ {
			if list[i] > list[i+1] {
				Swap(list, i)
			}
		}
	}
	return
}

// displayNumbers prints the sorted list to stdout on a single line
func displayNumbers(theList []int) {
	fmt.Printf("The sorted list of %v integers: ", len(theList))
	for _, i := range theList {
		fmt.Print(i, " ")
	}
	fmt.Println()
}

func main() {
	theList := inputIntegerList()
	BubbleSort(theList)
	displayNumbers(theList)
}
