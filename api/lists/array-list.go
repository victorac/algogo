package lists

type Array struct {
	Length int
	data   []int
}

func (arr *Array) Push(value int) {
	arr.data = append(arr.data, value)
	arr.Length++
}

func (arr *Array) Pop() int {
	el := arr.data[len(arr.data)-1]
	arr.data = arr.data[:len(arr.data)-1]
	arr.Length--
	return el
}

func (arr *Array) Get(index int) int {
	return arr.data[index]
}

func (arr *Array) Delete(index int) {
	arr.data = append(arr.data[:index], arr.data[index+1:]...)
	arr.Length--
}
