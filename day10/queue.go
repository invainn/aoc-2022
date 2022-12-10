package main

type Queue[T any] []T

func (q Queue[T]) IsEmpty() bool {
	return len(q) == 0
}

func (q Queue[T]) Enqueue(t T) Queue[T] {
	return append(q, t)
}

func (q Queue[T]) Dequeue() (Queue[T], T) {
	if q.IsEmpty() {
		panic("Queue is empty!")
	}

	return q[1:], q[0]
}
