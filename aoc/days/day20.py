import collections
import dataclasses
import typing

from aoc.library import datatypes, funcs

_max_cheat = 20

QueueType: typing.TypeAlias = tuple[int, datatypes.Point]


@dataclasses.dataclass
class Day(datatypes.Day):
    save_picoseconds: int
    directions = {
        datatypes.Direction.right(),
        datatypes.Direction.left(),
        datatypes.Direction.up(),
        datatypes.Direction.down(),
    }

    def parse(self) -> datatypes.Grid:
        raw = []

        for line in funcs.split(self.indata):
            raw.append(list(line))

        return datatypes.Grid(raw)

    def first_star(self) -> datatypes.DayResult:
        grid = self.parse()
        visited = self._travel(grid.find({'S'}))

        cheats = 0
        for point, time in visited.items():
            for direct in self.directions:
                diff = visited.get(point.next(direct.increase(2)), 0) - time
                if diff - 2 >= self.save_picoseconds:
                    cheats += 1

        return cheats

    def second_star(self) -> datatypes.DayResult:
        grid = self.parse()
        visited = self._travel(grid.find({'S'}))

        cheats = 0

        path = sorted(visited, key=lambda point: visited[point])
        for time2 in range(self.save_picoseconds, len(path)):  # noqa:WPS518
            for time1 in range(time2 - self.save_picoseconds):
                distance = path[time1].distance(path[time2])
                if distance <= _max_cheat and time2 - time1 - distance >= self.save_picoseconds:
                    cheats += 1

        return cheats

    def _travel(self, start: datatypes.Point) -> dict[datatypes.Point, int]:  # noqa:WPS210,WPS231
        time = 0
        queue = collections.deque[QueueType]()
        visited = collections.defaultdict(int)

        queue.append((time, start))
        visited[start] = time

        while queue:
            time, point = queue.popleft()
            visited[point] = time

            if point == 'E':
                break

            for next_direct in self.directions:
                next_point = point.next(next_direct)

                if next_point.outside or next_point == '#' or next_point in visited:
                    continue

                queue.append((1 + time, next_point))

        return visited
