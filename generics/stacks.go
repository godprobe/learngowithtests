package generics

type Stack[T any] struct {
	values []T
}

// Push
func (s *Stack[T]) Push(value T) {
	s.values = append(s.values, value)
}

// Pop
func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}

	index := len(s.values) - 1
	el := s.values[index]
	s.values = s.values[:index]
	return el, true
}

// IsEmpty
func (s *Stack[T]) IsEmpty() bool {
	return len(s.values) == 0
}
