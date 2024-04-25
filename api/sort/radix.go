package sort

import (
	"math"
)

func getDigit(number int, place int, longestNumber int) int {
	var digit int
	for i := longestNumber - 1; i >= place; i-- {
		power := int(math.Pow10(i))
		digit = number / power
		number -= digit * power
	}
	return digit
}

func getLongest(arr []int) int {
	longest := 1
	for i := 0; i < len(arr); i++ {
		log := int(math.Log10(float64(arr[i]))) + 1
		if log > longest {
			longest = log
		}
	}
	return longest
}

type bucketItem struct {
	number int
	digit  int
}

func Radix(numbers []int) {
	longest := getLongest(numbers)
	for i := 0; i < longest; i++ {
		buckets := [10][]bucketItem{}
		for j := 0; j < len(numbers); j++ {
			digit := getDigit(numbers[j], i, longest)
			buckets[digit] = append(buckets[digit], bucketItem{numbers[j], digit})
		}
		// fmt.Printf("Buckets: %v\n", buckets)
		index := 0
		for j := 0; j < 10; j++ {
			// for each bucket
			for l := 0; l < len(buckets[j]); l++ {
				numbers[index] = buckets[j][l].number
				index++
			}
		}
	}
}
