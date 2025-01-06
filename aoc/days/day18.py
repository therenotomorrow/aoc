import collections
import dataclasses
import heapq
import math
import typing

from aoc.library import datatypes, funcs

_hashtag = '#'

VisitedType: typing.TypeAlias = dict[datatypes.Point, float]


@dataclasses.dataclass
class Day(datatypes.Day):
    size: int
    corrupted: int
    directions = {
        datatypes.Direction.right(),
        datatypes.Direction.left(),
        datatypes.Direction.up(),
        datatypes.Direction.down(),
    }

    def parse(self) -> tuple[datatypes.Grid, list[datatypes.Point]]:
        points = []
        grid = datatypes.Grid([list('.' * self.size) for _ in range(self.size)])

        for line in funcs.split(self.indata):
            col, row = map(int, line.split(','))
            points.append(datatypes.Point(row=row, col=col, mark=_hashtag, grid=grid))

        for point in points[: self.corrupted]:
            grid[point] = _hashtag

        return grid, points

    def first_star(self) -> datatypes.DayResult:
        grid, points = self.parse()

        return self._travel(grid.point(0, 0), grid.point(self.size - 1, self.size - 1))

    def second_star(self) -> datatypes.DayResult:
        grid, points = self.parse()
        left, right = self.corrupted, len(points) - 1

        while left < right:
            mid = left + (right - left) // 2

            for idx, point in enumerate(points[left:], start=left):
                grid[point] = '.' if idx > mid else _hashtag

            ans = self._travel(grid.point(0, 0), grid.point(self.size - 1, self.size - 1))

            if ans > -1:
                left = mid + 1
            else:
                right = mid - 1

        col = points[left].col
        row = points[left].row

        return f'{col},{row}'

    def _travel(self, start: datatypes.Point, finish: datatypes.Point) -> int:  # noqa:WPS210,WPS231
        score = 0
        visited: VisitedType = collections.defaultdict(lambda: math.inf)

        queue = [(score, start)]
        visited[start] = score

        while queue:
            score, point = heapq.heappop(queue)

            if point == finish:
                return score

            for next_direct in self.directions:
                next_point = point.next(next_direct)

                if next_point.outside or next_point == _hashtag:
                    continue

                new_score = 1 + score
                old_score = visited[next_point]

                if new_score < old_score:
                    heapq.heappush(queue, (new_score, next_point))
                    visited[next_point] = new_score

        return -1
