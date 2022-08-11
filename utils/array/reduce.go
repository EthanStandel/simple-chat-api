package array

func Reduce[T any, O any](input []T, callback func(T, O) O, initial O) O {
	result := initial
	for i := range input {
		result = callback(input[i], result)
	}

	return result
}
