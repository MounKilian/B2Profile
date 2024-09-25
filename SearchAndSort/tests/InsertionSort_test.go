package algo_test

import (
	"algo"
	"math/rand"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}
	testArray := algo.InsertionSort(array)
	if !IsSorted(testArray) {
		t.Errorf("Sorted array InsertionSort test fail, expected %+v, got %+v", array, testArray)
	}

	array = []int{5, 4, 3, 2, 1}
	testArray = algo.InsertionSort(array)
	if !IsSorted(testArray) {
		t.Errorf("Reversed array InsertionSort test fail, expected {1, 2, 3, 4, 5}, got %+v", testArray)
	}
}

func TestInsertionSortLarge(t *testing.T) {
    array := rand.Perm(10000)
    testArray := algo.InsertionSort(array)
    if !IsSorted(testArray) {
        t.Errorf("Large random array InsertionSort test fail")
    }
}

func IsSorted(array []int) bool {
	if len(array) <= 1 {
		return true
	}
	for i := 1; i < len(array); i++ {
		if array[i-1] > array[i] {
			return false
		}
	}
	return true
}
