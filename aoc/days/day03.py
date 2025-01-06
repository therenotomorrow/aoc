import dataclasses
import re

from aoc.library import datatypes


@dataclasses.dataclass
class Day(datatypes.Day):
    mul_pattern = re.compile(r'mul\((\d+),(\d+)\)')
    do_pattern = re.compile(r'(do\(\))|(don\'t\(\))')

    def parse(self) -> str:
        return self.indata.strip()

    def first_star(self) -> datatypes.DayResult:
        return sum(int(num1) * int(num2) for num1, num2 in self.mul_pattern.findall(self.parse()))

    def second_star(self) -> datatypes.DayResult:
        mul = self._find_mul(self.parse())
        doe = self._find_do(self.parse())

        ans = 0
        allow = True

        for idx, _ in enumerate(self.parse()):
            if doe.get(idx) is not None:
                allow = doe[idx]
                continue

            if mul.get(idx) is not None and allow:
                ans += mul[idx]

        return ans

    def _find_mul(self, memory: str) -> dict[int, int]:
        statements = {}

        for group in self.mul_pattern.finditer(memory):
            num1, num2 = map(int, group.groups())
            statements[group.start(0)] = num1 * num2

        return statements

    def _find_do(self, memory: str) -> dict[int, bool]:
        statements = {}

        for group in self.do_pattern.finditer(memory):
            statements[group.start(0)] = bool(group.group(1))

        return statements
