package sort

func Heap(array []int) {
	if len(array) <= 1 {
		return
	}
	// create a max heap
	lastParentIndex := len(array)/2 - 1
	for i := lastParentIndex; i >= 0; i-- {
		heapify(array, len(array), i)
	}

	for i := 1; i < len(array); i++ {
		swap(array, 0, len(array)-i)
		heapify(array, len(array)-i, 0)
	}
}

func swap(array []int, i int, j int) {
	tmp := array[i]
	array[i] = array[j]
	array[j] = tmp
}

func heapify(array []int, lenght int, index int) {
	lastParentIndex := lenght/2 - 1
	if index > lastParentIndex {
		return
	}
	var greaterChildIndex int

	if index*2+1 < lenght {
		greaterChildIndex = index*2 + 1
		if index*2+2 < lenght {
			if array[greaterChildIndex] < array[index*2+2] {
				greaterChildIndex = index*2 + 2
			}
		}
	}

	if greaterChildIndex > 0 && array[index] < array[greaterChildIndex] {
		swap(array, index, greaterChildIndex)
		heapify(array, lenght, greaterChildIndex)
	}

}
