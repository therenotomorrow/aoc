from aoc.library import datatypes

names = 'indata, want'


def ids() -> datatypes.StrsTuple:
    return 'example', 'result'


def first_star(day: datatypes.Day, want: datatypes.DayResult) -> None:
    got = day.first_star()

    assert got == want, f'first_star() = {got}, want = {want}'  # noqa:S101


def second_star(day: datatypes.Day, want: datatypes.DayResult) -> None:
    got = day.second_star()

    assert got == want, f'second_star() = {got}, want = {want}'  # noqa:S101
