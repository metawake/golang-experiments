package main

import (
	"fmt"
)

func getQsLTE(l []int, pivot int) []int {
	var result []int
	for _, elem := range l {
		if elem <= pivot {
			result = append(result, elem)
		}
	}
	return qs(result)
}

func getQsGT(l []int, pivot int) []int {
	var result []int
	for _, elem := range l {
		if elem > pivot {
			result = append(result, elem)
		}
	}
	return qs(result)
}

func qs(l []int) []int {
	if len(l) < 2 {
		return l
	}

	pivot := l[len(l)-1]
	l = l[:len(l)-1]

	return append(append(getQsLTE(l, pivot), pivot), getQsGT(l, pivot)...)
}

func test1() {
	var l []int
	l = qs(l)
	assertEqual(l, []int{})

	l = []int{10}
	l = qs(l)
	assertEqual(l, []int{10})
}

func test2() {
	l := []int{2, 1, 3}
	initialL := make([]int, len(l))
	copy(initialL, l)
	l = qs(l)
	assertEqual(l, quickSort(initialL))
}

func test3() {
	l := []int{9, 6, 1, 8, 5, 4, 2, 3, 7}
	initialL := make([]int, len(l))
	copy(initialL, l)
	l = qs(l)
	assertEqual(l, quickSort(initialL))
}

func runAllTests() {
	test1()
	test2()
	test3()

	fmt.Println("Completed: runAllTests()")
}

func assertEqual(a, b []int) {
	if len(a) != len(b) {
		panic("Lengths of slices are different")
	}
	for i := range a {
		if a[i] != b[i] {
			panic(fmt.Sprintf("Slices differ at index %d: %d != %d", i, a[i], b[i]))
		}
	}
}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	var left, right []int
	pivot := arr[0]

	for _, v := range arr[1:] {
		if v <= pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	return append(append(quickSort(left), pivot), quickSort(right)...)
}

func main() {
	runAllTests()
}
