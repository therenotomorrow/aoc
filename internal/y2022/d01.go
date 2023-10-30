// Package day01: https://adventofcode.com/2022/day/1
package day01

import "sort"

func SumCalories(nums []int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	return sum
}

func FindMostCalories(nums [][]int) int {
	max := nums[0][0]

	for _, cals := range nums {
		tmp := SumCalories(cals)

		if tmp > max {
			max = tmp
		}
	}

	return max
}

func FindTopCalories(nums [][]int, top int) int {
	groups := make([]int, len(nums))

	for elf, cals := range nums {
		groups[elf] = SumCalories(cals)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(groups)))

	return SumCalories(groups[0:top])
}
