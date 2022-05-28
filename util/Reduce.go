package util

func Reduce[T any](array []T, f func(x T, res T) T) T {
	var res T
	for i, t := range array {
		if i == 0 {
			res = t
		} else {
			res = f(t, res)
		}
	}
	return res
}
