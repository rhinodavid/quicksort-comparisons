package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// See: https://stackoverflow.com/questions/9862443/golang-is-there-a-better-way-read-a-file-of-integers-into-an-array
func readFile(fname string) (nums []int, err error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")
	// Assign cap to avoid resize on every append.
	nums = make([]int, 0, len(lines))

	for _, l := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(l) == 0 {
			continue
		}
		// Atoi better suits the job when we know exactly what we're dealing
		// with. Scanf is the more general option.
		n, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}

	return nums, nil
}

func selectFirstAsPivot(arr []int, start int, end int) int {
	return start
}

func selectLastAsPivot(arr []int, start int, end int) int {
	return end - 1
}

func selectMedianOfThreeAsPivot(arr []int, start int, end int) int {
	first := arr[start]
	last := arr[end-1]
	middle := arr[(end-1-start)/2+start]
	if (first < middle && first > last) || (first > middle && first < last) {
		return start
	}
	if (middle > first && middle < last) || (middle < first && middle > last) {
		return (end-1-start)/2 + start
	}
	return end - 1
}

func partition(arr []int, start int, end int) ([]int, int) {
	comparisons := (end - start) - 1
	if end == start {
		return arr, comparisons
	}
	// select pivot
	pivot := selectMedianOfThreeAsPivot(arr, start, end)
	pivotValue := arr[pivot]

	// swap pivot and first
	tmp := arr[start]
	arr[start] = arr[pivot]
	arr[pivot] = tmp

	i := start + 1
	for j := start + 1; j < end; j++ {
		if arr[j] < pivotValue {
			tmp := arr[i]
			arr[i] = arr[j]
			arr[j] = tmp
			i++
		}
	}

	// put pivot in its correct spot
	tmp = arr[i-1]
	arr[i-1] = arr[start]
	arr[start] = tmp

	// recurse on left half if necessary
	if (i-1)-start > 1 {
		_, comp := partition(arr, start, (i - 1))
		comparisons += comp
	}
	// recurse on right half if necessary
	if end-i > 1 {
		_, comp := partition(arr, i, end)
		comparisons += comp
	}

	return arr, comparisons
}

func main() {
	flag.Parse()
	if len(flag.Args()) < 1 {
		panic("Enter the name of the file with the integer list you'd like to sort and count comparisons.")
	}
	arr, err := readFile(flag.Args()[0])
	if err != nil {
		panic(err)
	}
	_, comparisons := partition(arr, 0, len(arr))

	fmt.Println(comparisons)
}
