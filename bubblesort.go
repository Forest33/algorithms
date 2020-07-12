package main

func BubbleSort(data []int) {
	n := len(data)
	for {
		var exchange = false
		for i := 0; i < n; i++ {
			if i+1 < n && data[i] > data[i+1] {
				data[i+1], data[i] = data[i], data[i+1]
				exchange = true
			}
		}
		if !exchange {
			return
		}
	}
}
