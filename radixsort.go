package main

func RadixStringSort(data []string, w int) {
	n := len(data)
	R := 256
	aux := make([]string, n)

	for d := w - 1; d >= 0; d-- {
		count := make([]int, R+1)
		for i := 0; i < n; i++ {
			count[data[i][d]+1]++
		}

		for r := 0; r < R; r++ {
			count[r+1] += count[r]
		}

		for i := 0; i < n; i++ {
			aux[count[data[i][d]]] = data[i]
			count[data[i][d]]++
		}

		for i := 0; i < n; i++ {
			data[i] = aux[i]
		}
	}
}