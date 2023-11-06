package datastruct

type Set interface {
	Add(val interface{})
	Del(val interface{})
	Populate(vals ...interface{})
	Contains(val interface{}) (ok bool)
	Size() int
	Values() []interface{}
}

type set struct {
	data map[interface{}]bool
}

func NewSet() Set {
	return &set{data: make(map[interface{}]bool)}
}

func (s *set) Add(val interface{}) {
	if s.data[val] {
		return
	}

	s.data[val] = true
}

func (s *set) Del(val interface{}) {
	delete(s.data, val)
}

func (s *set) Contains(val interface{}) (ok bool) {
	return s.data[val]
}

func (s *set) Size() int {
	return len(s.data)
}

func (s *set) Populate(vals ...interface{}) {
	for _, val := range vals {
		s.Add(val)
	}
}

func (s *set) Values() []interface{} {
	vals := make([]interface{}, 0, s.Size())

	for k := range s.data {
		vals = append(vals, k)
	}

	return vals
}
