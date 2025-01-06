import dataclasses
import typing

from aoc.library import datatypes, funcs


@dataclasses.dataclass(slots=True, kw_only=True, frozen=True)
class Position:
    row: int
    col: int


@dataclasses.dataclass(slots=True, kw_only=True, frozen=True)
class Velocity:
    row: int
    col: int


@dataclasses.dataclass(slots=True, kw_only=True, frozen=True)
class Robot:
    pos: Position
    vel: Velocity


Tiles: typing.TypeAlias = list[datatypes.Ints]


@dataclasses.dataclass
class Day(datatypes.Day):
    rows: int
    cols: int

    def parse(self) -> list[Robot]:
        robots = []

        for line in funcs.split(self.indata):
            parts: datatypes.Strs = line.split()

            col, row = map(int, parts[0].lstrip('p=').split(','))
            v_col, v_row = map(int, parts[1].lstrip('v=').split(','))

            robots.append(Robot(pos=Position(row=row, col=col), vel=Velocity(row=v_row, col=v_col)))

        return robots

    def first_star(self) -> datatypes.DayResult:
        robots = self.parse()
        seconds = 100

        return self._run(robots, seconds)

    def second_star(self) -> datatypes.DayResult:
        robots = self.parse()
        seconds = 10000
        tree_fact = 69552000  # empirically

        while seconds > 0:
            fact = self._run(robots, seconds)
            if fact == tree_fact:
                return seconds

            seconds -= 1

        return seconds

    def _run(self, robots: list[Robot], seconds: int) -> int:
        tiles = [[0 for _ in range(self.cols)] for _ in range(self.rows)]

        for robot in robots:
            tiles = self._move(tiles, robot, seconds)

        return self._calc_safety_factor(tiles)

    def _move(self, tiles: Tiles, robot: Robot, seconds: int) -> Tiles:
        new_row = robot.pos.row + seconds * robot.vel.row
        new_col = robot.pos.col + seconds * robot.vel.col

        new_row = new_row % self.rows
        new_col = new_col % self.cols

        tiles[new_row][new_col] += 1

        return tiles

    def _calc_first_squad(self, tiles: Tiles) -> int:
        return sum(tiles[row][col] for col in range(self.cols // 2) for row in range(self.rows // 2))

    def _calc_second_squad(self, tiles: Tiles) -> int:
        return sum(tiles[row][col] for col in range(self.cols // 2) for row in range(self.rows // 2 + 1, len(tiles)))

    def _calc_third_squad(self, tiles: Tiles) -> int:
        return sum(tiles[row][col] for col in range(self.cols // 2 + 1, len(tiles[0])) for row in range(self.rows // 2))

    def _calc_fourth_squad(self, tiles: Tiles) -> int:
        return sum(
            tiles[row][col]
            for col in range(self.cols // 2 + 1, len(tiles[0]))
            for row in range(self.rows // 2 + 1, len(tiles))
        )

    def _calc_safety_factor(self, tiles: Tiles) -> int:
        robots = {
            self._calc_first_squad(tiles),
            self._calc_second_squad(tiles),
            self._calc_third_squad(tiles),
            self._calc_fourth_squad(tiles),
        }

        mul = 1
        for num in robots:
            mul *= num

        return mul
