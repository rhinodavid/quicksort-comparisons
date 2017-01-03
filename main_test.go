package main

import "testing"

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
	nums := []int{3, 5, 6, 7, 1, 4, 2}
	sorted, _ := partition(nums, 0, len(nums))
	if !testEq(sorted, []int{1, 2, 3, 4, 5, 6, 7}) {
		t.Fatalf("Incorrect sorted array")
	}
}

// counting comparisons
func TestCountComparisons(t *testing.T) {
	nums := []int{3, 2, 1}
	_, comparisons := partition(nums, 0, len(nums))
	if comparisons != 2 {
		t.Fatalf("Incorrect number of comparisons, expected 2, got %d\n", comparisons)
	}
}
