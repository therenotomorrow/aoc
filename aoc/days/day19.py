import dataclasses
import functools

from aoc.library import datatypes, funcs


@dataclasses.dataclass
class Day(datatypes.Day):
    def parse(self) -> tuple[datatypes.Strs, datatypes.Strs]:
        parts = funcs.parts(self.indata)

        towels = _parse_towels(parts[0])
        patterns = _parse_patterns(parts[1])

        return towels, patterns

    def first_star(self) -> datatypes.DayResult:
        towels, patterns = self.parse()
        ans = 0

        for pattern in patterns:
            if _count_possible(pattern, tuple(towels)):
                ans += 1

        return ans

    def second_star(self) -> datatypes.DayResult:
        towels, patterns = self.parse()

        return sum(_count_possible(pattern, tuple(towels)) for pattern in patterns)


@functools.cache
def _count_possible(pattern: str, towels: datatypes.StrsTuple) -> int:
    if not pattern:
        return 1

    combinations = 0
    for towel in towels:
        if not pattern.startswith(towel):
            continue

        idx = len(towel)
        combinations += _count_possible(pattern[idx:], towels)

    return combinations


def _parse_towels(indata: str) -> datatypes.Strs:
    return indata.strip().split(', ')


def _parse_patterns(indata: str) -> datatypes.Strs:
    return list(funcs.split(indata))
