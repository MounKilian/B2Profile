package algo

func InsertionSort(array []int) []int {
	for i := 1; i < len(array); i++ {
		j := i
		for j > 0 {
			if array[j-1] > array[j] {
				array[j-1], array[j] = array[j], array[j-1]
			}
			j--
		}
	}
	return array
}
