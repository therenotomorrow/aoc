// Package day11: https://adventofcode.com/2022/day/11
package day11

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Operation struct {
	leftOp  string
	sign    string
	rightOp string
}

type Test struct {
	Div int
	T   *Monkey
	F   *Monkey
}

type Monkey struct {
	ID        int
	Items     []int
	Operation Operation
	Test      Test
	Inspects  int
}

func (m *Monkey) String() string {
	return fmt.Sprintf(`
Monkey %d:
  Starting items: %v
  Operation: new = %v
  Test: divisible by %d
    If true: throw to monkey %d
    If false: throw to monkey %d
`, m.ID, m.Items, m.Operation, m.Test.Div, m.Test.T.ID, m.Test.F.ID)
}

func NewMonkey() *Monkey {
	return &Monkey{}
}

func ParseMonkeyID(in string) int {
	re := regexp.MustCompile(`(?i)monkey (\d+)`)
	id, _ := strconv.Atoi(re.FindStringSubmatch(in)[1])
	return id
}

func ParseStartingItems(in string) []int {
	re := regexp.MustCompile(`Starting items: (.*)`)

	items := make([]int, 0)

	for _, i := range strings.Split(re.FindStringSubmatch(in)[1], ", ") {
		item, _ := strconv.Atoi(i)

		items = append(items, item)
	}

	return items
}

func ParseOperation(in string) Operation {
	re := regexp.MustCompile(`Operation: new = (.*)`)

	parts := strings.Fields(re.FindStringSubmatch(in)[1])

	return Operation{parts[0], parts[1], parts[2]}
}

func ParseTestDiv(in string) int {
	re := regexp.MustCompile(`Test: divisible by (\d+)`)
	div, _ := strconv.Atoi(re.FindStringSubmatch(in)[1])
	return div
}

func (m *Monkey) calcWorry1(val int) int {
	var right int

	if m.Operation.rightOp == "old" {
		right = val
	} else {
		right, _ = strconv.Atoi(m.Operation.rightOp)
	}

	switch m.Operation.sign {
	case "*":
		val *= right
	case "+":
		val += right
	}

	val /= 3
	m.Inspects++
	return val
}

func (m *Monkey) calcWorry2(val int, div int) int {
	var right int

	if m.Operation.rightOp == "old" {
		right = val
	} else {
		right, _ = strconv.Atoi(m.Operation.rightOp)
	}

	switch m.Operation.sign {
	case "*":
		val *= right
	case "+":
		val += right
	}
	if val != val%div {
		fmt.Println(m.Test.Div, val, val%div)
	}
	val %= div
	m.Inspects++
	return val
}

func (m *Monkey) TestIt(item int) {
	if item%m.Test.Div == 0 {
		m.Test.T.Items = append(m.Test.T.Items, item)
	} else {
		m.Test.F.Items = append(m.Test.F.Items, item)
	}
}

func PlayRound(monkeys []*Monkey, r bool) {
	mul := MUL(monkeys)

	for _, m := range monkeys {
		for len(m.Items) != 0 {
			val := m.Items[0]
			m.Items = m.Items[1:]

			var newItem int
			if r {
				newItem = m.calcWorry1(val)
			} else {
				newItem = m.calcWorry2(val, mul)
			}
			m.TestIt(newItem)
		}
	}
}

func MUL(monkeys []*Monkey) int {
	div := 1

	for _, m := range monkeys {
		div *= m.Test.Div
	}

	return div
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	data, _ := io.ReadAll(f)

	input := strings.Split(string(data), "\n\n")

	monkeys := make([]*Monkey, 0)
	for i := 0; i < len(input); i++ {
		monkeys = append(monkeys, NewMonkey())
	}

	for _, ms := range input {
		parts := strings.Split(ms, "\n")

		mID := ParseMonkeyID(parts[0])
		mItems := ParseStartingItems(parts[1])
		mOperation := ParseOperation(parts[2])
		mTestDiv := ParseTestDiv(parts[3])
		mTestT := ParseMonkeyID(parts[4])
		mTestF := ParseMonkeyID(parts[5])

		monkeys[mID].ID = mID
		monkeys[mID].Items = mItems
		monkeys[mID].Operation = mOperation
		monkeys[mID].Test = Test{Div: mTestDiv, T: monkeys[mTestT], F: monkeys[mTestF]}
	}

	for i := 0; i < 1000; i++ {
		PlayRound(monkeys, false)
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].Inspects > monkeys[j].Inspects
	})

	fmt.Println(monkeys[0].Inspects * monkeys[1].Inspects)
}
