""" Module day10: https://adventofcode.com/2022/day/10"""


class CPU(object):
    def __init__(self):
        self._clock = 0
        self._register = 1
        # ---- part 1
        self.signal = 0
        # ---- part 2
        self._crt = []
        self._sprite = []

        self.draw_sprite()

    def draw_sprite(self):
        self._sprite = ['.' for _ in range(40)]

        mid = 39 if (m := self._register) > 39 else m
        left = 0 if (l := self._register - 1) < 0 else l
        right = 39 if (r := self._register + 1) > 39 else r

        self._sprite[left] = self._sprite[mid] = self._sprite[right] = '#'

    def draw_crt(self):
        pos = len(self._crt)
        self._crt.append(self._sprite[pos])

    def clock(self):
        self._clock += 1
        self.draw_crt()

        match self._clock:
            case 20 | 60 | 100 | 140 | 180 | 220:
                self.signal += self._clock * self._register
            case 40 | 80 | 120 | 160 | 200 | 240:
                print(''.join(self._crt))
                self._crt = []

    def register(self, val: int):
        self._register += val
        self.draw_sprite()

    def noop(self) -> None:
        self.clock()

    def addx(self, val: int) -> None:
        self.clock()
        self.clock()
        self.register(val)


if __name__ == '__main__':
    with open('input.txt') as f:
        commands = f.readlines()

    cpu = CPU()

    for command in commands:
        command = command.strip()

        if command == 'noop':
            cpu.noop()
        else:
            command, val = command.split()
            cpu.addx(int(val))

    print(cpu.signal)
