package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

func main() {
	// There is a text, you need to count how many times each word occurs.
	const str = "shivam sharma core dev course maps attempt shivam course"
	wordCountMap := printWordCount(str)

	for key, value := range wordCountMap {
		fmt.Printf("%s: %d\n", key, value)
	}

	fmt.Println()
	// There is a very large array or slice of integers, it must be said which numbers are mentioned in it at least once.
	slice0 := []int{1, 2, 4, 6, 8, 9, 20, 1, 6}
	mapPresent := printPresentNumbers(slice0)
	fmt.Printf("Present numbers in slice: %v\n", mapPresent)

	fmt.Printf("\n\n")

	// There are two large arrays of numbers, you need to find which numbers are mentioned in both.
	slice1 := []int{1, 2, 3, 4, 7, 9, 19, 2, 6}

	slice2 := commonNumbers(slice0, slice1)
	fmt.Printf("slice1: %v\n", slice1)
	fmt.Printf("slice2: %v\n", slice2)
	fmt.Printf("Common numbers are : %v\n\n", slice2)

	// find nth fibonacci number using memoization in maps
	val1 := initFibonacci(5)
	fmt.Printf("fibonacci(5): %d\n", val1)

	fmt.Println("")

	// create a random multi dimensional map and print the sorted version of it
	// map[fee]map[nonce]exists
	m2 := generateRandom3dMap(3, 3)

	// you have map[fee]map[nonce]Account (Account [20]byte), you need to print it as sorted data: descending sorting by fee first and ascending by nonce last.
	sort3dMap(m2)
	fmt.Println(m2)
}

// generateRandom3dMap generates a 3d map with random values
func generateRandom3dMap(l1, l2 int) map[int]map[int]bool {
	m := make(map[int]map[int]bool)

	for i := 0; i < l1; i++ {
		// nolint:gosec
		idx := rand.Intn(1000)

		for {
			if _, ok := m[idx]; ok {
				// nolint:gosec
				idx = rand.Intn(1000)
			} else {
				break
			}
		}

		m[idx] = make(map[int]bool)

		for j := 0; j < l2; j++ {
			// nolint:gosec
			m[idx][rand.Intn(1000)] = true
		}
	}

	return m
}

// There is a text, you need to count how many times each word occurs.
func printWordCount(str string) map[string]int {
	strArr := strings.Fields(str)
	mapCount := make(map[string]int, 0)

	for _, val := range strArr {
		mapCount[val]++
	}

	fmt.Println("string ", str)

	return mapCount
}

// There is a very large array or slice of integers, it must be said which numbers are mentioned in it at least once.
func printPresentNumbers(ints []int) []int {
	mapPresent := make(map[int]bool, len(ints))

	for i := 0; i < len(ints); i++ {
		mapPresent[ints[i]] = true
	}

	fmt.Printf("Occurrence of each number in slice: %v\n", ints)

	ansArr := make([]int, 0, len(mapPresent))
	for key := range mapPresent {
		ansArr = append(ansArr, key)
	}

	return ansArr
}

// There are two large arrays of numbers, you need to find which numbers are mentioned in both.
func commonNumbers(ints1 []int, ints2 []int) []int {
	commonMap := make(map[int]bool, len(ints1))

	for i := 0; i < len(ints1); i++ {
		commonMap[ints1[i]] = true
	}

	min := len(ints1)
	if len(ints2) < len(ints1) {
		min = len(ints2)
	}

	ansArr := make([]int, 0, min)
	for i := 0; i < len(ints2); i++ {
		if len(ansArr) == min {
			break
		}

		if _, ok := commonMap[ints2[i]]; ok {
			ansArr = append(ansArr, ints2[i])
		}
	}

	return ansArr
}

// fibonacci uses map for memoization of values for seen indexes.
// It also returns the map for future use.
func initFibonacci(n int) int {
	fibMap := make(map[int]int, 0)
	fibMap[0] = 1
	fibMap[1] = 1

	return fibonnaci(n, fibMap)
}

func fibonnaci(n int, fibMap map[int]int) int {
	if val, ok := fibMap[n]; ok {
		return val
	}

	fibMap[n] = fibonnaci(n-1, fibMap) + fibonnaci(n-2, fibMap)

	return fibMap[n]
}

// you have map[fee]map[nonce]Account (Account [20]byte), you need to print it as sorted data: descending sorting by fee first and ascending by nonce last.
func sort3dMap(_3dmap map[int]map[int]bool) {
	keys1 := make([]int, 0)
	for k1 := range _3dmap {
		keys1 = append(keys1, k1)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(keys1)))

	for _, k1 := range keys1 {
		keys2 := make([]int, 0)
		for k2 := range _3dmap[k1] {
			keys2 = append(keys2, k2)
		}

		sort.Sort(sort.IntSlice(keys2))

		for _, k2 := range keys2 {
			fmt.Printf("Fee : %d ; Nonce : %d \n", k1, k2)
		}

		fmt.Println()
	}
}
