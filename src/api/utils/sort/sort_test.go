package sort

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	// Init
	elements := []int{9, 7, 5, 6, 3, 2, 1, 6, 4}

	// Execution
	BubbleSort(elements)

	// Validation
	if elements[0] != 1 {
		t.Error("First element should be 1")
	}

	if elements[len(elements)-1] != 9 {
		t.Error("Last element should be 9")
	}
}

func TestSortIncreasingOrder(t *testing.T) {
	// Init
	elements := []int{9, 7, 5, 6, 3, 2, 1, 6, 4}

	// Execution
	Sort(elements)

	// Validation
	if elements[0] != 1 {
		t.Error("First element should be 1")
	}

	if elements[len(elements)-1] != 9 {
		t.Error("Last element should be 9")
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	elements := GetElements(100000)

	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}

func BenchmarkSort(b *testing.B) {
	elements := GetElements(100000)

	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}
