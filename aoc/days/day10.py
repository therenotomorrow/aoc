import dataclasses

from aoc.library import datatypes, funcs


@dataclasses.dataclass
class Day(datatypes.Day):
    directions = {
        datatypes.Direction.up(),
        datatypes.Direction.down(),
        datatypes.Direction.left(),
        datatypes.Direction.right(),
    }

    def parse(self) -> datatypes.Grid:
        raw = []

        for line in funcs.split(self.indata):
            raw.append(list(line))

        return datatypes.Grid(raw)

    def first_star(self) -> datatypes.DayResult:
        grid = self.parse()
        ans = 0

        for line in grid:
            for point in line:
                if point != '0':
                    continue

                ans += len(self._travel_map(point))

        return ans

    def second_star(self) -> datatypes.DayResult:
        grid = self.parse()
        ans = 0

        for line in grid:
            for point in line:
                if point != '0':
                    continue

                ans += self._count_map(point)

        return ans

    def _find_next_moves(self, point: datatypes.Point) -> list[datatypes.Point]:
        points = []

        for direct in self.directions:
            next_point = point.next(direct)

            if next_point.outside:
                continue

            if _can_move(point, next_point):
                points.append(next_point)

        return points

    def _travel_map(self, point: datatypes.Point) -> datatypes.UniqPoints:
        if point == '9':
            return {point}

        ans = set()

        for next_point in self._find_next_moves(point):
            ans.update(self._travel_map(next_point))

        return ans

    def _count_map(self, point: datatypes.Point) -> int:
        if point == '9':
            return 1

        return sum(self._count_map(next_point) for next_point in self._find_next_moves(point))


def _can_move(from_point: datatypes.Point, to_point: datatypes.Point) -> bool:
    return int(to_point.mark) - int(from_point.mark) == 1
