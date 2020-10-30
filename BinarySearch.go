package main

import "fmt"

func BinarySearch(needle int, lst []int) bool {
	low := 0
	high := len(lst) - 1
	for low <= high {
		mid := (low + high) / 2
		if lst[mid] == needle {
			return true
		} else if lst[mid] < needle {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false

}
func main() {
	lst := []int{23, 56, 78, 79, 85, 89, 92, 95, 98}
	fmt.Println(BinarySearch(85, lst))
}
