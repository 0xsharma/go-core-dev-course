package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintWordCount(t *testing.T) {
	t.Parallel()

	str := "shivam sharma core dev course maps attempt shivam course"
	expected := map[string]int{"shivam": 2, "sharma": 1, "core": 1, "dev": 1, "course": 2, "maps": 1, "attempt": 1}
	assert.Equal(t, expected, printWordCount(str))
}

func TestPrintPresentNumbers(t *testing.T) {
	t.Parallel()

	slice1 := []int{1, 2, 3, 4, 5, 6, 1, 2, 3, 4}
	expected := []int{1, 2, 3, 4, 5, 6}
	ans := printPresentNumbers(slice1)
	sort.Ints(ans)
	assert.Equal(t, expected, ans)
}

func TestCommonNumbers(t *testing.T) {
	t.Parallel()

	slice1 := []int{1, 2, 3, 4, 5, 6}
	slice2 := []int{1, 2, 3, 6, 7, 8}
	expected := []int{1, 2, 3, 6}
	ans := commonNumbers(slice1, slice2)
	sort.Ints(ans)
	assert.Equal(t, expected, ans)
}

func TestFibbonaci(t *testing.T) {
	t.Parallel()

	n := 10
	expected := []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}

	for i := 0; i < n; i++ {
		assert.Equal(t, expected[i], initFibonacci(i))
	}
}
