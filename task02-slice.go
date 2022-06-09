package homework

func reverse(input []int64) (result []int64) {
	output := make([]int64, len(input))
	var l = len(input)
	for i := l - 1; i > -1; i-- {
		output[i] = input[l-i-1]
	}
	result = output
	return result
}
