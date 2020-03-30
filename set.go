package go_collections

type Set map[string]interface{}

func NewSet(list []string) *Set {
	if list == nil {
		return new(Set)
	}
	s := make(Set, len(list))
	for _, el := range list {
		s[el] = nil
	}
	return &s
}

func (s Set) Size() int {
	if s == nil {
		return 0
	}
	return len(s)
}

func (s Set) Empty() bool {
	return s == nil || len(s) == 0
}

func (s Set) Contains(el string) bool {
	if s == nil {
		return false
	}
	_, exists := s[el]
	return exists
}

func (s Set) ToArray() []string {
	array := make([]string, s.Size())
	i := 0
	for el := range s {
		array[i] = el
		i++
	}
	return array
}

func (s *Set) Add(el string) {
	if s == nil || *s == nil {
		*s = make(map[string]interface{})
	}
	(*s)[el] = nil
}

func (s *Set) AddAll(el ...string) {
	for _, listE := range el {
		s.Add(listE)
	}
}

func (s *Set) Merge(other Set) {
	for k := range other {
		s.Add(k)
	}
}

func (s *Set) Remove(el string) {
	delete(*s, el)
}
