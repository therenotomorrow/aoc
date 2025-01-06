import dataclasses

from aoc.library import datatypes, funcs


@dataclasses.dataclass
class Day(datatypes.Day):
    xmas_pattern = 'XMAS'
    mas_pattern = {'MAS', 'SAM'}
    directions = {
        datatypes.Direction.up(),
        datatypes.Direction.down(),
        datatypes.Direction.left(),
        datatypes.Direction.right(),
        datatypes.Direction.up_left(),
        datatypes.Direction.up_right(),
        datatypes.Direction.down_left(),
        datatypes.Direction.down_right(),
    }

    def parse(self) -> datatypes.Grid:
        raw = []

        for line in funcs.split(self.indata):
            raw.append(list(line))

        return datatypes.Grid(raw)

    def first_star(self) -> datatypes.DayResult:
        puzzle = self.parse()
        count = 0

        for line in puzzle:
            for point in line:
                if point != 'X':
                    continue

                count += self._find_xmas_pattern(point)

        return count

    def second_star(self) -> datatypes.DayResult:
        puzzle = self.parse()

        return sum(self._find_mas_pattern(point) for line in puzzle for point in line)

    def _find_xmas_pattern(self, start: datatypes.Point) -> int:
        count = 0

        for direct in self.directions:
            for idx, _ in enumerate(self.xmas_pattern):
                point = start.next(direct.increase(idx))

                if point.outside or point != self.xmas_pattern[idx]:
                    break
            else:
                count += 1

        return count

    def _find_mas_pattern_main_diag(self, point: datatypes.Point) -> bool:
        up_left = point.next(datatypes.Direction.up_left())
        down_right = point.next(datatypes.Direction.down_right())

        line = ''.join(map(str, (up_left, point, down_right)))

        return up_left.inside and down_right.inside and line in self.mas_pattern

    def _find_mas_pattern_sub_diag(self, point: datatypes.Point) -> bool:
        up_right = point.next(datatypes.Direction.up_right())
        down_left = point.next(datatypes.Direction.down_left())

        line = ''.join(map(str, (up_right, point, down_left)))

        return up_right.inside and down_left.inside and line in self.mas_pattern

    def _find_mas_pattern(self, start: datatypes.Point) -> bool:
        return self._find_mas_pattern_main_diag(start) and self._find_mas_pattern_sub_diag(start)
