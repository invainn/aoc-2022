package main

type Queue []string

func (q Queue) IsEmpty() bool {
	return len(q) == 0
}

func (q Queue) Enqueue(s string) Queue {
	return append(q, s)
}

func (q Queue) Dequeue() Queue {
	if q.IsEmpty() {
		panic("Queue is empty!")
	}

	return q[1:]
}

func (q Queue) ContainsDuplicates() bool {
	hashTable := map[string]bool{}
	for _, member := range q {
		if _, ok := hashTable[member]; ok {
			return true
		}

		hashTable[member] = true
	}

	return false
}
