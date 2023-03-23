package generics

type StackOfInts struct {
	values []int
}

// Push
func (s *StackOfInts) Push(value int) {
	s.values = append(s.values, value)
}

// Pop
func (s *StackOfInts) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	index := len(s.values) - 1
	el := s.values[index]
	s.values = s.values[:index]
	return el, true
}

// IsEmpty
func (s *StackOfInts) IsEmpty() bool {
	return len(s.values) == 0
}

type StackOfStrings struct {
	values []string
}

// Push
func (s *StackOfStrings) Push(value string) {
	s.values = append(s.values, value)
}

// Pop
func (s *StackOfStrings) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}

	index := len(s.values) - 1
	el := s.values[index]
	s.values = s.values[:index]
	return el, true
}

// IsEmpty
func (s *StackOfStrings) IsEmpty() bool {
	return len(s.values) == 0
}
