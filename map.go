package slice

// Map applies the callback fn to the elements of the values.
func Map[T any, Y any](values []T, fn func(v T) Y) []Y {
	v := make([]Y, 0, len(values))
	for i := 0; i < len(values); i++ {
		v = append(v, fn(values[i]))
	}
	return v
}
