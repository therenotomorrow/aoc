import collections
import dataclasses

from aoc.library import datatypes, funcs

_hashtag = '#'
_open_bracket = '['
_big_o = 'O'


@dataclasses.dataclass
class Day(datatypes.Day):
    def first_star(self) -> datatypes.DayResult:
        parts = funcs.parts(self.indata)

        grid = _parse_grid(parts[0])
        dirs = _parse_dirs(parts[1])
        start = grid.find({'@'})

        grid = _travel(grid, start, dirs)

        return _calc_gps(grid, _big_o)

    def second_star(self) -> datatypes.DayResult:
        parts = funcs.parts(self.indata)

        grid = _parse_wide(parts[0])
        dirs = _parse_dirs(parts[1])
        start = grid.find({'@'})

        for direct in dirs:
            if direct.name in {'<', '>'}:
                start = _move_horizontal(grid, start, direct)
            else:
                start = _move_vertical(grid, start, direct)

        return _calc_gps(grid, _open_bracket)


def _travel(grid: datatypes.Grid, point: datatypes.Point, directions: datatypes.Directions) -> datatypes.Grid:
    for direct in directions:
        next_point = point.next(direct)

        match next_point.mark:
            case '.':
                could_move = True
            case 'O':
                next_point, could_move = _find_next_free_point(grid, next_point, direct)
            case _:
                could_move = False

        if not could_move:
            continue

        grid[point] = next_point
        grid[next_point] = point
        point = next_point

    return grid


def _move_vertical(
    grid: datatypes.Grid,
    point: datatypes.Point,
    direct: datatypes.Direction,
) -> datatypes.Point:
    rows = _is_possible_move(point, direct)
    if not rows:
        return point

    diff = 1 if direct == '^' else -1
    reverse = direct == 'v'

    for point_row in sorted(rows, reverse=reverse):
        for point_col in rows[point_row]:
            curr_point = grid.point(point_row, point_col)
            prev_point = curr_point.prev(direct)

            if point_col in rows[point_row + diff]:
                grid[curr_point] = prev_point
            else:
                grid[curr_point] = '.'

    return point.next(direct)


def _calc_gps(grid: datatypes.Grid, mark: str) -> int:
    gps = 0

    for row in grid:
        for point in row:
            if point == mark:
                gps += 100 * point.row
                gps += point.col

    return gps


def _parse_grid(indata: str) -> datatypes.Grid:
    raw = []

    for line in funcs.split(indata):
        raw.append(list(line))

    return datatypes.Grid(raw)


def _parse_wide(indata: str) -> datatypes.Grid:
    raw = []

    for line in funcs.split(indata):
        lines: datatypes.Strs = []

        for mark in line:
            if mark in {'.', _hashtag}:
                lines.extend(mark * 2)
            elif mark == _big_o:
                lines.extend((_open_bracket, ']'))
            else:
                lines.extend(('@', '.'))

        raw.append(lines)

    return datatypes.Grid(raw)


def _parse_dirs(indata: str) -> datatypes.Directions:
    moves = []

    for line in funcs.split(indata):
        for name in line:
            moves.append(datatypes.Direction.new(name))

    return moves


def _find_next_free_point(
    grid: datatypes.Grid,
    point: datatypes.Point,
    direct: datatypes.Direction,
) -> tuple[datatypes.Point, bool]:
    box_point = point

    while box_point == _big_o:
        box_point = box_point.next(direct)

    found = box_point != _hashtag
    if found:
        grid[point] = box_point
        grid[box_point] = point

    return grid[point], found


def _move_horizontal(
    grid: datatypes.Grid,
    point: datatypes.Point,
    direct: datatypes.Direction,
) -> datatypes.Point:
    next_point = point.next(direct)

    while next_point.mark in {_open_bracket, ']'}:
        next_point = next_point.next(direct)

    if next_point == _hashtag:
        return point

    prev_point = next_point
    while next_point != point:
        prev_point = prev_point.prev(direct)
        grid[next_point] = prev_point
        next_point = next_point.prev(direct)

    return point.next(direct)


def _is_possible_move(  # noqa:WPS231
    point: datatypes.Point,
    direct: datatypes.Direction,
) -> dict[int, datatypes.IntsSet]:
    queue = collections.deque[datatypes.Point]()
    rows: dict[int, datatypes.IntsSet] = collections.defaultdict(set)

    queue.append(point.next(direct))

    while queue:
        curr: datatypes.Point = queue.pop()

        match curr.mark:
            case '#':
                return {}

            case ']':
                rows[curr.row] |= {curr.col - 1, curr.col}
                diag = datatypes.Direction.down_left() if direct == 'v' else datatypes.Direction.up_left()
                queue.extend((curr.next(direct), curr.next(diag)))

            case '[':
                rows[curr.row] |= {curr.col, curr.col + 1}
                diag = datatypes.Direction.up_right() if direct == '^' else datatypes.Direction.down_right()
                queue.extend((curr.next(direct), curr.next(diag)))

            case '.':
                rows[curr.row].add(curr.col)

    return rows
