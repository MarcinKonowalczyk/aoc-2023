package main

import (
	"fmt"
)

type mykey struct {
	id int
}

type myvalue struct {
	id int
}

func main() {

	my_map := make(map[mykey](chan func()))

	for i := 0; i < 100; i++ {

		myfun := make(chan func())

		my_map[mykey{i}] = myfun
	}

	for k, v := range my_map {
		fmt.Println(k, v)
	}
}

// Returns the keys of a map in sorted order.
// func mapKeysInSortedOrder[T cmp.Ordered, V any](m map[T]V) []T {
// 	keys := make([]T, 0, len(m))
// 	for k := range m {
// 		keys = append(keys, k)
// 	}
// 	sort.Slice(keys, func(i, j int) bool { return keys[i].Less(keys[j]) })
// 	return keys
// }
