package datastruct

type Stack interface {
	Push(val interface{})
	Pop() (val interface{}, ok bool)
	Peek() (val interface{}, ok bool)
	Size() int
	IsEmpty() bool
}

type sNode struct {
	val  interface{}
	next *sNode
}

type stack struct {
	head *sNode
	size int
}

func NewStack() Stack {
	return &stack{}
}

func (s *stack) Push(val interface{}) {
	s.head = &sNode{val: val, next: s.head}
	s.size++
}

func (s *stack) Pop() (val interface{}, ok bool) {
	if s.IsEmpty() {
		return
	}

	val = s.head.val
	s.head = s.head.next
	s.size--

	return val, true
}

func (s *stack) Peek() (val interface{}, ok bool) {
	if s.IsEmpty() {
		return
	}

	return s.head.val, true
}

func (s *stack) Size() int {
	return s.size
}

func (s *stack) IsEmpty() bool {
	return s.Size() == 0
}
