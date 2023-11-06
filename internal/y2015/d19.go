package y2015

import (
	"github.com/therenotomorrow/aoc/pkg/datastruct"
	"math/rand"
	"strings"
)

func parseReplacements(replacements []string) (keys []string, rMap map[string][]string) {
	keys = make([]string, 0)
	rMap = make(map[string][]string)

	for _, replacement := range replacements {
		parts := strings.Split(replacement, " => ")

		key := parts[0]
		val := parts[1]

		if _, ok := rMap[key]; !ok {
			keys = append(keys, key)
		}

		rMap[key] = append(rMap[key], val)
	}

	return keys, rMap
}

func generateMolecules(base string, pattern string, replacement []string) []string {
	distinct := datastruct.NewSet()

	for offset := 0; ; offset++ {
		find := strings.Index(base[offset:], pattern)

		if find == -1 {
			break
		}

		find += len(base[:offset])

		for _, sub := range replacement {
			v := base[:find] + sub

			if find+len(sub) <= len(base) {
				v += base[find+len(pattern):]
			}

			distinct.Add(v)
		}

		offset = find
	}

	vals := make([]string, distinct.Size())

	for i, val := range distinct.Values() {
		vals[i] = val.(string)
	}

	return vals
}

func Day19Part1(replacements []string, molecule string) int {
	distinct := datastruct.NewSet()
	keys, rMap := parseReplacements(replacements)

	for _, key := range keys {
		for _, newMolecule := range generateMolecules(molecule, key, rMap[key]) {
			distinct.Add(newMolecule)
		}
	}

	return distinct.Size()
}

func Day19Part2(replacements []string, molecule string) int {
	rMap := make([][]string, 0)
	init := make([]string, 0)

	for _, replacement := range replacements {
		parts := strings.Split(replacement, " => ")

		if parts[0] == "e" {
			init = append(init, parts[1])
			continue
		}

		rMap = append(rMap, parts)
	}

	target := molecule
	cnt := 0

	// thank you, Reddit <3
	for target != "e" {
		tmp := target

		for _, pairs := range rMap {
			a := pairs[0]
			b := pairs[1]

			if !strings.Contains(target, b) {
				continue
			}

			target = strings.Replace(target, b, a, 1)
			cnt++
		}

		for _, in := range init {
			if target == in {
				return cnt + 1
			}
		}

		if tmp == target {
			target = molecule
			cnt = 0
			rand.Shuffle(len(rMap), func(i, j int) {
				rMap[i], rMap[j] = rMap[j], rMap[i]
			})
		}
	}

	// impossible because of task description
	return 0
}
