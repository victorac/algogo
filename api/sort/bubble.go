package sort

import "fmt"

func Bubble(arr []int) {
	fmt.Println("Bubble Sort")
	// var arr = []int{1, 2, 3}
	// var arr = []int{3, 1, 2}
	var length = len(arr)
	var swap bool = true
	for swap {
		swap = false
		for i := 0; i < length-1; i++ {
			if arr[i] > arr[i+1] {
				aux := arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = aux
				swap = true
			}
		}
		length--
		fmt.Printf("Arr: %v, Swap: %v, new length: %v\n", arr, swap, length)
	}
}
