package sort

import "fmt"

func Insertion(arr []int) {
	fmt.Println("Insertion sort")

	if len(arr) < 2 {
		return
	}

	for i := 1; i < len(arr); i++ {
		numberToInsert := arr[i]
		j := i - 1
		for ; j >= 0 && arr[j] > numberToInsert; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = numberToInsert
	}
}
