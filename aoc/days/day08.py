import collections
import dataclasses
import typing

from aoc.library import datatypes, funcs

Frequencies = dict[str, list[datatypes.Point]]


@dataclasses.dataclass
class Day(datatypes.Day):
    def parse(self) -> datatypes.Grid:
        raw = []

        for line in funcs.split(self.indata):
            raw.append(list(line))

        return datatypes.Grid(raw)

    def first_star(self) -> datatypes.DayResult:
        grid = self.parse()
        anti_nodes = set()

        for antenna1, antenna2 in _yield_antennas(grid):
            point = grid.point(row=2 * antenna1.row - antenna2.row, col=2 * antenna1.col - antenna2.col)

            if point.outside:
                continue

            anti_nodes.add(point)

        return len(anti_nodes)

    def second_star(self) -> datatypes.DayResult:
        grid = self.parse()
        anti_nodes: datatypes.UniqPoints = set()

        for antenna1, antenna2 in _yield_antennas(grid):
            anti_nodes.update((antenna1, antenna2))
            diagonal = _create_diagonal(grid, antenna1, antenna2)
            anti_nodes.update(diagonal)

        return len(anti_nodes)


def _yield_antennas(grid: datatypes.Grid) -> typing.Iterable[tuple[datatypes.Point, datatypes.Point]]:
    frequencies = _collect_frequencies(grid)

    for antennas in frequencies.values():
        yield from ((antenna1, antenna2) for antenna1 in antennas for antenna2 in antennas if antenna1 != antenna2)


def _collect_frequencies(grid: datatypes.Grid) -> Frequencies:
    frequencies = collections.defaultdict(list)

    for row in grid:
        for point in row:
            if point == '.':
                continue

            frequencies[point.mark].append(point)

    return frequencies


def _create_diagonal(
    grid: datatypes.Grid,
    antenna1: datatypes.Point,
    antenna2: datatypes.Point,
) -> list[datatypes.Point]:
    diagonal = []

    diff_row = antenna1.row - antenna2.row
    diff_col = antenna1.col - antenna2.col

    point = grid.point(row=antenna1.row + diff_row, col=antenna1.col + diff_col)

    while point.inside:
        diagonal.append(point)

        point = grid.point(row=point.row + diff_row, col=point.col + diff_col)

    return diagonal
