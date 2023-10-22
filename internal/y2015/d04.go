package y2015

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"strings"
)

func Day04Part1(salt string) int {
	for num := 0; num < 1_100_000; num++ {
		numStr := strconv.Itoa(num)
		checksum := md5.Sum([]byte(salt + numStr))

		if strings.HasPrefix(hex.EncodeToString(checksum[:]), "00000") {
			return num
		}
	}

	// impossible because of task description
	return 0
}

func Day04Part2(salt string) int {
	for num := 0; num < 10_000_000; num++ {
		numStr := strconv.Itoa(num)
		checksum := md5.Sum([]byte(salt + numStr))

		if strings.HasPrefix(hex.EncodeToString(checksum[:]), "000000") {
			return num
		}
	}

	// impossible because of task description
	return 0
}
