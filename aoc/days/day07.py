import dataclasses

from aoc.library import datatypes, funcs

Combination = tuple[int, datatypes.Ints]


@dataclasses.dataclass
class Day(datatypes.Day):
    def parse(self) -> list[Combination]:
        calibration = []

        for line in funcs.split(self.indata):
            part1, part2 = line.split(': ')
            numbers = list(map(int, part2.split()))
            calibration.append((int(part1), numbers))

        return calibration

    def first_star(self) -> datatypes.DayResult:
        calibration = self.parse()
        total = 0

        for ans, numbers in calibration:
            if self._is_equal_first(ans, numbers[0], numbers):
                total += ans

        return total

    def second_star(self) -> datatypes.DayResult:
        calibration = self.parse()
        total = 0

        for ans, numbers in calibration:
            if self._is_equal_second(ans, numbers[0], numbers):
                total += ans

        return total

    def _is_equal_first(self, ans: int, curr: int, numbers: datatypes.Ints) -> bool:
        if len(numbers) == 1:
            return ans == curr

        other = numbers[1]

        is_sum = self._is_equal_first(ans, curr + other, numbers[1:])
        is_mul = self._is_equal_first(ans, curr * other, numbers[1:])

        return is_sum or is_mul

    def _is_equal_second(self, ans: int, curr: int, numbers: datatypes.Ints) -> bool:
        if len(numbers) == 1:
            return ans == curr

        other = numbers[1]
        another = int(str(curr) + str(other))

        is_con = self._is_equal_second(ans, another, numbers[1:])
        is_sum = self._is_equal_second(ans, curr + other, numbers[1:])
        is_mul = self._is_equal_second(ans, curr * other, numbers[1:])

        return is_con or is_sum or is_mul
