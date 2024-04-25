package recursion

func NestedAdd(arr []interface{}, index int) int {
	if index > len(arr)-1 {
		return 0
	}
	if nestedArr, ok := arr[index].([]interface{}); ok {
		return NestedAdd(nestedArr, 0) + NestedAdd(arr, index+1)
	}
	return arr[index].(int) + NestedAdd(arr, index+1)
}
