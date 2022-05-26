package tpe

type Option[T any] struct {
	value   T
	isEmpty bool
}

func NewOption[T any](value T) *Option[T] {
	return &Option[T]{
		value:   value,
		isEmpty: false,
	}
}

func EmptyOption[T any]() *Option[T] {
	return &Option[T]{
		isEmpty: true,
	}
}

func (opt *Option[T]) Get() T {
	return opt.value
}

func (opt *Option[T]) IsEmpty() bool {
	return opt.isEmpty
}

func (opt *Option[T]) NonEmpty() bool {
	return !opt.isEmpty
}
