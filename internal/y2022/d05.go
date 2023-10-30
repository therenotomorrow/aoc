// Package day05: https://adventofcode.com/2022/day/5
package main

import (
	"strconv"
	"strings"
)

type Stack struct {
	head *node
}

type node struct {
	val  string
	next *node
}

func (s *Stack) Push(val string) {
	s.head = &node{val: val, next: s.head}
}

func (s *Stack) Pop() string {
	n := s.head
	s.head = s.head.next

	return n.val
}

func (s *Stack) String() string {
	return s.head.val
}

func NewStack() Stack {
	return Stack{}
}

func NewStackOfCrates(matrix [][]string) []Stack {
	crates := make([]Stack, len(matrix[0]))

	for i := 0; i < len(matrix[0]); i++ {
		crates[i] = NewStack()

		for j := len(matrix) - 1; j >= 0; j-- {
			if crate := matrix[j][i]; crate != "" {
				crates[i].Push(crate)
			}
		}
	}

	return crates
}

func PerformMoves(crates []Stack, moves []string) []Stack {
	for _, move := range moves {
		parts := strings.Fields(move)

		size, _ := strconv.Atoi(parts[1])
		from, _ := strconv.Atoi(parts[3])
		to, _ := strconv.Atoi(parts[5])

		for i := 0; i < size; i++ {
			val := crates[from-1].Pop()
			crates[to-1].Push(val)
		}
	}

	return crates
}

func PerformMovesWithOrder(crates []Stack, moves []string) []Stack {
	for _, move := range moves {
		parts := strings.Fields(move)

		size, _ := strconv.Atoi(parts[1])
		from, _ := strconv.Atoi(parts[3])
		to, _ := strconv.Atoi(parts[5])

		addStack := NewStack()

		for i := 0; i < size; i++ {
			val := crates[from-1].Pop()
			addStack.Push(val)
		}

		for i := 0; i < size; i++ {
			val := addStack.Pop()
			crates[to-1].Push(val)
		}
	}

	return crates
}
