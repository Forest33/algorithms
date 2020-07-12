package main

func merge(a, b []int) []int {
	c := make([]int, len(a)+len(b))
	i := 0
	j := 0
	for k := 0; k < len(c); k++ {
		if i >= len(a) {
			c[k] = b[j]
			j++
		} else if j >= len(b) {
			c[k] = a[i]
			i++
		} else if a[i] <= b[j] {
			c[k] = a[i]
			i++
		} else {
			c[k] = b[j]
			j++
		}
	}
	return c
}

func MergeSort(input []int) []int {
	N := len(input)
	if N <= 1 {
		return input
	}
	a := make([]int, N/2)
	b := make([]int, N-N/2)
	for i := 0; i < len(a); i++ {
		a[i] = input[i]
	}
	for i := 0; i < len(b); i++ {
		b[i] = input[i+N/2]
	}
	return merge(MergeSort(a), MergeSort(b))
}
