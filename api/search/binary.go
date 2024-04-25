package search

func Binary(numbers []int, element int) int {
	var lo, hi = 0, len(numbers) - 1
	pivotIndex := hi / 2
	for lo <= hi {
		pivot := numbers[pivotIndex]
		if pivot < element {
			lo = pivotIndex + 1
		} else if pivot > element {
			hi = pivotIndex - 1
		} else {
			return pivot
		}
		pivotIndex = lo + (hi-lo+1)/2
	}
	return -1
}
