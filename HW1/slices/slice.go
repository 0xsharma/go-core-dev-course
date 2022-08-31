package main

import (
	"fmt"
	"sort"
)

func main() {

	// Init Slice with ints
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("slice1: ", slice1)

	//Add given int(N) to each of []int elements
	slice1 = addToEach(slice1, 10)
	fmt.Println("slice1 after adding 10: ", slice1)

	// Add the number int(N) to the end of the slice.
	slice1 = addToEnd(slice1, 21)
	fmt.Println("slice1 after adding 21 to end: ", slice1)

	// Add the number int(N) to the beginning of the slice.
	slice1 = addToBeginning(slice1, 0)
	fmt.Println("slice1 after adding 0 to beginning: ", slice1)

	// Take the last number of slice, return it to the user, and remove this element from slice.
	slice1, last := takeLast(slice1)
	fmt.Println("slice1 last : ", last)
	fmt.Println("slice1 after taking last: ", slice1)

	// Take the first slice number, return it to the user, and delete this element from slice.
	slice1, first := takeFirst(slice1)
	fmt.Println("slice1 first : ", first)
	fmt.Println("slice1 after taking first: ", slice1)

	// Take the i-th number of slice, return it to the user, and delete this element from slice. The number i is passed by the user to the function. Try to do it without allocating a new slice.
	slice1, ith := takeIth(slice1, 2)
	fmt.Println("slice1 ith : ", ith)
	fmt.Println("slice1 after taking ith: ", slice1)

	// Merge two slices and return a new one with all elements of the first and second without duplicates.
	slice2 := []int{11, 12, 13, 4, 5, 6, 7, 8, 9, 10}
	slice3 := merge(slice1, slice2)
	fmt.Println("slice1 : ", slice1)
	fmt.Println("slice2 : ", slice2)
	fmt.Println("slice3 after merging slice1 and slice2: ", slice3)

	// From the first slice, remove all the numbers that are in the second.
	fmt.Println("slice1 : ", slice1)
	fmt.Println("slice2 : ", slice2)
	slice3 = removeCommon(slice1, slice2)
	fmt.Println("slice1 after removing common elements of slice2 : ", slice3)

	// Shift all slice elements by 1 to the left. Zero index element becomes the last one, the first - zero, the last - penultimate.
	slice1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("slice1 : ", slice1)
	slice1 = shiftLeft(slice1)
	fmt.Println("slice1 after shiftLeft: ", slice1)

	// Shift all slice elements by N int to the left.
	slice1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("slice1 : ", slice1)
	slice1 = shiftLeftIter(slice1, 3)
	fmt.Println("slice1 after shiftLeft: ", slice1)

	// Shift all slice elements by 1 to the right.
	slice1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("slice1 : ", slice1)
	slice1 = shiftRight(slice1)
	fmt.Println("slice1 after shiftRight: ", slice1)

	// Shift all slice elements by N int to the right.
	slice1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("slice1 : ", slice1)
	slice1 = shiftRightIter(slice1, 3)
	fmt.Println("slice1 after shiftRight: ", slice1)

	// Return a copy of the passed slice to the user
	slice1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("slice1 : ", slice1)
	slice2 = returnCopy(slice1)
	fmt.Println("slice2 ( deep copy of slice1 ) : ", slice2)

	// swap all even index elements with the nearest odd indices.
	slice1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("slice1 : ", slice1)
	slice1 = swapEvenOdd(slice1)
	fmt.Println("slice1 after swapEvenOdd: ", slice1)

	// sort ascending slice ints
	slice1 = []int{1, 8, 4, 5, 9, 6, 2, 3, 7, 10}
	fmt.Println("slice1 : ", slice1)
	sortInts(slice1, true)
	fmt.Println("slice1 after sortAsc: ", slice1)
	sortInts(slice1, false)
	fmt.Println("slice1 after sortDesc: ", slice1)

	// sort slice strings
	sliceStr := []string{"a", "e", "f", "i", "d", "b", "c", "h", "g", "j"}
	fmt.Println("sliceStr : ", sliceStr)
	sortStrings(sliceStr, true)
	fmt.Println("sliceStr after sortAsc: ", sliceStr)
	sortStrings(sliceStr, false)
	fmt.Println("sliceStr after sortDesc: ", sliceStr)

}

