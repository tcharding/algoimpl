// Package sortable is my own implementation of different sort functions
// The package will be similar, if not exactly the same as the go provided
// sort package. I am using that to reference where I need to improve my code.
//
// Another reason for this is that, really, you need Len, Less and Swap if you
// are to make a generic sort package
package sort

type Sortable interface {
  // Len is the number of elements in the collection
  Len() int
  // LT returns whether the element at index i should
  // sort before the element at index j
  LT(i, j int) bool
  // Swaps the elements at indices i and j
  Swap(i, j int)
}

/*
   Insertion sort on Sortable type
   Does not have start and end indices yet, like the Go authors of sort
 */
func InsertionSort(stuff Sortable) {
	for j := 1; j < stuff.Len(); j++ { // from the second spot to the last
    for i := j; i > 0 && stuff.LT(i, i-1); i-- { // while num left is larger
      stuff.Swap(i, i-1) // slide that number right one position
		}
	}
}
