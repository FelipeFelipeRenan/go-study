package datastructures

import "errors"

type Stack struct {
	elements []any
}

func (s *Stack) Push(el any) {
	s.elements = append(s.elements, el)
}

func (s *Stack) Pop() (el any, err error) {
	if s.IsEmpty() {
		err = errors.New("empty stack")
		return
	}
	el = s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return 
}

func (s *Stack) Peek() (el any, err error) {
	if s.IsEmpty() {
		err = errors.New("empty stack")
		return
	}
	el = s.elements[len(s.elements)-1]
	return
}

func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack)Size() int  {
	return len(s.elements)
}