package jslices

// At method takes an integer value and returns the item at that index, allowing for positive and negative integers.
// Negative integers count back from the last item in the array.
//
// Example:
//
//	array1 := []int{5, 12, 8, 130, 44}
//
//	index := 2
//	fmt.Printf("Using an index of %d the item returned is %v", index, At(array1, index))
//	// Expected output: "Using an index of 2 the item returned is 8"
//
//	index = -2
//	fmt.Printf("Using an index of %d item returned is %v", index, At(array1, index))
//	// Expected output: "Using an index of -2 item returned is 130"
//
// (See: https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/at)
func At[ArrayItem any](array []ArrayItem, index int) ArrayItem {
	if index < 0 {
		index = len(array) + index
	}

	return array[index]
}
