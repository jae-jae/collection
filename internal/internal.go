package internal

// Min returns the minimum of the values.
func Min[T Ordered](values ...T) T {
	if len(values) == 0 {
		var zero T
		return zero
	}

	minValue := values[0]
	for _, value := range values {
		if value < minValue {
			minValue = value
		}
	}

	return minValue
}

// Zero returns the zero value of the type.
func Zero[T any]() T {
	var noop T
	return noop
}
