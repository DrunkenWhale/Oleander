package tpe

type Stream[T any] struct {
	// bool return value was used to mark
	// is this element be filtered
	mappingRule func(x T) (T, bool)

	generateRule func(i int) T
}

func (stream *Stream[T]) Map(newMappingRule func(x T) T) {
	if stream.mappingRule == nil {
		stream.mappingRule = func(x T) (T, bool) {
			return newMappingRule(x), true
		}
	} else {
		stream.mappingRule = func(x T) (T, bool) {
			tmp, ok := stream.mappingRule(x)
			if ok {
				return newMappingRule(tmp), true
			} else {
				return nil, false
			}
		}
	}
}

func (stream *Stream[T]) Filter(newFilterRule func(x T) bool) {
	if stream.mappingRule == nil {
		stream.mappingRule = func(x T) (T, bool) {
			if newFilterRule(x) {
				return x, true
			} else {
				return nil, false
			}
		}
	} else {
		stream.mappingRule = func(x T) (T, bool) {
			res, ok := stream.mappingRule(x)
			if ok {
				if newFilterRule(res) {
					return res, true
				} else {
					return nil, false
				}
			} else {
				return nil, false
			}
		}
	}
}

func (stream *Stream[T]) Generator(generatorRule func(index int) T) {
	stream.generateRule = generatorRule
}

func (stream *Stream[T]) Build(elementSize int) []T {
	res := make([]T, elementSize)
	res = res[0:0]
	index := 0
	for index < elementSize {
		tmp, ok := stream.mappingRule(stream.generateRule(index))
		if ok {
			res = append(res, tmp)
			index++
		}
	}
	return res
}
