package datastruct

type Queue interface {
	Enqueue(val interface{})
	Dequeue() (val interface{}, ok bool)
	Size() int
	IsEmpty() bool
}

type qNode struct {
	val  interface{}
	next *qNode
}

type queue struct {
	head *qNode
	tail *qNode
	size int
}

func NewQueue() Queue {
	return &queue{}
}

func (q *queue) Enqueue(val interface{}) {
	old := q.tail
	q.tail = &qNode{val: val}

	if q.Size() == 0 {
		q.head = q.tail
	} else {
		old.next = q.tail
	}

	q.size++
}

func (q *queue) Dequeue() (val interface{}, ok bool) {
	if q.IsEmpty() {
		return
	}

	val = q.head.val
	q.head = q.head.next
	q.size--

	return val, true
}

func (q *queue) Size() int {
	return q.size
}

func (q *queue) IsEmpty() bool {
	return q.Size() == 0
}
