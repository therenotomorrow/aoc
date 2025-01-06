import dataclasses
import typing

from aoc.library import datatypes, funcs


@dataclasses.dataclass
class Day(datatypes.Day):
    directions = {
        datatypes.Direction.up(),
        datatypes.Direction.down(),
        datatypes.Direction.left(),
        datatypes.Direction.right(),
    }
    _checks = {
        (datatypes.Direction.up(), datatypes.Direction.left()),
        (datatypes.Direction.right(), datatypes.Direction.up()),
        (datatypes.Direction.down(), datatypes.Direction.right()),
        (datatypes.Direction.left(), datatypes.Direction.down()),
    }

    def parse(self) -> datatypes.Grid:
        raw = []

        for line in funcs.split(self.indata):
            raw.append(list(line))

        return datatypes.Grid(raw)

    def first_star(self) -> datatypes.DayResult:
        ans = 0

        for region in self._collect_regions(self.parse()):
            perimeter = 0

            for point in region:
                for direct in self.directions:
                    if point.next(direct) not in region:
                        perimeter += 1

            ans += len(region) * perimeter

        return ans

    def second_star(self) -> datatypes.DayResult:  # noqa:WPS210,WPS231
        grid = self.parse()
        ans = 0

        for region in self._collect_regions(grid):
            num_sides = 0

            for point in region:
                for direct in self._checks:
                    point_a = point.next(direct[0])
                    point_b = point.next(direct[1])
                    point_ab = point_a.next(direct[1])

                    a_ok = point_a in region
                    b_ok = point_b in region
                    ab_ok = point_ab in region

                    both_outside = not a_ok and not b_ok
                    forms_corner = a_ok and b_ok and not ab_ok

                    if both_outside or forms_corner:
                        num_sides += 1

            ans += len(region) * num_sides

        return ans

    def _travel(self, point: datatypes.Point) -> typing.Iterable[datatypes.Point]:
        for direct in self.directions:
            next_point = point.next(direct)

            if next_point.outside:
                continue

            yield next_point

    def _collect_region(self, visited: datatypes.UniqPoints, start: datatypes.Point) -> datatypes.UniqPoints:
        if start in visited:
            return set()

        visited.add(start)
        region = {start}

        for next_point in self._travel(start):
            if next_point.mark != start.mark:
                continue

            region.update(self._collect_region(visited, next_point))

        return region

    def _collect_regions(self, grid: datatypes.Grid) -> typing.Iterable[datatypes.UniqPoints]:
        visited: datatypes.UniqPoints = set()

        for row in grid:
            for point in row:
                if point in visited:
                    continue

                region = self._collect_region(visited, point)
                if region:
                    yield region
