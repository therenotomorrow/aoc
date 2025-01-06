import dataclasses
import typing

Ints: typing.TypeAlias = list[int]
Strs: typing.TypeAlias = list[str]
IntsTuple: typing.TypeAlias = tuple[int, ...]
StrsTuple: typing.TypeAlias = tuple[str, ...]
IntsSet: typing.TypeAlias = set[int]
StrsSet: typing.TypeAlias = set[str]


@dataclasses.dataclass(slots=True, kw_only=True, frozen=True)
class Point:
    row: int
    col: int
    mark: str
    grid: 'Grid'

    def __eq__(self, other: object) -> bool:
        if isinstance(other, str):
            return self.mark == other

        if not isinstance(other, Point):
            return NotImplemented  # pragma: no cover

        return (self.row, self.col) == (other.row, other.col)

    def __lt__(self, other: object) -> bool:
        if not isinstance(other, Point):
            return NotImplemented  # pragma: no cover

        return (self.row, self.col) < (other.row, other.col)

    def __hash__(self) -> int:
        return hash((self.row, self.col))

    def __str__(self) -> str:
        return self.mark

    @property
    def inside(self) -> bool:
        return 0 <= self.row < self.grid.rows and 0 <= self.col < self.grid.cols

    @property
    def outside(self) -> bool:
        return not self.inside

    def next(self, direct: 'Direction') -> 'Point':
        next_point = Point(row=self.row + direct.row, col=self.col + direct.col, mark='', grid=self.grid)

        return self.grid[next_point] if next_point.inside else next_point

    def prev(self, direct: 'Direction') -> 'Point':
        return self.next(direct.opposite)

    def distance(self, other: 'Point') -> int:
        return abs(self.row - other.row) + abs(self.col - other.col)


_up = '^'
_down = 'v'
_left = '<'
_right = '>'


@dataclasses.dataclass(slots=True, kw_only=True, frozen=True)
class Direction:  # noqa:WPS214
    row: int = 0
    col: int = 0
    name: str = ''

    def __eq__(self, other: object) -> bool:
        if isinstance(other, str):
            return self.name == other

        if not isinstance(other, Direction):
            return NotImplemented  # pragma: no cover

        return (self.row, self.col, self.name) == (other.row, other.col, other.name)

    def __lt__(self, other: object) -> bool:
        if not isinstance(other, Direction):
            return NotImplemented  # pragma: no cover

        return self.name < other.name

    def __hash__(self) -> int:
        return hash((self.row, self.col, self.name))

    def increase(self, mul: int) -> 'Direction':
        mul_row = mul * self.row
        mul_col = mul * self.col

        return Direction(row=mul_row, col=mul_col, name=self.name)

    @classmethod
    def new(cls, name: str) -> 'Direction':
        if name == _up:
            return cls.up()
        if name == _down:
            return cls.down()
        if name == _left:
            return cls.left()

        return cls.right()

    @property
    def turn_right(self) -> 'Direction':
        if self.name == _up:
            return self.right()
        if self.name == _right:
            return self.down()
        if self.name == _down:
            return self.left()

        return self.up()

    @property
    def opposite(self) -> 'Direction':
        if self.name == _up:
            return self.down()
        if self.name == _right:
            return self.left()
        if self.name == _down:
            return self.up()

        return self.right()

    @classmethod
    def up(cls) -> 'Direction':
        return cls(row=-1, name=_up)

    @classmethod
    def down(cls) -> 'Direction':
        return cls(row=1, name=_down)

    @classmethod
    def left(cls) -> 'Direction':
        return cls(col=-1, name=_left)

    @classmethod
    def right(cls) -> 'Direction':
        return cls(col=1, name=_right)

    @classmethod
    def up_left(cls) -> 'Direction':
        return cls(row=-1, col=-1)

    @classmethod
    def down_left(cls) -> 'Direction':
        return cls(row=1, col=-1)

    @classmethod
    def up_right(cls) -> 'Direction':
        return cls(row=-1, col=1)

    @classmethod
    def down_right(cls) -> 'Direction':
        return cls(row=1, col=1)


class Grid:
    def __init__(self, grid: list[Strs]):
        self._points = []

        for row, line in enumerate(grid):
            points = []

            for col, cell in enumerate(line):
                points.append(Point(row=row, col=col, mark=cell, grid=self))

            self._points.append(points)

        self.rows = len(self._points)
        self.cols = len(self._points[0])
        self.size = self.rows * self.cols

    def __iter__(self) -> typing.Iterator[list[Point]]:
        return iter(self._points)

    def __getitem__(self, point: Point) -> Point:
        return self._points[point.row][point.col]

    def __setitem__(self, point: Point, other: Point | str) -> None:
        if isinstance(other, Point):
            self._points[point.row][point.col] = Point(row=point.row, col=point.col, mark=other.mark, grid=self)
            self._points[other.row][other.col] = Point(row=other.row, col=other.col, mark=point.mark, grid=self)
        else:
            self._points[point.row][point.col] = Point(row=point.row, col=point.col, mark=other, grid=self)

    def point(self, row: int, col: int) -> Point:
        point = Point(row=row, col=col, mark='', grid=self)

        return self[point] if point.inside else point

    def find(self, marks: set[str]) -> Point:
        found = self.point(0, 0)

        for row in self:
            for point in row:
                if point.mark in marks:
                    self[point] = '.'
                    found = point

                    break

        return found

    def points(self) -> typing.Iterable[Point]:
        for row in self._points:
            yield from row


UniqPoints: typing.TypeAlias = set[Point]
Directions: typing.TypeAlias = list[Direction]
DayResult: typing.TypeAlias = int | str


@dataclasses.dataclass
class Day:
    indata: str

    def first_star(self) -> DayResult:
        raise NotImplementedError()  # pragma: no cover

    def second_star(self) -> DayResult:
        raise NotImplementedError()  # pragma: no cover
