import dataclasses
import functools
import typing

from aoc.library import datatypes, funcs

_start = 'A'
_hashtag = '#'

GraphType: typing.TypeAlias = dict[datatypes.StrsTuple, str]


def _create_graph(keypad: datatypes.Grid) -> GraphType:
    graph: GraphType = {}

    for from_point in keypad.points():
        for to_point in keypad.points():
            path = ''

            path += '<' * (from_point.col - to_point.col)
            path += 'v' * (to_point.row - from_point.row)
            path += '^' * (from_point.row - to_point.row)
            path += '>' * (to_point.col - from_point.col)

            if keypad.point(from_point.row, to_point.col) == '#' or keypad.point(to_point.row, from_point.col) == '#':
                path = path[::-1]

            graph[(from_point.mark, to_point.mark)] = f'{path}A'

    return graph


@dataclasses.dataclass
class Day(datatypes.Day):
    numeric_keypad = _create_graph(
        datatypes.Grid(
            [
                ['7', '8', '9'],
                ['4', '5', '6'],
                ['1', '2', '3'],
                [_hashtag, '0', _start],
            ]
        )
    )
    direct_keypad = _create_graph(
        datatypes.Grid(
            [
                [_hashtag, '^', _start],
                ['<', 'v', '>'],
            ]
        )
    )

    def parse(self) -> datatypes.Strs:
        return list(funcs.split(self.indata))

    def first_star(self) -> datatypes.DayResult:
        robots = 2
        total_complexity = 0

        for code in self.parse():
            path = self._press_code(code)
            total_complexity += int(code[:-1]) * self._follow_robot(path, robots)

        return total_complexity

    def second_star(self) -> datatypes.DayResult:
        robots = 25
        total_complexity = 0

        for code in self.parse():
            path = self._press_code(code)
            total_complexity += int(code[:-1]) * self._follow_robot(path, robots)

        return total_complexity

    @classmethod
    @functools.cache
    def _follow_robot(cls, sequence: list[str], robot: int) -> int:
        if robot == 0:
            return len(sequence)

        prev = _start
        total_length = 0

        for char in sequence:
            total_length += cls._follow_robot(cls.direct_keypad[(prev, char)], robot - 1)
            prev = char

        return total_length

    @classmethod
    def _press_code(cls, code: str) -> str:
        prev = _start
        path = ''

        for curr in code:
            path += cls.numeric_keypad[(prev, curr)]
            prev = curr

        return path
