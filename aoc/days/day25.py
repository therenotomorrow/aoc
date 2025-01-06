import dataclasses

from aoc.library import datatypes, funcs

_overlap_sum = 5


@dataclasses.dataclass
class Day(datatypes.Day):
    def parse(self) -> tuple[list[datatypes.Ints], list[datatypes.Ints]]:
        locks: list[datatypes.Ints] = []
        keys: list[datatypes.Ints] = []

        for part in funcs.parts(self.indata, num=0):
            rows = list(funcs.split(part))
            heights = [column.count('#') - 1 for column in zip(*rows)]

            if rows[0] == '#####':
                keys.append(heights)
            else:
                locks.append(heights)

        return locks, keys

    def first_star(self) -> datatypes.DayResult:
        locks, keys = self.parse()
        ans = 0

        for key in keys:
            for lock in locks:
                if all(lock_val + key_val <= _overlap_sum for lock_val, key_val in zip(lock, key)):
                    ans += 1

        return ans

    def second_star(self) -> datatypes.DayResult:
        return 'Happy New Year!'
