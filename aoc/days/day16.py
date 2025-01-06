import collections
import dataclasses
import heapq
import math
import typing

from aoc.library import datatypes, funcs

VisitedType: typing.TypeAlias = dict[tuple[datatypes.Point, datatypes.Direction], tuple[float, datatypes.UniqPoints]]


@dataclasses.dataclass
class Day(datatypes.Day):
    directions = {
        datatypes.Direction.right(),
        datatypes.Direction.left(),
        datatypes.Direction.up(),
        datatypes.Direction.down(),
    }

    def parse(self) -> tuple[datatypes.Grid, datatypes.Point]:
        raw = []

        for line in funcs.split(self.indata):
            raw.append(list(line))

        grid = datatypes.Grid(raw)
        start = grid.find({'S'})

        return grid, start

    def first_star(self) -> datatypes.DayResult:
        grid, start = self.parse()
        direct = datatypes.Direction.right()

        return self._travel(start, direct)[0]

    def second_star(self) -> datatypes.DayResult:
        grid, start = self.parse()
        direct = datatypes.Direction.right()

        return self._travel(start, direct)[1]

    def _travel(self, point: datatypes.Point, direct: datatypes.Direction) -> datatypes.IntsTuple:  # noqa:WPS210,WPS231
        score = 0
        visited: VisitedType = collections.defaultdict(lambda: (math.inf, set()))

        queue = [(score, point, direct)]
        visited[(point, direct)] = (score, {point})

        while queue:
            score, point, direct = heapq.heappop(queue)

            if point == 'E':
                break

            for next_direct in self.directions:
                next_point = point.next(next_direct)

                if next_point.outside or next_point == '#':
                    continue

                add_score = 0 if direct == next_direct else 1000
                new_score = 1 + score + add_score
                old_score = visited[(next_point, next_direct)][0]

                new_visited_cells = visited[(point, direct)][1] | {next_point}

                if new_score < old_score:
                    heapq.heappush(queue, (new_score, next_point, next_direct))
                elif new_score == old_score:
                    new_visited_cells |= visited[(next_point, next_direct)][1]
                else:
                    continue

                visited[(next_point, next_direct)] = (new_score, new_visited_cells)

        return score, len(visited[(point, direct)][1])
