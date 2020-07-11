package main

func binarySearch(data []int, key int) int {
	lo := 0
	hi := len(data) - 1
	for lo <= hi {
		mid := lo + (hi - lo) / 2
		if key < data[mid] { // left side
			hi = mid - 1
		} else if key > data[mid] { // right side
			lo = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
