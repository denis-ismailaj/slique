package slique

import "testing"

func TestEnqueueDequeue(t *testing.T) {
	initialSize := 50
	sq := Slique[int]{}

	// Insert enough items to make the queue grow a couple of times.
	n := initialSize * 10

	var input []int

	for i := 0; i < n; i++ {
		sq.Enqueue(i)
		input = append(input, i)
	}

	for i := 0; i < n; i++ {
		if sq.Dequeue(1)[0] != input[i] {
			t.Fatal("Items missing or reordered.")
		}
	}

	// Test for panic when dequeuing when empty.
	sq.Dequeue(50)
}
