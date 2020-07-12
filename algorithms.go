package main

import (
	"fmt"
	"sort"
)

func main() {
	data := []int{2, 123, 443, 223, 3, 5, 6, 432, 1, 999, 333}
	sort.Ints(data)
	fmt.Println(data, "=>", BinarySearch(data, 432))

	data = []int{2, 123, 443, 223, 3, 5, 6, 432, 1, 999, 333}
	fmt.Print(data, " => ")
	BubbleSort(data)
	fmt.Println(data)

	data_str := []string{"zz", "yy", "dd", "cc", "uu", "zu", "bb", "cc", "aa", "ee", "xx", "zz"}
	fmt.Print(data_str, " => ")
	RadixStringSort(data_str, 2)
	fmt.Println(data_str)

	data = []int{2, 443, 123, 223, 7890123, 3, 5, 6, 432, 1, 999, 333}
	fmt.Print(data, " => ")
	fmt.Println(MergeSort(data))
}
