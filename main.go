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

func selectFirstAsPivot(start int, end int) int {
	return start
}

func partition(arr []int, start int, end int) ([]int, int) {
	comparisons := (end - start) - 1
	fmt.Printf("starting with start %d, end %d\n", start, end)
	if end == start {
		return arr, comparisons
	}
	// select pivot
	pivot := selectFirstAsPivot(start, end)

	pivotValue := arr[pivot]
	fmt.Println(pivotValue)
	// swap pivot and first
	tmp := arr[start]
	arr[start] = arr[pivot]
	arr[pivot] = tmp

	fmt.Println(arr)
	i := start + 1
	for j := start + 1; j < end; j++ {
		fmt.Println(i)
		if arr[j] < pivotValue {
			tmp := arr[i]
			arr[i] = arr[j]
			arr[j] = tmp
			i++
		}
		fmt.Println(arr)
	}
	tmp = arr[i-1]
	arr[i-1] = arr[pivot]
	arr[pivot] = tmp
	fmt.Println(arr)
	fmt.Println("i", i)
	if (i-1)-start > 1 {
		_, comp := partition(arr, start, (i - 1))
		comparisons += comp
	}
	if end-i > 1 {
		_, comp := partition(arr, i, end)
		comparisons += comp
	}

	return arr, comparisons
}

func countQuicksortComparisons(arr []int) (comparisons int) {
	comparisons = 0
	return comparisons
}

func main() {
	flag.Parse()
	if len(flag.Args()) < 1 {
		panic("Enter the name of the file with the integer list you'd like to sort and count comparisons.")
	}
	_, err := readFile(flag.Args()[0])
	if err != nil {
		panic(err)
	}
	comparisons := 0
	fmt.Println(comparisons)
}
