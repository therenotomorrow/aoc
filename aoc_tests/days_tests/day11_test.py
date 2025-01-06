import pytest

from aoc.days import day11
from aoc.library import testit

example = """
125 17
"""
answer = """
8069 87014 98 809367 525 0 9494914 5
"""


class TestDay:
    @pytest.mark.parametrize(testit.names, [(example, 55312), (answer, 183484)], ids=testit.ids())
    def test_first_star(self, indata: str, want: int) -> None:
        testit.first_star(day11.Day(indata), want)

    @pytest.mark.parametrize(testit.names, [(example, 65601038650482), (answer, 218817038947400)], ids=testit.ids())
    def test_second_star(self, indata: str, want: int) -> None:
        testit.second_star(day11.Day(indata), want)
