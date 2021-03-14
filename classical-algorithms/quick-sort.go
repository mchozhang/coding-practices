package main

import "fmt"

func quickSort(arr []int, start int, end int) {
	if start >= end {
		return
	}

	pivot := arr[start]

	low := start + 1
	high := end

	// correct position of the pivot in the end
	pivotEnds := 0
	for low <= high {
		if arr[low] > pivot {
			arr[low], arr[high] = arr[high], arr[low]
			high--
			pivotEnds = high
		} else {
			low++
			pivotEnds = high - 1
		}
	}

	// swap the pivot to correct position
	arr[start], arr[pivotEnds] = arr[pivotEnds], arr[start]

	quickSort(arr, start, high-1)
	quickSort(arr, high+1, end)
}

func main() {
	// arr := []int{5, 1, 2, 4, 9, 3, 7, 8, 0}
	// arr := []int{3,5,3,5}
	arr := []int{2,2,1,1,1,1,2}

	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
