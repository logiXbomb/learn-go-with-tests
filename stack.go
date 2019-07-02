package main

type Stack struct {
	empty bool
}

func (s *Stack) push(i int) {
	s.empty = false
}

func (s *Stack) pop() {
	panic("waffles")
}

func (s *Stack) isEmpty() bool {
	return s.empty
}

func NewStack() *Stack {
	return &Stack{
		empty: true,
	}
}
