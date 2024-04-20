package datastructures

type Set struct {
	elements map[any]bool
}

func NewSet() Set {
	return Set{elements: make(map[any]bool)}
}

func (s *Set) Add(el any) {
	s.elements[el] = false
}

func (s *Set) Remove(el any) {
	delete(s.elements, el)
}

func (s *Set) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Set) Size() int {
	return len(s.elements)
}

func (s *Set) List(list []any) {
	for k := range s.elements {
		list = append(list, k)
	}
	return
}

// Verify if the set has the given element
func (s *Set) Has(el any) (ok bool) {
	_, ok = s.elements[el]
	return
}

func (s *Set) Union(sets ...Set) (u Set) {
	u = sets[0].Copy()
	for _, set := range sets[1:len(sets)] {
		for k := range set.elements {
			u.Add(k)
		}
	}

	return
}

func (s *Set) Copy() (u Set) {
	u = NewSet()
	for k := range s.elements {
		u.Add(k)
	}
	return
}


func Intersect(sets ...Set)(i Set)  {
	i = sets[0].Copy()
	for k := range i.elements{
		for _, set := range sets[1:len(sets)]{
			if !set.Has(k){
				i.Remove(k)
			}
		}
	}
	return
}


func (s Set)Difference(t Set) Set  {
	for k:= range t.elements{
		if s.Has(k){
			delete(s.elements, k)
		}
	}
	return s
}

func (s Set)IsSubset(t Set) bool{
	for k := range s.elements{
		if !t.Has(k) {
			return false
		}
	}
	return true
	
}