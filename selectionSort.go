package main

import (
	"fmt"
)

func selectionSort(arr []int) {
	len := len(arr)
	for i := 0; i < len-1; i++ {
		minIndex := i
		for j := i + 1; j < len; j++ {
			if arr[j] < arr[minIndex] {
				arr[j], arr[minIndex] = arr[minIndex], arr[j]
			}
		}
	}
	fmt.Println("\nAfter SelectionSort")
	for _, val := range arr {
		fmt.Println(val)
	}
}
func main() {
	sample := []int{13, 14, 15, 2, 1}
	selectionSort(sample)
	sample = []int{30, 4, 15, 21, 11, 7, 28, -1, -3}
	selectionSort(sample)
}
