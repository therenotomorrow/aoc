package y2015

import (
	"math"
	"strconv"
	"strings"
)

func Day14Part1(deer []string, time int) int {
	maxDistance := math.MinInt

	for _, d := range deer {
		parts := strings.Fields(d)

		speed, _ := strconv.Atoi(parts[3])
		flyTime, _ := strconv.Atoi(parts[6])
		restTime, _ := strconv.Atoi(parts[13])

		circles := time / (flyTime + restTime)
		distance := speed * flyTime * circles
		left := time % (flyTime + restTime)

		if left < flyTime {
			distance += speed * left
		} else {
			distance += speed * flyTime
		}

		if distance > maxDistance {
			maxDistance = distance
		}
	}

	return maxDistance
}

func Day14Part2(deer []string, time int) int {
	table := make(map[int][]int)
	distances := make([]int, len(deer))
	scores := make([]int, len(deer))
	times := make([]int, len(deer))
	states := make([]bool, len(deer))

	for i, d := range deer {
		parts := strings.Fields(d)

		speed, _ := strconv.Atoi(parts[3])
		flyTime, _ := strconv.Atoi(parts[6])
		restTime, _ := strconv.Atoi(parts[13])

		times[i] = flyTime
		table[i] = []int{speed, flyTime, restTime}
	}

	for ; time > 0; time-- {
		maxDistance := math.MinInt

		for i, stats := range table {
			times[i]--

			if !states[i] {
				distances[i] += stats[0]

				if times[i] == 0 {
					states[i] = !states[i]
					times[i] = stats[2]
				}
			} else {
				if times[i] == 0 {
					states[i] = !states[i]
					times[i] = stats[1]
				}
			}
		}

		for _, distance := range distances {
			if distance > maxDistance {
				maxDistance = distance
			}
		}

		for i, distance := range distances {
			if distance == maxDistance {
				scores[i]++
			}
		}
	}

	maxScore := math.MinInt
	for _, score := range scores {
		if score > maxScore {
			maxScore = score
		}
	}

	return maxScore
}
