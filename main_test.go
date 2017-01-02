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
func TestPivot(t *testing.T) {
	nums := []int{3, 5, 7, 1, 4, 2}
	sorted := partition(nums, 0, len(nums))
	fmt.Println(sorted)
	if !testEq(sorted, []int{1, 2, 3, 4, 5}) {
		t.Fatalf("Incorrect sorted array")
	}
}
