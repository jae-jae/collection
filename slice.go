package collect

import "reflect"

type SliceCollection[T any] struct {
	data []T
}

func NewSlice[T any](data []T) *SliceCollection[T] {
	return &SliceCollection[T]{
		data: data,
	}
}

// ToSlice returns the slice.
func (c *SliceCollection[T]) ToSlice() []T {
	return c.data
}

// Contains returns true if the slice contains the values.
func (c *SliceCollection[T]) Contains(values ...T) bool {
	hasValue := func(value T) bool {
		for _, item := range c.data {
			if reflect.DeepEqual(item, value) {
				return true
			}
		}

		return false
	}

	for _, item := range values {
		if !hasValue(item) {
			return false
		}
	}

	return true
}

// ContainsAny returns true if the slice contains any of the values.
func (c *SliceCollection[T]) ContainsAny(values ...T) bool {
	for _, item := range values {
		if c.Contains(item) {
			return true
		}
	}

	return false
}

// Equals returns true if the slice equals the values.
func (c *SliceCollection[T]) Equals(data []T) bool {
	return reflect.DeepEqual(c.data, data)
}

// HasKey returns true if the slice contains the keys.
func (c *SliceCollection[T]) HasKey(keys ...int) bool {
	for _, item := range keys {
		if item < 0 || item >= len(c.data) {
			return false
		}
	}

	return true
}

// HasAnyKey returns true if the slice contains any of the keys.
func (c *SliceCollection[T]) HasAnyKey(keys ...int) bool {
	for _, item := range keys {
		if item >= 0 && item < len(c.data) {
			return true
		}
	}

	return false
}

// Get returns the value for the key.
func (c *SliceCollection[T]) Get(key int) (T, bool) {
	if !c.HasKey(key) {
		return zero[T](), false
	}

	return c.data[key], true
}

// MustGet returns the value for the key or panics if the key does not exist.
func (c *SliceCollection[T]) MustGet(key int) T {
	return c.data[key]
}

// Set sets the value for the key.
func (c *SliceCollection[T]) Set(key int, value T) *SliceCollection[T] {
	c.data[key] = value
	return c
}

// Push adds the value to the end of the slice.
func (c *SliceCollection[T]) Push(values ...T) *SliceCollection[T] {
	c.data = append(c.data, values...)
	return c
}

// Pop removes and returns the last value of the slice.
func (c *SliceCollection[T]) Pop() T {
	value := c.data[len(c.data)-1]
	c.data = c.data[:len(c.data)-1]
	return value
}

// Shift removes and returns the first value of the slice.
func (c *SliceCollection[T]) Shift() T {
	value := c.data[0]
	c.data = c.data[1:]
	return value
}

// Unshift adds the value to the beginning of the slice.
func (c *SliceCollection[T]) Unshift(values ...T) *SliceCollection[T] {
	c.data = append(values, c.data...)
	return c
}

// Reverse reverses the slice.
func (c *SliceCollection[T]) Reverse() *SliceCollection[T] {
	for i, j := 0, len(c.data)-1; i < j; i, j = i+1, j-1 {
		c.data[i], c.data[j] = c.data[j], c.data[i]
	}

	return c
}

// Each iterates over the slice.
func (c *SliceCollection[T]) Each(fn func(i int, value T)) *SliceCollection[T] {
	for i, value := range c.data {
		fn(i, value)
	}

	return c
}

// Map iterates over the slice and returns a new slice with the results.
func (c *SliceCollection[T]) Map(fn func(i int, value T) T) *SliceCollection[T] {
	for i, value := range c.data {
		c.data[i] = fn(i, value)
	}

	return c
}

// Filter iterates over the slice and returns a new slice with the results.
func (c *SliceCollection[T]) Filter(fn func(i int, value T) bool) *SliceCollection[T] {
	for i, value := range c.data {
		if fn(i, value) {
			continue
		}

		c.data = append(c.data[:i], c.data[i+1:]...)
	}

	return c
}

// Merge merges the slice with the values.
func (c *SliceCollection[T]) Merge(values ...[]T) *SliceCollection[T] {
	for _, value := range values {
		c.data = append(c.data, value...)
	}

	return c
}

// Chunk splits the slice into smaller slices.
func (c *SliceCollection[T]) Chunk(size int) [][]T {
	chunks := make([][]T, 0)

	for i := 0; i < len(c.data); i += size {
		chunks = append(chunks, c.data[i:min(i+size, len(c.data))])
	}

	return chunks
}
