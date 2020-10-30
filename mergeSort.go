package main

import (
	"fmt"
	"math/rand"
	"time"
)

func MergeSort(lst []int) []int {

	if len(lst) < 2 {
		return lst
	}
	mid := (len(lst)) / 2
	return Merge(MergeSort(lst[:mid]), MergeSort(lst[mid:]))
}

func Merge(left, right []int) []int {

	size, i, j := len(left)+len(right), 0, 0
	slice := make([]int, size, size)
	count := 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			slice[count] = left[i]
			count, i = count+1, i+1
		} else {
			slice[count] = right[j]
			count, j = count+1, j+1
		}
	}
	for i < len(left) {
		slice[count] = left[i]
		count, i = count+1, i+1
	}
	for j < len(right) {
		slice[count] = right[j]
		count, j = count+1, j+1
	}

	return slice
}
func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}
func main() {
	slice := generateSlice(10)
	fmt.Println("unsorted data:", slice)
	fmt.Println("Sorted data:", MergeSort(slice))
}
