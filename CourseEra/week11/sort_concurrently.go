// Write a program to sort an array of integers.
// The program should partition the array into 4 parts, each of which is sorted by a different goroutine.
// Each partition should be of approximately equal size.
// Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

// The program should prompt the user to input a series of integers.
// Each goroutine which sorts ¼ of the array "should print the subarray that it will sort".
// When sorting is complete, the main goroutine "should print the entire sorted list."

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func sort_array(sli []int) {
	/* Sort slice of Ints in place*/
	sort.Ints(sli)
	fmt.Println(sli) // Prints the sorted slice by each of the goroutine.
	wg.Done()
}

func create_subslices(num_slices int, main_slice []int) [][]int {
	/* Generates an array of almost equal subslices based on the parent slice */
	slice_length := len(main_slice)
	subslice_list := make([][]int, num_slices)

	l_idx := 0
	for itr := 1; itr <= num_slices; itr++ {
		r_idx := slice_length * itr / num_slices
		subslice_list[itr-1] = main_slice[l_idx:r_idx]
		l_idx = r_idx
	}
	return subslice_list
}

func merge(left, right []int) []int {
	/* Merge two sorted arrays into one. */
	merged_array := make([]int, 0)
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			merged_array = append(merged_array, left[0])
			left = left[1:]
		} else {
			merged_array = append(merged_array, right[0])
			right = right[1:]
		}
	}

	if len(left) > 0 {
		merged_array = append(merged_array, left...)
	}

	if len(right) > 0 {
		merged_array = append(merged_array, right...)
	}

	return merged_array
}

func main() {

	// Ask the user for input
	fmt.Printf("Insert all (integer) elements to be sorted> ")
	this_scanner := bufio.NewScanner(os.Stdin)
	this_scanner.Scan()
	user_input := this_scanner.Text()

	// Process input data and convert each element, if possible,  to integer
	array_elements_str := strings.Split(strings.Trim(user_input, " "), " ")
	var array_elements_int []int
	for _, elem := range array_elements_str {
		int_elem, err := strconv.Atoi(elem)
		if err != nil {
			fmt.Printf("Invalid element: %s in the inputs.\n", elem)
			return
		}
		array_elements_int = append(array_elements_int, int_elem)

	}

	// Split the Slice into 4 subslices
	num_of_divisions := 4
	all_subslices := create_subslices(num_of_divisions, array_elements_int)

	// Sort each slice using different goroutines concurrently
	wg.Add(num_of_divisions)
	for _, this_subslice := range all_subslices {
		go sort_array(this_subslice) // All subslices are sorted in place
	}
	wg.Wait() // Wait for all the goroutines to get completed

	// Merge the individually sorted subslices into a single slice
	merged_array := merge(merge(all_subslices[0], all_subslices[1]), merge(all_subslices[2], all_subslices[3]))

	fmt.Println(merged_array) // Print the sorted array all in one place.
}
