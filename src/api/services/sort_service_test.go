package services

import (
	"testing"
	"github.com/tajud99n/go-testing/src/api/utils/sort"
)

func TestSortLessThan10000(t *testing.T) {
	elements := sort.GetElements(1000)

	Sort(elements)

	if elements[0] != 0 {
		t.Error("First element should be 0")
	}

	if elements[len(elements)-1] != 999 {
		t.Error("Last element should be 999")
	}
}

func TestSortMoreThan10000(t *testing.T) {
	elements := sort.GetElements(100000)

	Sort(elements)

	if elements[0] != 0 {
		t.Error("First element should be 0")
	}

	if elements[len(elements)-1] != 99999 {
		t.Error("Last element should be 99999")
	}
}

func BenchmarkSortMoreThan10000(b *testing.B) {
	elements := sort.GetElements(100000)

	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}

func BenchmarkSortLessThan10000(b *testing.B) {
	elements := sort.GetElements(5000)

	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}