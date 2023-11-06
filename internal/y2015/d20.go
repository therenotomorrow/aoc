package y2015

import "math"

func Day20Part1(n int) int {
	ans := 0

	for i := 1; i <= math.MaxInt; i++ {
		sum := 0

		// thank you, Reddit <3
		for j := 1; j <= int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				sum += (j + i/j) * 10
			}
		}

		if sum >= n {
			ans = i
			break
		}
	}

	return ans
}

func Day20Part2(n int) int {
	ans := 0

	for i := 1; i <= math.MaxInt; i++ {
		sum := 0
		count := 0

		// thank you, Reddit <3
		for j := 1; j <= int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				sum += (j + i/j) * 11
			}

			count++

			if count == 50 {
				break
			}
		}

		if sum >= n {
			ans = i
			break
		}
	}

	return ans
}
