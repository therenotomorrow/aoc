import dataclasses

from aoc.library import datatypes, funcs


@dataclasses.dataclass
class Day(datatypes.Day):
    def parse(self) -> list[datatypes.Ints]:
        reports = []

        for line in funcs.split(self.indata):
            reports.append(list(map(int, line.split())))

        return reports

    def first_star(self) -> datatypes.DayResult:
        return sum(_is_safe(report) for report in self.parse())

    def second_star(self) -> datatypes.DayResult:
        safe = 0

        for report in self.parse():
            for exclude, _ in enumerate(report):
                left, right = exclude, exclude + 1

                if _is_safe(report[:left] + report[right:]):
                    safe += 1

                    break

        return safe


def _is_safe(report: datatypes.Ints) -> bool:
    is_asc = (report[1] - report[0]) > 0
    diffs = []

    for idx, curr in enumerate(report[1:], start=1):
        diff = curr - report[idx - 1]

        if not is_asc:
            diff = -diff

        diffs.append(1 <= diff <= 3)

    return all(diffs)
