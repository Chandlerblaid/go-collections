package go_collections

type SimpleList []interface{}

func (s SimpleList) Empty() bool {
	return len(s) == 0
}

func (s SimpleList) Size() int {
	return len(s)
}

func (s SimpleList) Peek() interface{} {
	return s[len(s)-1]
}

func (s SimpleList) PeekFront() interface{} {
	return s[0]
}

func (s *SimpleList) Put(i interface{}) {
	*s = append(*s, i)
}

func (s *SimpleList) Pop() interface{} {
	d := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return d
}

func (s *SimpleList) PopFront() interface{} {
	d := (*s)[0]
	*s = (*s)[1:]
	return d
}
