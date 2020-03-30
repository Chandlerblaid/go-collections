package go_collections

const (
	MINFILL   float32 = 0.1
	MAXFILL   float32 = 0.9
	FILLRATIO float32 = 0.8
	MINCAP    int     = 10
)

type Stack struct {
	list         SimpleList
	fillRatio    float32
	minFillRatio float32
	maxFillRatio float32
	minCapacity  int
}

type StackOptionFunc func(*Stack) error

func SetMinFillPrct(p float32) StackOptionFunc {
	return func(s *Stack) error {
		s.minFillRatio = p
		return nil
	}
}

func SetFillRatio(p float32) StackOptionFunc {
	return func(s *Stack) error {
		s.fillRatio = p
		return nil
	}
}

func SetMaxFillPrct(p float32) StackOptionFunc {
	return func(s *Stack) error {
		s.maxFillRatio = p
		return nil
	}
}

func SetMinCapacity(cap int) StackOptionFunc {
	return func(s *Stack) error {
		s.minCapacity = cap
		return nil
	}
}

func NewStack(slice []interface{}, opts ...StackOptionFunc) Stack {
	if slice == nil {
		slice = []interface{}{}
	}

	ret := Stack{
		list:         slice,
		minCapacity:  MINCAP,
		fillRatio:    FILLRATIO,
		minFillRatio: MINFILL,
		maxFillRatio: MAXFILL,
	}

	for _, opt := range opts {
		if err := opt(&ret); err != nil {
			return Stack{}
		}
	}

	ret.rightsize()

	return ret
}

func (s Stack) pressureUpperLimit() int {
	return int(float32(s.list.Size())*1.0 + (1.0 - s.fillRatio*s.maxFillRatio))
}

func (s Stack) pressureLowerLimit() int {
	return int(float32(s.list.Size()) * s.minFillRatio)
}

func (s Stack) idealCapacity() int {
	ideal := int(float32(s.list.Size())*1.0 + (1.0 - s.fillRatio))
	if ideal < s.minCapacity {
		ideal = s.minCapacity
	}
	return ideal
}

func (s *Stack) rightsize() {
	if cap(s.list) < s.minCapacity || s.list.Size() > s.pressureUpperLimit() || s.list.Size() < s.pressureLowerLimit() {
		array := make([]interface{}, s.idealCapacity())

		if s.list.Size() == 0 {
			s.list = array[:0]
			return
		}

		copy(array, s.list)
		s.list = array[:s.list.Size()]
	}
}

func (s Stack) Empty() bool {
	return s.list.Empty()
}

func (s Stack) Size() int {
	return s.list.Size()
}

func (s Stack) Peek() interface{} {
	return s.list.Peek()
}

func (s *Stack) Put(i interface{}) {
	s.list.Put(i)
	s.rightsize()
}

func (s *Stack) Pop() interface{} {
	d := s.list.Pop()
	s.rightsize()
	return d
}
