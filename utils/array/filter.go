package array

func Filter[T any](input []T, callback func(T) bool) []T {
	output := []T{}
	for i := range input {
		if callback(input[i]) {
			output = append(output, input[i])
		}
	}

	return output
}
