import dataclasses

from aoc.library import datatypes, funcs


@dataclasses.dataclass(slots=True, kw_only=True)
class Computer:
    reg_a: int
    reg_b: int
    reg_c: int
    instructions = {0: 'adv', 1: 'bxl', 2: 'bst', 3: 'jnz', 4: 'bxc', 5: 'out', 6: 'bdv', 7: 'cdv'}

    def instruction(self, code: int) -> str:
        return self.instructions[code]

    def reset(self, reg_a: int) -> None:
        self.reg_a = reg_a
        self.reg_b = 0
        self.reg_c = 0

    def run(self, program: datatypes.Ints) -> datatypes.Ints:  # noqa:WPS231
        out = []
        idx = 0

        while idx < len(program):
            in_code = program[idx]
            op_code = program[idx + 1]
            idx += 2

            match self.instruction(in_code):  # noqa:WPS242
                case 'adv':
                    self.reg_a >>= self._combo(op_code)
                case 'bxl':
                    self.reg_b ^= op_code
                case 'bst':
                    self.reg_b = self._combo(op_code) % 8
                case 'jnz':
                    idx = op_code if self.reg_a else idx
                case 'bxc':
                    self.reg_b ^= self.reg_c
                case 'out':
                    out.append(self._combo(op_code) % 8)
                case 'bdv':
                    self.reg_b = self.reg_a >> self._combo(op_code)  # pragma: no cover
                case 'cdv':
                    self.reg_c = self.reg_a >> self._combo(op_code)

        return out

    def _combo(self, code: int) -> int:
        match code:
            case 4:
                return self.reg_a
            case 5:
                return self.reg_b
            case 6:
                return self.reg_c  # pragma: no cover
            case _:
                return code


@dataclasses.dataclass
class Day(datatypes.Day):
    def parse(self) -> tuple[Computer, datatypes.Ints]:
        lines = [line.split(': ')[-1] for line in funcs.split(self.indata)]

        comp = Computer(reg_a=int(lines[0]), reg_b=int(lines[1]), reg_c=int(lines[2]))
        program = [int(num) for num in lines[3].split(',')]

        return comp, program

    def first_star(self) -> datatypes.DayResult:
        comp, program = self.parse()
        out = comp.run(program)

        return ','.join(map(str, out))

    def second_star(self) -> datatypes.DayResult:
        comp, program = self.parse()

        reg_a = 0
        for idx in reversed(range(len(program))):
            reg_a <<= 3
            comp.reset(reg_a)

            while comp.run(program) != program[idx:]:
                reg_a += 1
                comp.reset(reg_a)

        return reg_a
