package collect

import (
	"golang.org/x/exp/constraints"
)

// min returns the minimum of the values.
func min[T constraints.Ordered](values ...T) T {
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

// zero returns the zero value of the type.
func zero[T any]() T {
	var noop T
	return noop
}
