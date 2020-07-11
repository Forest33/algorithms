package main

import (
	"fmt"
	"sort"
)

func main() {
	lst := []int{2, 123, 443, 223, 3, 5, 6, 432, 1, 999, 333}
	sort.Ints(lst)
	fmt.Println(lst, "=>", binarySearch(lst, 432))

	lst = []int{2, 123, 443, 223, 3, 5, 6, 432, 1, 999, 333}
	fmt.Print(lst, " => ")
	bubbleSort(lst)
	fmt.Println(lst)
}
