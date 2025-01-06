import collections
import dataclasses
import typing

from aoc.library import datatypes, funcs

Update: typing.TypeAlias = datatypes.Ints
Updates: typing.TypeAlias = list[Update]
Ruleset: typing.TypeAlias = dict[int, datatypes.IntsSet]


@dataclasses.dataclass
class Day(datatypes.Day):
    def parse(self) -> tuple[Updates, Ruleset]:
        parts = funcs.parts(self.indata)

        ruleset = _parse_ruleset(parts[0])
        updates = _parse_updates(parts[1])

        return updates, ruleset

    def first_star(self) -> datatypes.DayResult:
        updates, ruleset = self.parse()
        total = 0

        for update in updates:
            if _is_correct_update(update, ruleset):
                total += update[len(update) // 2]

        return total

    def second_star(self) -> datatypes.DayResult:
        updates, ruleset = self.parse()
        total = 0

        for update in updates:
            if _is_correct_update(update, ruleset):
                continue

            update = _sort_update(update, ruleset)
            total += update[len(update) // 2]

        return total


def _parse_ruleset(indata: str) -> Ruleset:
    ruleset = collections.defaultdict(set)

    for line in funcs.split(indata):
        before, after = tuple(map(int, line.split('|')))
        ruleset[before].add(after)

    return ruleset


def _parse_updates(indata: str) -> Updates:
    updates = []

    for line in funcs.split(indata):
        updates.append([int(num) for num in line.split(',')])

    return updates


def _is_correct_update(update: Update, rules: Ruleset) -> bool:
    used = set()

    for num in update:
        for related_page in rules.get(num, set()):
            if related_page in used:
                return False

        used.add(num)

    return True


def _sort_update(update: Update, rules: Ruleset) -> Update:
    return sorted(
        update,
        key=lambda rule: sum(rule in rules.get(num, set()) for num in update),
        reverse=True,
    )
