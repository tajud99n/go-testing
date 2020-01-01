package services

import (
	"testing"
)

func TestSort(t *testing.T) {
	elements := []int{7, 6, 4, 0, 2, 5}

	Sort(elements)

	if elements[0] != 0 {
		t.Error("First element should be 0")
	}

	if elements[len(elements)-1] != 7 {
		t.Error("Last element should be 7")
	}
}
