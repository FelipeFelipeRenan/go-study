package datastructures

type Stack struct {
	elements []any
}


func (s *Stack)Push(el any){
	s.elements = append(s.elements, el)
}