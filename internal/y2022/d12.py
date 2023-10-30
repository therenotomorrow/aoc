""" Module day12: https://adventofcode.com/2022/day/12"""
import dataclasses
import itertools
import pprint
import typing
import functools


class CoordinateError(Exception):
    pass


@dataclasses.dataclass(repr=False)
class Coordinate(object):
    height: str
    x: int
    y: int
    visited: bool
    cost: int = 1

    def __repr__(self) -> str:
        return f'({self.y}, {self.x})[{self.height}]'


class HeightMap(object):
    def __init__(self, map_: list[list[Coordinate]]):
        self._map = map_
        self.start = self._find_coord('S')
        self.end = self._find_coord('E')
        self.end.height = '{'
        self.found = False

    def _find_coord(self, height: typing.Literal['S', 'E']) -> Coordinate:
        for row in self._map:
            for coord in row:
                if coord.height == height:
                    return coord

        raise CoordinateError(f'cannot find {height=}')

    def _next_steps(self, curr: Coordinate) -> list[Coordinate]:
        can = []

        if (x := curr.x - 1) > -1:
            can.append(self._map[curr.y][x])

        if (x := curr.x + 1) < len(self._map[0]):
            can.append(self._map[curr.y][x])

        if (y := curr.y - 1) > -1:
            can.append(self._map[y][curr.x])

        if (y := curr.y + 1) < len(self._map):
            can.append(self._map[y][curr.x])

        if curr.height == 'S':
            return can

        return [c for c in can if 0 <= ord(c.height) - ord(curr.height) <= 1]

    def point_at(self, coordinate: Coordinate):
        return self._map[coordinate.y][coordinate.x]

    def walk(self, queue: list[tuple[int, Coordinate, Coordinate]] = None) -> None:
        if not queue:
            queue = [(0, self.start, self.start)]
        cycles = 0
        while queue:
            cycles += 1
            curr_cost, curr_point, prev_point = queue.pop(0)
            if not self.point_at(curr_point).visited:
                self.point_at(curr_point).cost = curr_cost
                self.point_at(curr_point).visited = True
                self.point_at(curr_point).previous_point = prev_point
                if curr_point == self.end:
                    break
                else:
                    curr_cost += 1
                    for next_point in self._next_steps(curr_point):
                        queue.append((curr_cost, next_point, curr_point))
                    queue.sort(key=lambda l: l[0])
            if len(queue) == 1:
                print('asd')
        print(f"Completed in {cycles} cycles")

    #
    # def find_the_way(self) -> int:
    #     curr = self.start
    #     steps = [curr]
    #
    #     while not self.found:
    #         curr = steps.pop(0)
    #
    #         if curr.visited:
    #             continue
    #
    #         if curr.height == 'E':
    #             self.found = True
    #             return curr.cost
    #
    #         curr.visited = True
    #
    #         for nex in self._next_steps(curr):
    #             steps.append((val+1, nex))
    #
    #     return val
    #


def create_map(height_map: list[list[str]]) -> list[list[int]]:
    vertical = []

    for row in height_map:
        horizontal = []

        for mountain in row:
            match mountain:
                case 'S':
                    height = 0
                case 'E':
                    height = ord('z') - ord('a') + 2
                case _:
                    height = ord(mountain) - ord('a') + 1

            horizontal.append(height)

        vertical.append(horizontal)

    return vertical


def find_start(height_map: list[list[int]]) -> (int, int):
    for i, row in enumerate(height_map):
        for j, col in enumerate(row):
            if hm[i][j] == 0:
                return (i, j)


def create_visited(height_map: list[list[str]]) -> list[list[int]]:
    vertical = []

    for row in height_map:
        vertical.append([False for _ in row])

    return vertical


def create_weight(height_map: list[list[str]]) -> list[list[int]]:
    vertical = []

    for row in height_map:
        vertical.append([0 for _ in row])

    return vertical


if __name__ == '__main__':
    part = True
    lines = []

    with open('input.txt') as f:
        for i, line in enumerate(f):
            lines.append([i for i in line.strip()])

    hm = create_map(lines)
    visited = create_visited(lines)
    costs = create_weight(lines)

    queue = []
    if not part:
        queue = [(0, find_start(hm))]
    else:
        for i, row in enumerate(hm):
            for j, col in enumerate(row):
                if hm[i][j] == 1:
                    queue.append((0, (i, j)))

    while queue:
        cost, (r, c) = queue.pop(0)

        if visited[r][c]:
            continue

        visited[r][c] = True
        costs[r][c] = cost

        if hm[r][c] == 27:
            print(costs[r][c])
            break

        cost += 1

        points = [
            (y, x)
            for y, x in ((r - 1, c), (r + 1, c), (r, c - 1), (r, c + 1))
            if 0 <= y < len(hm) and 0 <= x < len(hm[0])
        ]
        for y, x in points:
            if hm[y][x] <= hm[r][c] + 1:
                queue.append((cost, (y, x)))

        queue.sort(key=lambda x: x[0])
