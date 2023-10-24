package datastruct_test

import (
	"github.com/therenotomorrow/aoc/pkg/datastruct"
	"testing"
)

func TestSetAddContainsSize(t *testing.T) {
	s := datastruct.NewSet()

	// empty set
	if got := s.Size(); got != 0 {
		t.Errorf("Size() = %v, want = %v", got, 0)
	}

	s.Add(42)
	s.Add("asd")

	if got := s.Size(); got != 2 {
		t.Errorf("Size() = %v, want = %v", got, 2)
	}

	// regular usage `.Add()`
	if is42, is36 := s.Contains(42), s.Contains("asd"); !is42 || !is36 {
		t.Errorf("Add() = (%v, %v), want = (%v, %v)", is42, is36, true, true)
	}

	s.Add(42)

	// override value by `.Add()`
	if got := s.Size(); got != 2 {
		t.Errorf("Size() = %v, want = %v", got, 2)
	}
}

func TestSetPopulate(t *testing.T) {
	s := datastruct.NewSet()

	s.Populate(1, 2, 3, 4, 5)

	for i := 1; i <= 5; i++ {
		if got := s.Contains(i); !got {
			t.Errorf("Populate() = %v, want = %v", got, i)
		}
	}
}
