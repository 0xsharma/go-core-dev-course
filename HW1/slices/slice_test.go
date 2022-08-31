package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddToEach(t *testing.T) {
	t.Parallel()

	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	slice1 = addToEach(slice1, 10)
	assert.Equal(t, expected, slice1)
}

func TestAddToEnd(t *testing.T) {
	t.Parallel()

	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 21}
	slice1 = addToEnd(slice1, 21)
	assert.Equal(t, expected, slice1)
}

func TestAddToBeginning(t *testing.T) {
	t.Parallel()

	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice1 = addToBeginning(slice1, 0)
	assert.Equal(t, expected, slice1)
}

func TestTakeLast(t *testing.T) {
	t.Parallel()

	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expectedLast := 10
	slice1, last := takeLast(slice1)
	assert.Equal(t, expected, slice1)
	assert.Equal(t, expectedLast, last)
}

func TestTakeFirst(t *testing.T) {
	t.Parallel()

	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}
	expectedFirst := 1
	slice1, first := takeFirst(slice1)
	assert.Equal(t, expected, slice1)
	assert.Equal(t, expectedFirst, first)
}

func TestTakeIth(t *testing.T) {
	t.Parallel()

	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{1, 2, 4, 5, 6, 7, 8, 9, 10}
	expectedIth := 3
	slice1, ith := takeIth(slice1, 2)
	assert.Equal(t, expected, slice1)
	assert.Equal(t, expectedIth, ith)
}

func TestMerge(t *testing.T) {
	t.Parallel()

	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice2 := []int{1, 2, 3, 4, 15, 16, 17, 18, 19, 20}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 15, 16, 17, 18, 19, 20}
	slice1 = merge(slice1, slice2)
	assert.Equal(t, expected, slice1)
}

func TestRemoveCommon(t *testing.T) {
	t.Parallel()

	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice2 := []int{1, 2, 3, 4, 15, 16, 17, 18, 19, 20}
	expected := []int{5, 6, 7, 8, 9, 10}
	slice1 = removeCommon(slice1, slice2)
	assert.Equal(t, expected, slice1)
}

func TestRemoveElement(t *testing.T) {
	t.Parallel()

	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{1, 2, 3, 4, 5, 7, 8, 9, 10}
	slice1 = removeElement(slice1, 6)
	assert.Equal(t, expected, slice1)
}

func TestShiftLeft(t *testing.T) {
	t.Parallel()

	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 1}
	slice1 = shiftLeft(slice1)
	assert.Equal(t, expected, slice1)
}

func TestShiftRight(t *testing.T) {
	t.Parallel()

	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{10, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice1 = shiftRight(slice1)
	assert.Equal(t, expected, slice1)
}

func TestShiftLeftIter(t *testing.T) {
	t.Parallel()

	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{4, 5, 6, 7, 8, 9, 10, 1, 2, 3}
	slice1 = shiftLeftIter(slice1, 3)
	assert.Equal(t, expected, slice1)
}

func TestShiftRightIter(t *testing.T) {
	t.Parallel()

	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{8, 9, 10, 1, 2, 3, 4, 5, 6, 7}
	slice1 = shiftRightIter(slice1, 3)
	assert.Equal(t, expected, slice1)
}

func TestSwapEvenOdd(t *testing.T) {
	t.Parallel()

	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{2, 1, 4, 3, 6, 5, 8, 7, 10, 9}
	slice1 = swapEvenOdd(slice1)
	assert.Equal(t, expected, slice1)
}

func TestSortInts(t *testing.T) {
	t.Parallel()

	slice1 := []int{4, 2, 3, 8, 5, 9, 10, 6, 7, 1}
	expectedDesc := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	expectedAsc := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// descending
	sortInts(slice1, false)
	assert.Equal(t, expectedDesc, slice1)

	//ascending
	sortInts(slice1, true)
	assert.Equal(t, expectedAsc, slice1)
}

func TestSortStrings(t *testing.T) {
	t.Parallel()

	slice1 := []string{"d", "f", "g", "e", "h", "a", "c", "b", "i", "j"}
	expectedDesc := []string{"j", "i", "h", "g", "f", "e", "d", "c", "b", "a"}
	expectedAsc := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

	// descending
	sortStrings(slice1, false)
	assert.Equal(t, expectedDesc, slice1)

	//ascending
	sortStrings(slice1, true)
	assert.Equal(t, expectedAsc, slice1)
}
