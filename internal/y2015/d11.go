package y2015

func isSecure(password []rune) bool {
	isIncrease := false

	for i := range password[:len(password)-2] {
		if password[i+2]-password[i+1] == 1 && password[i+1]-password[i] == 1 {
			isIncrease = true
			break
		}
	}

	if !isIncrease {
		return false
	}

	for _, r := range password {
		if r == 'i' || r == 'o' || r == 'l' {
			return false
		}
	}

	cntDoubles := 0
	prev := rune(0)

	for i := range password[:len(password)-1] {
		if password[i+1] != password[i] {
			continue
		}

		if prev != password[i] {
			prev = password[i]
			cntDoubles++
		}
	}

	return cntDoubles > 1
}

func genNextPassword(password []rune) []rune {
	var reminder int32 = 1

	for i := len(password) - 1; i >= 0; i-- {
		password[i] += reminder

		if password[i] <= 'z' {
			break
		}

		password[i] = 'a'
	}

	return password
}

func Day11Part1(prev string) string {
	curr := []rune(prev)

	for !isSecure(curr) {
		curr = genNextPassword(curr)
	}

	return string(curr)
}

func Day11Part2(prev string) string {
	curr := []rune(prev)

	// skip the secure password
	curr = genNextPassword(curr)

	for !isSecure(curr) {
		curr = genNextPassword(curr)
	}

	return string(curr)
}
