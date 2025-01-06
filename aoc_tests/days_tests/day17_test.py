import pytest

from aoc.days import day17
from aoc.library import testit

example = """
Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0
"""
answer = """
Register A: 37293246
Register B: 0
Register C: 0

Program: 2,4,1,6,7,5,4,4,1,7,0,3,5,5,3,0
"""


class TestDay:
    @pytest.mark.parametrize(testit.names, [(example, '5,7,3,0'), (answer, '1,5,0,1,7,4,1,0,3')], ids=testit.ids())
    def test_first_star(self, indata: str, want: str) -> None:
        testit.first_star(day17.Day(indata), want)

    @pytest.mark.parametrize(testit.names, [(example, 117440), (answer, 47910079998866)], ids=testit.ids())
    def test_second_star(self, indata: str, want: int) -> None:
        testit.second_star(day17.Day(indata), want)
