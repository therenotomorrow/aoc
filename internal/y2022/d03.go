// Package day03: https://adventofcode.com/2022/day/3
package day03

import (
	"fmt"
	"strings"
	"unicode"
)

func FindDuplicateItem(rucksack string) rune {
	var (
		half      = len(rucksack) / 2
		items     = make(map[rune]struct{})
		duplicate rune
	)

	for _, item := range rucksack[:half] {
		items[item] = struct{}{}
	}

	for _, item := range rucksack[half:] {
		if _, ok := items[item]; ok {
			duplicate = item
			break
		}
	}

	return duplicate
}

func GetPriority(r rune) int {
	var cut int

	if unicode.IsUpper(r) {
		cut = 38
	} else {
		cut = 96
	}

	return int(r) - cut
}

type Set struct {
	state map[rune]struct{}
}

func NewSet() Set {
	return Set{state: make(map[rune]struct{})}
}

func (s *Set) Set(r rune) {
	s.state[r] = struct{}{}
}

func (s *Set) Pop() rune {
	for key := range s.state {
		return key
	}

	return 0
}

func (s *Set) Intersect(other Set) Set {
	res := NewSet()

	for key := range other.state {
		_, ok := s.state[key]
		if ok {
			res.Set(key)
		}
	}

	return res
}

func (s *Set) String() string {
	var sb strings.Builder

	for k := range s.state {
		sb.WriteString(fmt.Sprintf("|%c|", k))
	}

	return sb.String()
}

func FindDuplicateItemInGroup(rucksacks []string) rune {
	var sets []Set

	for _, rucksack := range rucksacks {
		set := NewSet()

		for _, item := range rucksack {
			set.Set(item)
		}

		sets = append(sets, set)
	}

	fs := sets[0]
	for _, set := range sets[1:] {
		fs = fs.Intersect(set)
	}

	return fs.Pop()
}

func SumDuplicateItems(rucksacks []string) int {
	total := 0

	for _, ruck := range rucksacks {
		item := FindDuplicateItem(ruck)

		total += GetPriority(item)
	}

	return total
}

func SumDuplicateGroups(rucksacks []string) int {
	total := 0

	for i := 0; i < len(rucksacks); i += 3 {
		item := FindDuplicateItemInGroup(rucksacks[i : i+3])

		total += GetPriority(item)
	}

	return total
}
