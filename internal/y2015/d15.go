package y2015

import (
	"math"
	"strconv"
	"strings"
)

const teaspoons = 100

func parseIngredients(ingredients []string) [][]int {
	props := make([][]int, len(ingredients))

	for i, ing := range ingredients {
		parts := strings.Fields(ing)

		capacity, _ := strconv.Atoi(strings.TrimSuffix(parts[2], ","))
		durability, _ := strconv.Atoi(strings.TrimSuffix(parts[4], ","))
		flavor, _ := strconv.Atoi(strings.TrimSuffix(parts[6], ","))
		texture, _ := strconv.Atoi(strings.TrimSuffix(parts[8], ","))
		calories, _ := strconv.Atoi(parts[10])

		props[i] = []int{capacity, durability, flavor, texture, calories}
	}

	return props
}

func genReceipts(nCookies int, remain int) [][]int {
	if nCookies == 1 {
		return [][]int{{remain}}
	}

	var result [][]int

	for i := 0; i <= remain+1; i++ {
		for _, sub := range genReceipts(nCookies-1, remain-i) {
			result = append(result, append([]int{i}, sub...))
		}
	}

	return result
}

func calcReceipt(receipt []int, props [][]int, calGuard int) int {
	score := 1

	if calGuard != 0 {
		cal := 0

		for j, r := range receipt {
			cal += r * props[j][len(props[0])-1]
		}

		if cal != calGuard {
			return 0
		}
	}

	for i := 0; i < len(props[0])-1; i++ {
		sum := 0

		for j, r := range receipt {
			sum += r * props[j][i]
		}

		score *= sum

		if score < 0 {
			return 0
		}
	}

	return score
}

func Day15Part1(ingredients []string) int {
	props := parseIngredients(ingredients)
	maxScore := math.MinInt

	for _, receipt := range genReceipts(len(props), teaspoons) {
		score := calcReceipt(receipt, props, 0)

		if score > maxScore {
			maxScore = score
		}
	}

	return maxScore
}

func Day15Part2(ingredients []string) int {
	props := parseIngredients(ingredients)
	maxScore := math.MinInt

	for _, receipt := range genReceipts(len(props), teaspoons) {
		score := calcReceipt(receipt, props, 500)

		if score > maxScore {
			maxScore = score
		}
	}

	return maxScore
}
