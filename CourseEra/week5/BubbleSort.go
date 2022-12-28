package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Swap(sli []int, index int) {
	this_value := sli[index]
	sli[index] = sli[index+1]
	sli[index+1] = this_value
}

func BubbleSort(sli []int) {
	len_slice := len(sli)
	for idx := 0; idx < len_slice; idx++ {
		for idy := 0; idy < len_slice-idx-1; idy++ {
			this_value := sli[idy]
			next_value := sli[idy+1]
			if this_value > next_value {
				Swap(sli, idy)
			}
		}
	}
}

func GetDataFromUser() []int {
	input_list := make([]int, 0, 10)
	var user_input string
	for idx := 0; idx < 10; idx++ {
		fmt.Print("Insert an Integer. (X to discontinue): ")
		fmt.Scan(&user_input)
		if strings.Compare(strings.Trim(user_input, " "), "X") == 0 {
			break
		}
		int_user_input, err := strconv.Atoi(user_input)
		if err != nil {
			fmt.Printf("Invalid integer provided: '%s'. ", user_input)
			idx-- // to offset the iteration for invalid insertion
			continue
		}

		input_list = append(input_list, int_user_input)
	}
	return input_list
}

func main() {
	list_of_integers := GetDataFromUser()
	BubbleSort(list_of_integers)
	fmt.Println(list_of_integers)
}
