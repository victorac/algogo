package recursion

func Factorial(value int) int {
	if value < 2 {
		return 1
	}
	return value * Factorial(value-1)
}
