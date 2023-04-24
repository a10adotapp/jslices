package jslices

// Concat is used to merge two or more arrays.
// This method does not change the existing arrays, but instead returns a new array.
//
// Example:
//
//	array1 := []string{"a", "b", "c"}
//	array2 := []string{"d", "e", "f"}
//
// array3 := Concat(array1, array2)
//
//	fmt.Printf("%+v\n", array3)
//	// Expected output: [a b c d e f]
//
// (See: https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/concat)
func Concat[ArrayItem any](array []ArrayItem, arrayOrArrayItems ...any) []ArrayItem {
	for _, arrayOrArrayItem := range arrayOrArrayItems {
		switch arrayOrArrayItem := arrayOrArrayItem.(type) {
		case ArrayItem:
			array = append(array, arrayOrArrayItem)
		case []ArrayItem:
			array = append(array, arrayOrArrayItem...)
		}
	}

	return array
}
