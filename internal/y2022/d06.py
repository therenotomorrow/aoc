""" Module day06: https://adventofcode.com/2022/day/6"""


def is_unique(subseq: str) -> bool:
    return len(set(subseq)) == len(subseq)


def find_signal(sequence: str, marker: int) -> int:
    signal = 0

    for i in range(0, len(sequence)):
        if is_unique(sequence[i : i + marker]):
            signal = i + marker
            break

    return signal
