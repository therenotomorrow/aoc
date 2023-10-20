package y2015

func Day01Part1(stairs string) int {
	floor := 0

	for _, stair := range stairs {
		switch stair {
		case '(':
			floor++
		case ')':
			floor--
		}
	}

	return floor
}

func Day01Part2(stairs string) int {
	floor := 0

	for pos, stair := range stairs {
		switch stair {
		case '(':
			floor++
		case ')':
			floor--
		}

		if floor == -1 {
			return pos + 1
		}
	}

	// impossible because of task description
	return 0
}
