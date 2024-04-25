package sort

func merge(arr1 []int, arr2 []int) []int {
	var merged []int
	for sentinel1, sentinel2 := 0, 0; sentinel1 < len(arr1) || sentinel2 < len(arr2); {
		if sentinel1 > len(arr1)-1 {
			// Arr1 is empty
			merged = append(merged, arr2[sentinel2:]...)
			sentinel2++
			break
		}
		if sentinel2 > len(arr2)-1 {
			// Arr2 is empty
			merged = append(merged, arr1[sentinel1:]...)
			sentinel1++
			break
		}
		if arr1[sentinel1] > arr2[sentinel2] {
			merged = append(merged, arr2[sentinel2])
			sentinel2++
		} else {
			merged = append(merged, arr1[sentinel1])
			sentinel1++
		}
	}
	return merged
}

func MergeSort(arr []int) []int {
	if len(arr) > 1 {
		divider := len(arr) / 2
		// fmt.Printf("arr1: %v\n", arr[:divider])
		// fmt.Printf("arr2: %v\n", arr[divider:])
		sorted1 := MergeSort(arr[:divider])
		sorted2 := MergeSort(arr[divider:])
		return merge(sorted1, sorted2)
	}
	return arr
}
