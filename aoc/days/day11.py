import dataclasses
import functools

from aoc.library import datatypes, funcs


@dataclasses.dataclass
class Day(datatypes.Day):
    def parse(self) -> datatypes.Ints:
        stones: datatypes.Ints = []

        for line in funcs.split(self.indata):
            stones.extend(map(int, line.split()))

        return stones

    def first_star(self) -> datatypes.DayResult:
        stones = self.parse()
        blinks = 25

        return _blink(tuple(stones), blinks)

    def second_star(self) -> datatypes.DayResult:
        stones = self.parse()
        blinks = 75

        return _blink(tuple(stones), blinks)


@functools.cache
def _transform_stone(stone: int) -> datatypes.IntsTuple:
    if stone == 0:
        return (1,)

    as_str = str(stone)
    lstr = len(as_str)
    year = 2024

    if lstr % 2 == 0:
        half = lstr // 2
        return int(as_str[:half]), int(as_str[half:])

    return (stone * year,)


@functools.cache
def _blink(stones: datatypes.IntsTuple, blink: int) -> int:
    if blink == 0:
        return len(stones)

    if len(stones) == 1:
        stones = _transform_stone(stones[0])

        return _blink(stones, blink - 1)

    return sum(_blink((stone,), blink) for stone in stones)
