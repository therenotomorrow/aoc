import collections
import dataclasses
import functools

from aoc.library import datatypes, funcs

_times = 2000
_prune = 16777216
_num64 = 64
_num32 = 32
_num2048 = 2048


@dataclasses.dataclass
class Day(datatypes.Day):
    def parse(self) -> datatypes.Ints:
        return [num for num in map(int, funcs.split(self.indata))]

    def first_star(self) -> datatypes.DayResult:
        secrets = self.parse()
        ans = 0

        for secret in secrets:
            for _ in range(_times):
                secret = _generate_secret(secret)

            ans += secret

        return ans

    def second_star(self) -> datatypes.DayResult:
        secrets = self.parse()
        bananas: dict[datatypes.IntsTuple, int] = collections.defaultdict(int)

        for secret in secrets:
            seen = set()
            four_diffs = []

            for _ in range(_times):
                new_secret = _generate_secret(secret)
                four_diffs.append(_last_digit(new_secret) - _last_digit(secret))
                secret = new_secret

                if len(four_diffs) > 4:
                    four_diffs = four_diffs[1:]

                four_changes = tuple(four_diffs)

                if four_changes not in seen:
                    bananas[four_changes] += _last_digit(secret)
                    seen.add(four_changes)

        return max(bananas.values())


@functools.cache
def _generate_secret(secret: int) -> int:
    secret = (secret ^ secret * _num64) % _prune
    secret = (secret ^ secret // _num32) % _prune
    return (secret ^ secret * _num2048) % _prune


@functools.cache
def _last_digit(secret: int) -> int:
    return secret % 10
