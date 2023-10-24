package y2015

import "github.com/therenotomorrow/aoc/pkg/datastruct"

var vowels = datastruct.NewSet()

func isNicePart1(s string) bool {
	var cntVowels, cntTwice, cntDoubles int

	for i := 0; i < len(s); i++ {
		if vowels.Contains(s[i]) {
			cntVowels++
		}

		// index error for `cntTwice` and `cntDoubles`
		if i > len(s)-2 {
			break
		}

		if s[i] == s[i+1] {
			cntTwice++
		}

		switch s[i : i+2] {
		case "ab", "cd", "pq", "xy":
			cntDoubles++
		}
	}

	return cntVowels > 2 && cntTwice > 0 && cntDoubles < 1
}

func Day05Part1(strings []string) int {
	cnt := 0

	vowels.Populate(byte('a'), byte('e'), byte('i'), byte('o'), byte('u'))

	for _, str := range strings {
		if isNicePart1(str) {
			cnt++
		}
	}

	return cnt
}

func isNicePart2(s string) bool {
	isMiddle := false
	isPair := false

	used := make(map[string]int)
	for i := 0; i < len(s)-1; i++ {
		if j, ok := used[s[i:i+2]]; ok {
			if i-j > 1 {
				isPair = true
			}
		}

		used[s[i:i+2]] = i

		// index error for `isMiddle`
		if i >= len(s)-2 {
			break
		}

		if s[i] == s[i+2] {
			isMiddle = true
		}
	}

	return isPair && isMiddle
}

func Day05Part2(strings []string) int {
	cnt := 0

	for _, str := range strings {
		if isNicePart2(str) {
			cnt++
		}
	}

	return cnt
}
