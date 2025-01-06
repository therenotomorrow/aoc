import dataclasses

from aoc.library import datatypes, funcs


@dataclasses.dataclass
class Day(datatypes.Day):
    def parse(self) -> tuple[datatypes.Grid, datatypes.Point]:
        raw = []

        for line in funcs.split(self.indata):
            raw.append(list(line))

        grid = datatypes.Grid(raw)
        start = grid.find({'<', '>', '^', 'v'})

        return grid, start

    def first_star(self) -> datatypes.DayResult:
        grid, start = self.parse()
        direction = datatypes.Direction.new(start.mark)

        return len(_travel(start, direction))

    def second_star(self) -> datatypes.DayResult:
        grid, start = self.parse()
        direction = datatypes.Direction.new(start.mark)

        ans = 0

        for point in _travel(start, direction):
            if start == point:
                continue

            grid[point] = '#'
            ans += _detect_cycle(start, direction, grid.size)
            grid[point] = '.'

        return ans


def _detect_cycle(point: datatypes.Point, direct: datatypes.Direction, tries: int) -> bool:
    while tries > 0:
        tries -= 1

        next_point = point.next(direct)

        if next_point.outside:
            break

        if next_point == '#':
            direct = direct.turn_right
        else:
            point = next_point

    return tries == 0


def _travel(point: datatypes.Point, direct: datatypes.Direction) -> datatypes.UniqPoints:
    visited = set()

    while True:
        next_point = point.next(direct)

        if next_point.outside:
            break

        if next_point == '#':
            direct = direct.turn_right
            continue

        point = next_point
        visited.add(point)

    return visited
