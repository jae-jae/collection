package collect

import (
	"fmt"
	"reflect"
)

// MapCollection is a collection of maps.
type MapCollection[K comparable, V any] struct {
	data map[K]V
}

// NewMap returns a new MapCollection.
func NewMap[K comparable, V any](data map[K]V) *MapCollection[K, V] {
	return &MapCollection[K, V]{
		data: data,
	}
}

// Keys returns the keys of the collection.
func (c *MapCollection[K, V]) Keys() *SliceCollection[K] {
	keys := make([]K, 0, len(c.data))
	for k := range c.data {
		keys = append(keys, k)
	}

	return NewSlice[K](keys)
}

// Values returns the values of the collection.
func (c *MapCollection[K, V]) Values() *SliceCollection[V] {
	values := make([]V, 0, len(c.data))
	for _, v := range c.data {
		values = append(values, v)
	}

	return NewSlice[V](values)
}

// ToMap returns the map.
func (c *MapCollection[K, V]) ToMap() map[K]V {
	return c.data
}

// HasKey returns true if the map contains the keys.
func (c *MapCollection[K, V]) HasKey(keys ...K) bool {
	for _, item := range keys {
		_, ok := c.data[item]
		if !ok {
			return false
		}
	}

	return true
}

// HasAnyKey returns true if the map contains any of the keys.
func (c *MapCollection[K, V]) HasAnyKey(keys ...K) bool {
	for _, item := range keys {
		_, ok := c.data[item]
		if ok {
			return true
		}
	}

	return false
}

// Contains returns true if the map contains the values.
func (c *MapCollection[K, V]) Contains(values ...V) bool {
	return c.Values().Contains(values...)
}

// ContainsAny returns true if the map contains any of the values.
func (c *MapCollection[K, V]) ContainsAny(values ...V) bool {
	return c.Values().ContainsAny(values...)
}

// Equals returns true if the map equals the values.
func (c *MapCollection[K, V]) Equals(data map[K]V) bool {
	return reflect.DeepEqual(c.data, data)
}

// Get returns the value for the key.
func (c *MapCollection[K, V]) Get(key K) (V, bool) {
	value, ok := c.data[key]
	return value, ok
}

// MustGet returns the value for the key or panics if the key does not exist.
func (c *MapCollection[K, V]) MustGet(key K) V {
	value, ok := c.data[key]
	if !ok {
		panic(fmt.Sprintf("key [%v] not exist", key))
	}

	return value
}

// Set sets the value for the key.
func (c *MapCollection[K, V]) Set(key K, value V) *MapCollection[K, V] {
	c.data[key] = value
	return c
}

// Delete deletes the value for the key.
func (c *MapCollection[K, V]) Delete(key K) *MapCollection[K, V] {
	delete(c.data, key)
	return c
}

// Clear clears the map.
func (c *MapCollection[K, V]) Clear() *MapCollection[K, V] {
	c.data = make(map[K]V)
	return c
}

// Size returns the size of the map.
func (c *MapCollection[K, V]) Size() int {
	return len(c.data)
}

// IsEmpty returns true if the map is empty.
func (c *MapCollection[K, V]) IsEmpty() bool {
	return len(c.data) == 0
}

// Clone returns a clone of the map.
func (c *MapCollection[K, V]) Clone() *MapCollection[K, V] {
	m := make(map[K]V, len(c.data))
	for k, v := range m {
		m[k] = v
	}
	return NewMap[K, V](m)
}

// Merge merges the map with the given map.
func (c *MapCollection[K, V]) Merge(items ...map[K]V) *MapCollection[K, V] {
	for _, item := range items {
		for k, v := range item {
			c.data[k] = v
		}
	}
	return c
}

// MergeBy merges the map with the given map using the given function.
func (c *MapCollection[K, V]) MergeBy(mergeFn func(V, V) V, items ...map[K]V) *MapCollection[K, V] {
	for _, item := range items {
		for k, v := range item {
			if c.HasKey(k) {
				c.data[k] = mergeFn(c.data[k], v)
			} else {
				c.data[k] = v
			}
		}
	}

	return c
}

// Filter filters the map with the given function.
func (c *MapCollection[K, V]) Filter(predicate func(key K, value V) bool) *MapCollection[K, V] {
	filtered := make(map[K]V)
	for key, value := range c.data {
		if predicate(key, value) {
			filtered[key] = value
		}
	}
	return NewMap[K, V](filtered)
}

// Map maps the map with the given function.
func (c *MapCollection[K, V]) Map(mapper func(key K, value V) V) *MapCollection[K, V] {
	mapped := make(map[K]V)
	for key, value := range c.data {
		mapped[key] = mapper(key, value)
	}
	return NewMap[K, V](mapped)
}

// MapKeys maps the keys with the given function.
func (c *MapCollection[K, V]) MapKeys(mapper func(key K, value V) K) *MapCollection[K, V] {
	mapped := make(map[K]V)
	for key, value := range c.data {
		mapped[mapper(key, value)] = value
	}
	return NewMap[K, V](mapped)
}

// Reduce reduces the map with the given function.
func (c *MapCollection[K, V]) Reduce(reducer func(key K, value V, previous V) V, initial V) V {
	previous := initial
	for key, value := range c.data {
		previous = reducer(key, value, previous)
	}
	return previous
}

// Each executes the given function for each item in the map.
func (c *MapCollection[K, V]) Each(fn func(key K, value V)) *MapCollection[K, V] {
	for key, value := range c.data {
		fn(key, value)
	}

	return c
}

// Find returns the first item that matches the given predicate.
func (c *MapCollection[K, V]) Find(fn func(key K, value V) bool) (K, V, bool) {
	for key, value := range c.data {
		if fn(key, value) {
			return key, value, true
		}
	}

	return zero[K](), zero[V](), false
}

// FindAll returns all items that match the given predicate.
func (c *MapCollection[K, V]) FindAll(fn func(key K, value V) bool) *MapCollection[K, V] {
	results := make(map[K]V)
	for key, value := range c.data {
		if fn(key, value) {
			results[key] = value
		}
	}
	return NewMap[K, V](results)
}
