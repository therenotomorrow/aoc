""" Module day02: https://adventofcode.com/2022/day/2"""
import typing

# A, X - Rock (1)
# B, Y - Paper(2)
# C, Z - Scissors (3)
Choice = typing.Literal['A', 'B', 'C', 'X', 'Y', 'Z']
# X - lose, Y - draw, Z - win
Desire = typing.Literal['X', 'Y', 'Z']
# part 1 - decide, part 2 - desire
Strategy = typing.Literal['decide', 'desire']
DecideOutcome = typing.Callable[[tuple[Choice, Choice]], int]
DesireOutcome = typing.Callable[[tuple[Choice, Desire]], int]

scores = {'X': 1, 'Y': 2, 'Z': 3}


def decide_outcome(rnd: tuple[Choice, Choice]) -> int:
    _, ch2 = rnd

    match rnd:
        case ('A', 'Z') | ('B', 'X') | ('C', 'Y'):
            outcome = 0
        case ('A', 'X') | ('B', 'Y') | ('C', 'Z'):
            outcome = 3
        case ('A', 'Y') | ('B', 'Z') | ('C', 'X'):
            outcome = 6
        case _:
            raise ValueError(f'{rnd=} is not expected')

    return outcome + scores[ch2]


def desire_outcome(rnd: tuple[Choice, Desire]) -> int:
    ch1, _ = rnd

    match rnd:
        case ('A', 'X') | ('B', 'Z') | ('C', 'Y'):
            outcome = decide_outcome((ch1, 'Z'))
        case ('A', 'Y') | ('B', 'X') | ('C', 'Z'):
            outcome = decide_outcome((ch1, 'X'))
        case ('A', 'Z') | ('B', 'Y') | ('C', 'X'):
            outcome = decide_outcome((ch1, 'Y'))
        case _:
            raise ValueError(f'{rnd=} is not expected')

    return outcome


def choose_strategy(strategy: Strategy) -> DecideOutcome | DesireOutcome:
    match strategy:
        case 'decide':
            fn = decide_outcome
        case 'desire':
            fn = desire_outcome
        case _:
            raise ValueError(f'{strategy=} is not excepted')

    return fn


def play(rounds: list[tuple[Choice, Choice]], strategy: Strategy) -> int:
    fn = choose_strategy(strategy)

    return sum(fn(rnd) for rnd in rounds)
