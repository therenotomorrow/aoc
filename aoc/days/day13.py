import dataclasses
import typing

from aoc.library import datatypes, funcs


@dataclasses.dataclass(slots=True, kw_only=True, frozen=True)
class Button:
    cost_x: int
    cost_y: int


@dataclasses.dataclass(slots=True, kw_only=True, frozen=True)
class Prize:
    value_x: int
    value_y: int


NumPair: typing.TypeAlias = tuple[int, int]
Game = tuple[Button, Button, Prize]


@dataclasses.dataclass
class Day(datatypes.Day):
    def parse(self) -> list[Game]:
        button_a = Button(cost_x=0, cost_y=0)
        button_b = Button(cost_x=0, cost_y=0)
        prize = Prize(value_x=0, value_y=0)

        games = []
        times = 1

        for line in funcs.split(self.indata):
            match times:
                case 1:
                    cost_x, cost_y = _parse_nums(line.split('+'))
                    button_a = Button(cost_x=cost_x, cost_y=cost_y)
                case 2:
                    cost_x, cost_y = _parse_nums(line.split('+'))
                    button_b = Button(cost_x=cost_x, cost_y=cost_y)
                case 3:
                    value_x, value_y = _parse_nums(line.split('='))
                    prize = Prize(value_x=value_x, value_y=value_y)

            times += 1
            if times > 3:
                games.append((button_a, button_b, prize))
                times = 1

        return games

    def first_star(self) -> datatypes.DayResult:
        games = self.parse()
        shift = 0

        return sum(_resolve(*game, shift=shift) for game in games)

    def second_star(self) -> datatypes.DayResult:
        games = self.parse()
        shift = 10000000000000

        return sum(_resolve(*game, shift=shift) for game in games)


def _resolve(button_a: Button, button_b: Button, prize: Prize, *, shift: int) -> int:
    shifted = Prize(value_x=prize.value_x + shift, value_y=prize.value_y + shift)

    times_a, ok_a = _calc(button_b, button_a, shifted)
    times_b, ok_b = _calc(button_a, button_b, shifted)

    return 3 * times_a + times_b if ok_a and ok_b else 0


def _parse_nums(parts: datatypes.Strs) -> NumPair:
    left = int(parts[-2].split(',')[0])
    right = int(parts[-1])

    return left, right


def _calc(left: Button, right: Button, prize: Prize) -> tuple[int, bool]:
    divider = right.cost_x * left.cost_y - left.cost_x * right.cost_y
    dividend = prize.value_x * left.cost_y - prize.value_y * left.cost_x

    ans = dividend / divider

    return int(ans), ans == int(ans)
