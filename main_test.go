package main

import (
	"fmt"
	"testing"
)

func testEq(a, b []int) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

// partition
func XTestPivot(t *testing.T) {
	nums := []int{3, 5, 6, 7, 1, 4, 2}
	sorted, _ := partition(nums, 0, len(nums))
	fmt.Println(sorted)
	if !testEq(sorted, []int{1, 2, 3, 4, 5, 6, 7}) {
		t.Fatalf("Incorrect sorted array")
	}
}

func TestCountPartitions(t *testing.T) {
	nums := []int{3, 2, 1}
	_, comparisons := partition(nums, 0, len(nums))
	fmt.Println(comparisons)
	if comparisons != 3 {
		t.Fatalf("Incorrect number of comparisons, expected 3, got %d\n", comparisons)
	}
}
