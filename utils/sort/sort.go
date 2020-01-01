package sort

// BubbleSort - sort a slice of integers in ascending order
func BubbleSort(elements []int) {
	walkThrough := true
	for walkThrough {
		walkThrough = false

		for i := 0; i < len(elements)-1; i++ {
			if elements[i] > elements[i+1] {
				walkThrough = true
				elements[i], elements[i+1] = elements[i+1], elements[i]
			}
		}
	}
}
