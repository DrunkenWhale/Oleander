package util

func Filter[T any](array []T, f func(x T) bool) []T {
	res := make([]T, 0)
	for _, t := range array {
		if f(t) {
			res = append(res, t)
		}
	}
	return res
}
