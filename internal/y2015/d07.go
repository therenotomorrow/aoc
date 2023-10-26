package y2015

import (
	"github.com/therenotomorrow/aoc/pkg/datastruct"
	"strconv"
	"strings"
)

type Wires struct {
	data map[string]uint16
}

func (w *Wires) connect(s string) (uint16, bool) {
	val, err := strconv.Atoi(s)

	if err == nil {
		return uint16(val), true
	}

	wire, ok := w.data[s]

	return wire, ok
}

func (w *Wires) send(val uint16, where string) {
	w.data[where] = val
}

func (w *Wires) and(a uint16, b uint16, where string) {
	w.data[where] = a & b
}

func (w *Wires) or(a uint16, b uint16, where string) {
	w.data[where] = a | b
}

func (w *Wires) lshift(a uint16, val int, where string) {
	w.data[where] = a << val
}

func (w *Wires) rshift(a uint16, val int, where string) {
	w.data[where] = a >> val
}

func (w *Wires) not(a uint16, where string) {
	w.data[where] = ^a
}

func parseCommand(cmd interface{}) ([]string, string) {
	toWire := cmd.([]string)[1]
	how := strings.Fields(cmd.([]string)[0])

	return how, toWire
}

func Day07Part1(instructions []string, wireID string) uint16 {
	wires := Wires{data: make(map[string]uint16)}
	commands := datastruct.NewQueue()

	for _, inst := range instructions {
		parts := strings.Split(inst, " -> ")

		commands.Enqueue(parts)
	}

	for !commands.IsEmpty() {
		cmd, _ := commands.Dequeue()

		how, toWire := parseCommand(cmd)

		switch len(how) {
		case 1:
			val, ok := wires.connect(how[0])

			if !ok {
				commands.Enqueue(cmd)
				continue
			}

			wires.send(val, toWire)
		case 2:
			val, ok := wires.connect(how[1])

			if !ok {
				commands.Enqueue(cmd)
				continue
			}

			wires.not(val, toWire)
		default:
			left, ok1 := wires.connect(how[0])
			right, ok2 := wires.connect(how[2])

			if !ok1 || !ok2 {
				commands.Enqueue(cmd)
				continue
			}

			switch how[1] {
			case "LSHIFT":
				wires.lshift(left, int(right), toWire)
			case "RSHIFT":
				wires.rshift(left, int(right), toWire)
			case "AND":
				wires.and(left, right, toWire)
			case "OR":
				wires.or(left, right, toWire)
			}
		}
	}

	return wires.data[wireID]
}

func Day07Part2(instructions []string, wireID string) uint16 {
	wires := Wires{data: make(map[string]uint16)}
	commands := datastruct.NewQueue()

	for _, inst := range instructions {
		parts := strings.Split(inst, " -> ")

		commands.Enqueue(parts)
	}

	for !commands.IsEmpty() {
		cmd, _ := commands.Dequeue()

		how, toWire := parseCommand(cmd)

		switch len(how) {
		case 1:
			val, ok := wires.connect(how[0])

			if !ok {
				commands.Enqueue(cmd)
				continue
			}

			wires.send(val, toWire)
		case 2:
			val, ok := wires.connect(how[1])

			if !ok {
				commands.Enqueue(cmd)
				continue
			}

			wires.not(val, toWire)
		default:
			left, ok1 := wires.connect(how[0])
			right, ok2 := wires.connect(how[2])

			if !ok1 || !ok2 {
				commands.Enqueue(cmd)
				continue
			}

			switch how[1] {
			case "LSHIFT":
				wires.lshift(left, int(right), toWire)
			case "RSHIFT":
				wires.rshift(left, int(right), toWire)
			case "AND":
				wires.and(left, right, toWire)
			case "OR":
				wires.or(left, right, toWire)
			}
		}
	}

	return wires.data[wireID]
}
