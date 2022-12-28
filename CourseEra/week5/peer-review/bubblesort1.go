// Write a Bubble Sort program in Go. The program
// should prompt the user to type in a sequence of up to 10 integers. The program
// should print the integers out on one line, in sorted order, from least to
// greatest. Use your favorite search tool to find a description of how the bubble
// sort algorithm works.

// As part of this program, you should write a
// function called BubbleSort() which
// takes a slice of integers as an argument and returns nothing. The BubbleSort() function should modify the 
// slice so that the elements are in sorted order.

// A recurring operation in the bubble sort algorithm is
// the Swap operation which swaps the position of two adjacent elements in the
// slice. You should write a Swap() function which performs this operation. Your Swap()
// function should take two arguments, a slice of integers and an index value i which
// indicates a position in the slice. The Swap() function should return nothing, but it should swap
// the contents of the slice in position i with the contents in position i+1.
package main

import(
	"fmt"
)

func Swap(input_slice []int, index int){
	temp := input_slice[index]
	input_slice[index] = input_slice[index+1]
	input_slice[index+1] = temp
}

func BubbleSort(input_slice []int){
	for index:=0;index<(len(input_slice)-1); index++{
		for index2:=0;index2<(len(input_slice)-1 - index); index2++{
			if input_slice[index2] > input_slice[index2+1]{
				Swap(input_slice,index2)
			}
		}
	}
}

func main(){
	var input_number int
	integer_slice := make([]int,10)
	for input_index:=0;input_index<10;input_index++{
		fmt.Printf("enter a number '%d': ",(input_index + 1))
		fmt.Scan(&input_number)
		integer_slice[input_index] = input_number
	}
	BubbleSort(integer_slice)
	fmt.Println(integer_slice)
}