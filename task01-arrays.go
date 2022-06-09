package homework

func average(input [15]float32) (result float32) {
	var N uint32 = 0
	var sum float32 = 0
	for _, elem := range input {
		if elem > 0 {
			sum += elem
			N++
		}
	}
	result = sum / float32(N)
	return result
}
