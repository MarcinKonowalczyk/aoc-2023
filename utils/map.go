package utils

// NOTE: T must be comparable to be a map key

// Return values of a map
func MapValues[T comparable, V any](m map[T]V) []V {
	values := make([]V, 0)
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

// Return keys of a map
func MapKeys[T comparable, V any](m map[T]V) []T {
	keys := make([]T, 0)
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// MapEntry is a key-value pair
type MapEntry[T comparable, V any] struct {
	Key   T
	Value V
}

// Return map entries as a slice of MapEntry structs (key-value pairs)
func MapItems[T comparable, V any](m map[T]V) []MapEntry[T, V] {
	items := make([]MapEntry[T, V], 0)
	for k, v := range m {
		items = append(items, MapEntry[T, V]{k, v})
	}
	return items
}
