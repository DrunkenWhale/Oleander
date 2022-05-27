package tpe

type Try[T any] struct {
	left  error
	right T
}
