package main

import "fmt"

func main() {
	fmt.Println(insertionSort([...]int{3, 10, 10, 9, 2}, true))
	fmt.Println(linearSearch([...]int{3, 10, 10, 9, 2}, 0))
}

func insertionSort(A [5]int, isNonincresing bool) [5]int {
	if isNonincresing {
		return nonincresingInsertionSort(A)
	}

	return incresingInsertionSort(A)
}

func incresingInsertionSort(A [5]int) [5]int {
	for j := 1; j < len(A); j++ {
		var key = A[j]
		var i = j - 1;
		for i > -1 && A[i] > key {
			A[i + 1] = A[i]
			i--
		}
		A[i + 1] = key
	}

	return A
}


func nonincresingInsertionSort(A [5]int) [5]int {
	for j := 1; j < len(A); j++ {
		var key = A[j]
		var i = j - 1

		for i > -1 && A[i] < key {
			A[i + 1] = A[i]
			i--
		}
		A[i + 1] = key
	}

	return A
}

func linearSearch(A [5]int, v int) int {
	for j := 0; j < len(A); j++ {
		if A[j] == v {
			return j
		}
	}

	return -1
}