package go_collections

type Set map[interface{}]interface{}

func NewSet(list []interface{}) Set {
	s := make(Set, len(list))
	for _, el := range list {
		s[el] = nil
	}
	return s
}

func (s Set) Size() int {
	return len(s)
}

func (s Set) Empty() bool {
	return len(s) == 0
}

func (s Set) Contains(el interface{}) bool {
	_, exists := s[el]
	return exists
}

func (s Set) ToArray() []interface{} {
	array := make([]interface{}, s.Size())
	var i int
	for el, _ := range s {
		array[i] = el
		i++
	}
	return array
}

func (s *Set) Add(el interface{}) {
	(*s)[el] = nil
}

func (s *Set) AddAll(el []interface{}) {
	for _, e := range el {
		(*s)[e] = nil
	}
}

func (s *Set) Remove(el interface{}) {
	delete(*s, el)
}
