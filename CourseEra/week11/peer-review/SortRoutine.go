package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	arr := make([]int, 0)
	c := make(chan []int)

	readInputs(&arr)
	if len(arr) < 4 {
		fmt.Println("Enter atleast 4 integers")
		os.Exit(1)
	}
	m := int64(math.Floor(float64(len(arr)) / 2.0))
	l1 := m / 2
	l2 := int64(math.Floor(float64(len(arr)-int(m)) / 2.0))

	go bubbleSort(arr[0:l1], c)
	go bubbleSort(arr[l1:m], c)
	go bubbleSort(arr[m:m+l2], c)
	go bubbleSort(arr[m+l2:len(arr)], c)
	arr1 := <-c
	arr2 := <-c
	arr3 := <-c
	arr4 := <-c
	go mergeSort(arr1, arr2, c)
	go mergeSort(arr3, arr4, c)
	arr5 := <-c
	arr6 := <-c
	go mergeSort(arr5, arr6, c)
	arr = <-c
	fmt.Println(arr)
}

func readInputs(arr *[]int) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter the sequence of numbers space seperated in single line")
	scanner.Scan()
	str := scanner.Text()
	intStr := strings.Split(str, " ")
	for _, s := range intStr {
		val, err := strconv.Atoi(s)
		if err == nil {
			*arr = append(*arr, val)
		}
	}
}

func bubbleSort(arr []int, c chan []int) {
	len := len(arr)
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-1-i; j++ {
			if arr[j] > arr[j+1] {
				swap(arr, j)
			}
		}
	}
	fmt.Println(arr)
	c <- arr
}

func mergeSort(arr1 []int, arr2 []int, c chan []int) {
	arr := make([]int, 0)

	i := 0
	j := 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			arr = append(arr, arr1[i])
			i++
		} else {
			arr = append(arr, arr2[j])
			j++
		}
	}

	for i < len(arr1) {
		arr = append(arr, arr1[i])
		i++
	}

	for j < len(arr2) {
		arr = append(arr, arr2[j])
		j++
	}
	fmt.Println(arr)
	c <- arr
}

func swap(arr []int, idx int) {
	arr[idx], arr[idx+1] = arr[idx+1], arr[idx]
}
