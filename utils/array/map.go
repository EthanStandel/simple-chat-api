package array

func Map[T any, O any](input []T, callback func(T) O) []O {
	output := make([]O, len(input))
	for i := range input {
		output[i] = callback(input[i])
	}

	return output
}
