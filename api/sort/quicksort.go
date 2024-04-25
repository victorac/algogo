package sort

func Quicksort(arr []int, startIndex int, endIndex int) {
	if startIndex >= endIndex || startIndex < 0 {
		return
	}

	pivot := arr[endIndex]
	nextPivotIndex := startIndex

	for j := startIndex; j < endIndex; j++ {
		if arr[j] <= pivot {
			if j != nextPivotIndex {
				// arr[nextPivotIndex] has the first value starting from startIndex that is greater than the pivot;
				// swap it arr[nextPivotIndex] with arr[j] because arr[j] is less or equal to the pivot;
				temp := arr[j]
				arr[j] = arr[nextPivotIndex]
				arr[nextPivotIndex] = temp
			}
			nextPivotIndex++
		}
	}

	// arr[nextPivotIndex] has the first value of the partition that is greater than the pivot
	// swap it with the pivot
	arr[endIndex] = arr[nextPivotIndex]
	arr[nextPivotIndex] = pivot

	Quicksort(arr, startIndex, nextPivotIndex-1)
	Quicksort(arr, nextPivotIndex+1, endIndex)
}
