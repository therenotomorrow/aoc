package datastruct_test

import (
	"github.com/therenotomorrow/aoc/pkg/datastruct"
	"testing"
)

func TestStackPushPop(t *testing.T) {
	s := datastruct.NewStack()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	if got, ok := s.Pop(); got != 3 || !ok {
		t.Errorf("Pop() = (%v, %v), want = (%v, %v)", got, ok, 3, true)
	}

	if got, ok := s.Pop(); got != 2 || !ok {
		t.Errorf("Pop() = (%v, %v), want = (%v, %v)", got, ok, 2, true)
	}

	if got, ok := s.Pop(); got != 1 || !ok {
		t.Errorf("Pop() = (%v, %v), want = (%v, %v)", got, ok, 1, true)
	}

	if got, ok := s.Pop(); got != nil || ok {
		t.Errorf("Pop() = (%v, %v), want = (%v, %v)", got, ok, nil, false)
	}
}

func TestStackPeek(t *testing.T) {
	s := datastruct.NewStack()

	s.Push(42)

	if got, ok := s.Peek(); got != 42 || !ok {
		t.Errorf("Peek() = (%v, %v), want = (%v, %v)", got, ok, 42, true)
	}

	s.Pop()

	if got, ok := s.Peek(); got != nil || ok {
		t.Errorf("Peek() = (%v, %v), want = (%v, %v)", got, ok, nil, false)
	}
}

func TestStackIsEmptySize(t *testing.T) {
	s := datastruct.NewStack()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	if got := s.Size(); got != 3 {
		t.Errorf("Size() = %v, want = %v", got, 3)
	}
	if got := s.IsEmpty(); got {
		t.Errorf("IsEmpty() = %v, want = %v", got, false)
	}

	s.Pop()

	if got := s.Size(); got != 2 {
		t.Errorf("Size() = %v, want = %v", got, 2)
	}

	s.Pop()
	s.Pop()

	if got := s.Size(); got != 0 {
		t.Errorf("Size() = %v, want = %v", got, 0)
	}
	if got := s.IsEmpty(); !got {
		t.Errorf("IsEmpty() = %v, want = %v", got, true)
	}
}
