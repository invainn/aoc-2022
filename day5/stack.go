package main

type Stack []string

func (s Stack) IsEmpty() bool {
	return len(s) == 0
}

func (s Stack) Push(v string) Stack {
	return append(s, v)
}

func (s Stack) Pop() (Stack, string) {
	if s.IsEmpty() {
		panic("Stack is empty!")
	}

	return s[:len(s)-1], s[len(s)-1]
}

func (s Stack) StackFromStrings(strs []string) Stack {
	for _, str := range strs {
		s = s.Push(str)
	}

	return s
}
