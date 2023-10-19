package y2015

func NotQuiteLisp(stairs string, isBasement bool) int {
	floor := 0

	for pos, stair := range stairs {
		switch stair {
		case '(':
			floor++
		case ')':
			floor--
		}

		if isBasement && floor == -1 {
			return pos + 1
		}
	}

	return floor
}
