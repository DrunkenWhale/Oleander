package util

func Map[T any, F any](array []T, f func(x T) F) []F {
	res := make([]F, len(array))
	for i, t := range array {
		res[i] = f(t)
	}
	return res
}
