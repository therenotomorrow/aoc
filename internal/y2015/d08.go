package y2015

import "strconv"

func Day08Part1(array []string) int {
	numSym := 0
	numMem := 0

	for _, s := range array {
		numSym += len(s)
		s, _ = strconv.Unquote(s)
		numMem += len(s)
	}

	return numSym - numMem
}

func Day08Part2(array []string) int {
	numSym := 0
	numEnc := 0

	for _, s := range array {
		numSym += len(s)
		s = strconv.Quote(s)
		numEnc += len(s)
	}

	return numEnc - numSym
}
