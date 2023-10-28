package y2015

import (
	"github.com/therenotomorrow/aoc/pkg/datastruct"
	"strconv"
	"strings"
)

func playTurn(seq string) string {
	res := strings.Builder{}
	stack := datastruct.NewStack()

	for _, r := range seq {
		if stack.IsEmpty() {
			stack.Push(r)
			continue
		}

		peek, _ := stack.Peek()
		if peek.(rune) == r {
			stack.Push(r)
			continue
		}

		res.WriteString(strconv.Itoa(stack.Size()))
		res.WriteRune(peek.(rune))

		for !stack.IsEmpty() {
			stack.Pop()
		}

		stack.Push(r)
	}

	if !stack.IsEmpty() {
		peek, _ := stack.Peek()
		res.WriteString(strconv.Itoa(stack.Size()))
		res.WriteRune(peek.(rune))
	}

	return res.String()
}

func Day10Part1(seq string, times int) int {
	for i := 0; i < times; i++ {
		seq = playTurn(seq)
	}

	return len(seq)
}

func Day10Part2(seq string) int {
	// Conway's Game of Life
	for i := 0; i < 50; i++ {
		seq = playTurn(seq)
	}

	return len(seq)
}
