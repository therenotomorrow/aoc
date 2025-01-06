import collections
import dataclasses

from aoc.library import datatypes, funcs


@dataclasses.dataclass
class Day(datatypes.Day):
    def parse(self) -> tuple[datatypes.Ints, datatypes.Ints]:
        left_list = []
        right_list = []

        for line in funcs.split(self.indata):
            left, right = map(int, line.split())

            left_list.append(left)
            right_list.append(right)

        return left_list, right_list

    def first_star(self) -> datatypes.DayResult:
        left_list, right_list = self.parse()

        left_list.sort()
        right_list.sort()

        return sum(abs(right - left) for right, left in zip(right_list, left_list))

    def second_star(self) -> datatypes.DayResult:
        left_list, right_list = self.parse()

        count = collections.Counter(right_list)

        return sum(num * count.get(num, 0) for num in left_list)
