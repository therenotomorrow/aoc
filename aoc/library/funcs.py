import typing

from aoc.library import datatypes


def split(indata: str) -> typing.Iterable[str]:
    yield from filter(bool, map(str.strip, indata.split('\n')))


def parts(indata: str, *, num: int = 2) -> datatypes.StrsTuple:
    comps = indata.split('\n\n')

    if num == 2:
        return comps[0], comps[1]

    return tuple(comps)
