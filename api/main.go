package main

import (
	"algo/lists"
	"algo/recursion"
	"algo/routes"
	"algo/search"
	"algo/sort"
	"algo/trees"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
)

func runSorting() {
	var arr1 = []int{64, 34, 25, 12, 22, 11, 90}
	var arr2 = []int{1, 2, 3}
	var arr3 = []int{3, 2, 1}
	var arr4 = []int{3, 1, 2}
	var arrays = [][]int{
		arr1,
		arr2,
		arr3,
		arr4,
	}
	for i := 0; i < len(arrays); i++ {
		fmt.Printf("Start array: %v\n", arrays[i])
		// Bubble sort
		// sort.Bubble(arr)
		// Insertion sort
		// sort.Insertion(arrays[i])
		// // Merge sort
		// sorted := sort.MergeSort(arrays[i])
		// fmt.Printf("%v\n", sorted)
		// Quicksort
		// sort.Quicksort(arrays[i], 0, len(arrays[i])-1)
		sort.Radix(arrays[i])
		fmt.Printf("End array: %v\n", arrays[i])
	}
}

var nestedArr1 = []interface{}{
	1,
	2,
	[]interface{}{
		3,
		4,
	},
	5,
	[]interface{}{
		6,
		7,
		[]interface{}{
			8,
		},
	},
}
var nestedArr2 = []interface{}{
	1,
	2,
	3,
	4,
}
var nestedArr3 = []interface{}{}
var nestedArr4 = []interface{}{
	1,
}
var nestedArr5 = []interface{}{
	[]interface{}{
		[]interface{}{
			[]interface{}{
				[]interface{}{
					5,
				},
			},
		},
	},
}
var nestedArrays = [][]interface{}{
	nestedArr1,
	nestedArr2,
	nestedArr3,
	nestedArr4,
	nestedArr5,
}

func runRecursive() {
	fmt.Println("Nested Arrays")
	for i := 0; i < len(nestedArrays); i++ {
		fmt.Printf("%v\n", nestedArrays[i])
		sum := recursion.NestedAdd(nestedArrays[i], 0)
		fmt.Printf("\tSum: %d\n", sum)
	}
	fmt.Println("Factorial")
	fmt.Printf("5!: %v\n", recursion.Factorial(5))
	fmt.Printf("4!: %v\n", recursion.Factorial(4))
	fmt.Printf("3!: %v\n", recursion.Factorial(3))
	fmt.Printf("2!: %v\n", recursion.Factorial(2))
	fmt.Printf("1!: %v\n", recursion.Factorial(1))
	fmt.Printf("0!: %v\n", recursion.Factorial(0))

}

type SearchParam struct {
	numbers []int
	element int
}

func runSearch() {
	arrays := []SearchParam{
		SearchParam{[]int{1, 2, 3, 4, 5}, 2},
		SearchParam{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 5},
		SearchParam{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 8},
		SearchParam{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 10},
		SearchParam{[]int{10, 20, 30, 40, 50, 60, 70, 80}, 15},
	}
	for i := 0; i < len(arrays); i++ {
		fmt.Printf("Searching %v, for element: %v\n", arrays[i].numbers, arrays[i].element)
		fmt.Printf("Position: %v\n", search.Binary(arrays[i].numbers, arrays[i].element))
	}
}

func runArrayList() {
	arr := lists.Array{}

	arr.Push(1)
	arr.Push(2)
	arr.Push(3)

	arr.Pop()

	arr.Push(35)

	arr.Delete(0)

	for i := 0; i < arr.Length; i++ {
		fmt.Println(arr.Get(i))
	}

}

func runLinkedList() {
	arr := lists.LinkedList{}

	arr.Push(1)
	arr.Push(2)
	arr.Push(6)

	arr.Pop()

	arr.Push(35)
	arr.Push(36)

	arr.Delete(0)
	arr.Delete(0)

	arr.Pop()
	arr.Push(56)
	arr.Push(509)

	for i := 0; i < arr.Length; i++ {
		node := arr.Get(i)
		fmt.Println(node.Value)
	}

}

func runBST() {
	tree := trees.BST{}

	tree.Push(5)
	tree.Push(9)
	tree.Push(3)
	tree.Push(4)
	tree.Push(7)
	tree.Push(12)
	tree.Print()
	tree.Delete(4)
	tree.Delete(7)
	tree.Delete(12)
	tree.Delete(5)
	tree.Print()
	tree.Delete(9)
	tree.Print()

}

type Movie struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ReleaseDate int    `json:"release_date"`
}

func loadData() []Movie {
	file, err := os.Open("data/movies.json")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	movies := []Movie{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&movies)
	if err != nil {
		fmt.Println(err)
	}
	return movies

}

func main() {
	var movies []Movie
	var wg sync.WaitGroup
	serverReady := make(chan bool)

	wg.Add(2)

	go func() {
		defer wg.Done()
		routes.Setup()
		serverReady <- true
		routes.Run()
	}()

	go func() {
		defer wg.Done()
		movies = loadData()
		<-serverReady
		// send data to server
		for i := 0; i < len(movies); i++ {
			jsonData, err := json.Marshal(movies[i])
			if err != nil {
				continue
			}

			_, err = http.Post("http://localhost:5000/v1/trees", "application/json", bytes.NewBuffer(jsonData))
			if err != nil {
				//handle error
			}

		}
	}()
	wg.Wait()
}
