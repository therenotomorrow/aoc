package datastruct_test

import (
	"github.com/therenotomorrow/aoc/pkg/datastruct"
	"reflect"
	"testing"
)

func TestQueueEnqueueDequeue(t *testing.T) {
	q := datastruct.NewQueue()

	q.Enqueue([]string{"1", "2"})
	q.Enqueue([]string{"2", "3"})
	q.Enqueue([]string{"3", "4"})

	if got, ok := q.Dequeue(); !reflect.DeepEqual(got, []string{"1", "2"}) || !ok {
		t.Errorf("Dequeue() = (%v, %v), want = (%v, %v)", got, ok, []string{"1", "2"}, true)
	}

	if got, ok := q.Dequeue(); !reflect.DeepEqual(got, []string{"2", "3"}) || !ok {
		t.Errorf("Dequeue() = (%v, %v), want = (%v, %v)", got, ok, []string{"2", "3"}, true)
	}

	if got, ok := q.Dequeue(); !reflect.DeepEqual(got, []string{"3", "4"}) || !ok {
		t.Errorf("Dequeue() = (%v, %v), want = (%v, %v)", got, ok, []string{"3", "4"}, true)
	}

	if got, ok := q.Dequeue(); got != nil || ok {
		t.Errorf("Dequeue() = (%v, %v), want = (%v, %v)", got, ok, nil, false)
	}
}

func TestQueueIsEmptySize(t *testing.T) {
	s := datastruct.NewQueue()

	s.Enqueue(1)
	s.Enqueue(2)
	s.Enqueue(3)

	if got := s.Size(); got != 3 {
		t.Errorf("Size() = %v, want = %v", got, 3)
	}
	if got := s.IsEmpty(); got {
		t.Errorf("IsEmpty() = %v, want = %v", got, false)
	}

	s.Dequeue()

	if got := s.Size(); got != 2 {
		t.Errorf("Size() = %v, want = %v", got, 2)
	}

	s.Dequeue()
	s.Dequeue()

	if got := s.Size(); got != 0 {
		t.Errorf("Size() = %v, want = %v", got, 0)
	}
	if got := s.IsEmpty(); !got {
		t.Errorf("IsEmpty() = %v, want = %v", got, true)
	}
}
