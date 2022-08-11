package array

func Find[T any](input []T, callback func(T) bool) (T, int) {
	for i := range input {
		if callback(input[i]) {
			return input[i], i
		}
	}

	return *new(T), -1
}
