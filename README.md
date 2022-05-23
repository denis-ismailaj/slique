# Slique

`Slique` (pronounced: slick) is a slice-backed FIFO queue that enables independent reads and writes.

## How it works

Independent reads and writes are achieved by not modifying slice headers when dequeuing.
Instead of re-slicing, a separate index is stored that holds the _reader head_.
Dequeued items are held in memory until slice capacity is full, and at that point they are discarded all at once.

Concurrent reads are policed with a mutex lock.
As long as slice capacity is sufficient a single writer can operate lock-free without disturbing reads.
However, while the slice is being resized no reads can be made.

`Slique` uses generics for the type of its elements (upper bound `any`).

## Room for improvement

To save memory when an index is dequeued even though the index remains in the slice,
an optimization could be made that if `T` is a pointer type it is nullified
(potentially leading to its value being garbage collected), or if it is not it can be replaced
with the zero-value of that type (which may or may not occupy less space).