package tpe

type Either[T any] struct {
	left  error
	right T
}

func Left[T any](err error) *Either[T] {
	return &Either[T]{
		left: err,
	}
}

func Right[T any](right T) *Either[T] {
	return &Either[T]{
		left:  nil,
		right: right,
	}
}

func NewEither[T any](left error, right T) *Either[T] {
	return &Either[T]{
		left:  left,
		right: right,
	}
}

func (either *Either[T]) Left() error {
	return either.left
}

func (either *Either[T]) Right() T {
	return either.right
}
