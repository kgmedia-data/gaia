package repo

func convertSlice[T any, R any](input []T, convertFunc func(T) R) []R {
	var result []R
	for _, item := range input {
		result = append(result, convertFunc(item))
	}
	return result
}
