""" Module day04: https://adventofcode.com/2022/day/4"""

Pair = tuple[str, ...]


def create_pairs(pairs_s: str) -> Pair:
    return tuple(pairs_s.split(','))


def is_pairs_fully_overlapped(pairs_t: Pair) -> bool:
    left, right = pairs_t

    left_l, left_r = map(int, left.split('-'))
    right_l, right_r = map(int, right.split('-'))

    if right_l <= left_l <= left_r <= right_r:
        return True

    if left_l <= right_l <= right_r <= left_r:
        return True

    return False


def is_pairs_overlapped(pairs_t: Pair) -> bool:
    left, right = pairs_t

    left_l, left_r = map(int, left.split('-'))
    right_l, right_r = map(int, right.split('-'))

    left_s = set(range(left_l, left_r + 1))
    right_s = set(range(right_l, right_r + 1))

    inter = left_s.intersection(right_s)

    return len(inter) > 0


def sum_fully_overlapped(pairs: list[Pair]) -> int:
    return sum(is_pairs_fully_overlapped(pair) for pair in pairs)


def sum_overlapped(pairs: list[Pair]) -> int:
    return sum(is_pairs_overlapped(pair) for pair in pairs)
