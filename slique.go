package slique

import (
	"sync"
)

type Slique[T any] struct {
	slice  []T
	head   int // Index to the reader head on the slice
	readMu sync.Mutex
}

func (q *Slique[T]) Enqueue(elements ...T) {
	origLen := len(q.slice)

	// Check if growing capacity is needed.
	// If that is the case, get a read lock and repair head index.
	total := origLen + len(elements)
	if total > cap(q.slice) {
		q.readMu.Lock()
		defer q.readMu.Unlock()

		// We exhausted the current backing array but that doesn't mean we've used up all the capacity
		// because reads leave empty places in the beginning which will now be discarded.
		// Calculate the real capacity needed for the new array.
		realLen := origLen - q.head
		realTotal := realLen + len(elements)

		// Triple the size (+1 in case it's 0)
		newSize := realTotal*3 + 1
		newSlice := make([]T, realTotal, newSize)

		// Copy the existing data
		copy(newSlice, q.slice[q.head:])
		q.slice = newSlice

		// Reset head to first index
		q.head = 0
	}

	q.slice = q.slice[:total]
	copy(q.slice[origLen:], elements)
}

func (q *Slique[T]) Dequeue(n int) []T {
	q.readMu.Lock()
	defer q.readMu.Unlock()

	var t []T

	for i := 0; i < n; i++ {
		if q.head >= len(q.slice) {
			// No additional data available to read.
			return t
		}

		// Retrieve item and save it.
		t = append(t, q.slice[q.head])

		// Move read head to the next position.
		q.head++
	}
	return t
}