// Add given int(N) to each of []int elements
func addToEach(slice []int, N int) []int {
	for i := 0; i < len(slice); i++ {
		slice[i] = slice[i] + N
	}
	return slice
}

// Add the number int(N) to the end of the slice.
func addToEnd(slice []int, N int) []int {
	slice = append(slice, N)
	return slice
}

// Add the number int(N) to the beginning of the slice.
func addToBeginning(slice []int, N int) []int {
	slice = append([]int{N}, slice...)
	return slice
}

// Take the last number of slice, return it to the user, and remove this element from slice.
func takeLast(slice []int) ([]int, int) {
	last := slice[len(slice)-1]
	slice = slice[:len(slice)-1]
	return slice, last
}

// Take the first slice number, return it to the user, and delete this element from slice.
func takeFirst(slice []int) ([]int, int) {
	first := slice[0]
	slice = slice[1:]
	return slice, first
}

// Take the i-th number of slice, return it to the user, and delete this element from slice. The number i is passed by the user to the function. Try to do it without allocating a new slice.
func takeIth(slice []int, i int) ([]int, int) {
	ith := slice[i]
	slice = append(slice[:i], slice[i+1:]...)
	return slice, ith
}

// Merge two slices and return a new one with all elements of the first and second without duplicates.
func merge(slice1 []int, slice2 []int) []int {
	slice := make([]int, 0, len(slice1)+len(slice2))

	m := make(map[int]bool)

	for _, v := range slice1 {
		if !m[v] {
			m[v] = true
			slice = append(slice, v)
		}
	}
	for _, v := range slice2 {
		if !m[v] {
			m[v] = true
			slice = append(slice, v)
		}
	}

	return slice
}

// From the first slice, remove all the numbers that are in the second.
func removeCommon(slice1 []int, slice2 []int) []int {
	for _, v := range slice2 {
		slice1 = removeElement(slice1, v)
	}
	return slice1
}

// Remove Element from slice
func removeElement(slice []int, element int) []int {
	for i, v := range slice {
		if v == element {
			slice = append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

// Shift all slice elements by 1 to the left. Zero index element becomes the last one, the first - zero, the last - penultimate.Shift all slice elements by 1 to the left. Zero index element becomes the last one, the first - zero, the last - penultimate.
func shiftLeft(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}
	slice = append(slice[1:], slice[:1]...)

	return slice
}

// Shift all slice elements by N int to the left.
func shiftLeftIter(slice []int, i int) []int {
	if len(slice) < 2 || i < 1 || i >= len(slice) {
		return slice
	}

	slice = append(slice[i:], slice[:i]...)

	return slice
}

// Shift all slice elements by 1 to the right.
func shiftRight(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}
	slice = append(slice[len(slice)-1:], slice[:len(slice)-1]...)

	return slice
}

// Shift all slice elements by N int to the right.
func shiftRightIter(slice []int, i int) []int {
	if len(slice) < 2 || i < 1 || i >= len(slice) {
		return slice
	}

	slice = append(slice[len(slice)-i:], slice[:len(slice)-i]...)

	return slice
}

// Return a copy of the passed slice to the user
func returnCopy(slice []int) []int {
	cpy := make([]int, len(slice))
	copy(cpy, slice)

	return cpy
}

// swap all even index elements with the nearest odd indices.
func swapEvenOdd(slice []int) []int {
	for i := 0; i < len(slice); i += 2 {
		slice[i], slice[i+1] = slice[i+1], slice[i]
	}
	return slice
}

// sort slice ints. true for ascending order, false for descending order
func sortInts(slice []int, asc bool) {
	sort.Slice(slice, func(i, j int) bool {
		if asc {
			return slice[i] < slice[j]
		} else {
			return slice[i] > slice[j]
		}
	})
}

// sort slice strings. true for ascending order, false for descending order
func sortStrings(slice []string, asc bool) {
	sort.Slice(slice, func(i, j int) bool {
		if asc {
			return slice[i] < slice[j]
		} else {
			return slice[i] > slice[j]
		}
	})
}
