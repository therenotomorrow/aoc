package y2015

import "math"

func Day17Part1(target int, containers []int) int {
	cnt := 0
	// technic here: we use bits as took container or not
	choices := 1 << len(containers)

	for ; choices > 0; choices-- {
		choice := choices
		sum := 0

		for _, container := range containers {
			if choice&1 == 1 {
				sum += container
			}

			choice >>= 1

			if choice == 0 {
				break
			}
		}

		if sum == target {
			cnt++
		}
	}

	return cnt
}

func Day17Part2(target int, containers []int) int {
	minCnt := 0
	currMin := math.MaxInt
	// technic here: we use bits as took container or not
	choices := 1 << len(containers)

	for ; choices > 0; choices-- {
		choice := choices
		sum := 0
		currCnt := 0

		for _, container := range containers {
			if choice&1 == 1 {
				sum += container
				currCnt++
			}

			choice >>= 1

			if choice == 0 {
				break
			}
		}

		if sum != target {
			continue
		}

		if currCnt < currMin {
			currMin = currCnt
			minCnt = 0
		}

		if currCnt == currMin {
			minCnt++
		}
	}

	return minCnt
}
