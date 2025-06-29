package value

type SliceValue struct {
	InternalValue interface{} // []int, []float64, etc.
	ElementType   string      // "int", "float", etc.
}

func (s *SliceValue) Type() string {
	return "[]" + s.ElementType
}

func (s *SliceValue) Value() interface{} {
	return s.InternalValue
}

func (s *SliceValue) Copy() IVOR {
	// Implementar la copia profunda según el tipo
	switch v := s.InternalValue.(type) {
	case []int:
		newSlice := make([]int, len(v))
		copy(newSlice, v)
		return &SliceValue{
			InternalValue: newSlice,
			ElementType:   s.ElementType,
		}
	case []float64:
		// Similar para float64
		// Otros tipos...
	}
	return s
}
