package utils

import (
	"cmp"
	"sort"
)

// NOTE: T must be comparable to be a map key

// Return keys and values of a map as two slices, in whatever order the map
// iterates over its entries. The Nth key in the returned slice corresponds to
// the Nth value in the returned slice.
func MapKeysAndValues[T comparable, V any](m map[T]V) ([]T, []V) {
	keys := make([]T, 0)
	values := make([]V, 0)
	for k, v := range m {
		keys = append(keys, k)
		values = append(values, v)
	}
	return keys, values
}

// Return keys and values of a map as two slices, in sorted order. The Nth key
// in the returned slice corresponds to the Nth value in the returned slice.
func MapKeysAndValuesSorted[T cmp.Ordered, V any](m map[T]V) ([]T, []V) {
	keys, values := MapKeysAndValues(m)
	cmp := func(i, j int) bool { return keys[i] < keys[j] }
	sort.SliceStable(keys, cmp)
	return keys, values
}

// Return keys and values of a map as two slices, in sorted order. The Nth key
// in the returned slice corresponds to the Nth value in the returned slice.
// The less function is used to compare keys.
func MapKeysAndValuesSortFunc[T comparable, V any](m map[T]V, less func(T, T) bool) ([]T, []V) {
	keys, values := MapKeysAndValues(m)
	cmp := func(i, j int) bool { return less(keys[i], keys[j]) }
	sort.SliceStable(keys, cmp)
	return keys, values
}

// Item is a key-value pair
type Item[T comparable, V any] struct {
	Key   T
	Value V
}

// Return map entries as a slice of Item structs (key-value pairs)
func MapItems[T comparable, V any](m map[T]V) []Item[T, V] {
	items := make([]Item[T, V], 0)
	for k, v := range m {
		items = append(items, Item[T, V]{k, v})
	}
	return items
}

// Return map entries as a slice of Item structs (key-value pairs), in sorted
// order.
func MapItemsSorted[T cmp.Ordered, V any](m map[T]V) []Item[T, V] {
	items := MapItems(m)
	cmp := func(i, j int) bool { return items[i].Key < items[j].Key }
	sort.SliceStable(items, cmp)
	return items
}

// Return map entries as a slice of Item structs (key-value pairs), in sorted
// order. The less function is used to compare keys.
func MapItemsSortFunc[T comparable, V any](m map[T]V, less func(T, T) bool) []Item[T, V] {
	items := MapItems(m)
	cmp := func(i, j int) bool { return less(items[i].Key, items[j].Key) }
	sort.SliceStable(items, cmp)
	return items
}