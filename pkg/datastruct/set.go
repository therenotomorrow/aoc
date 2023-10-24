package datastruct

type Set interface {
	Add(val interface{})
	Populate(vals ...interface{})
	Contains(val interface{}) (ok bool)
	Size() int
}

type set struct {
	data map[interface{}]bool
	size int
}

func NewSet() Set {
	return &set{data: make(map[interface{}]bool)}
}

func (s *set) Add(val interface{}) {
	if s.data[val] {
		return
	}

	s.data[val] = true
	s.size++
}

func (s *set) Contains(val interface{}) (ok bool) {
	return s.data[val]
}

func (s *set) Size() int {
	return s.size
}

func (s *set) Populate(vals ...interface{}) {
	for _, val := range vals {
		s.Add(val)
	}
}
