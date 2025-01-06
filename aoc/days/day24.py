import dataclasses
import typing

from aoc.library import datatypes, funcs

_xor = 'XOR'

WiresType: typing.TypeAlias = dict[str, int]
ConnectionsType: typing.TypeAlias = dict[str, datatypes.Strs]


@dataclasses.dataclass
class Day(datatypes.Day):
    def parse(self) -> tuple[WiresType, ConnectionsType]:
        parts = funcs.parts(self.indata)

        wires = _parse_wires(parts[0])
        connections = {}

        for line in funcs.split(parts[1]):
            *details, _, res = line.split()
            left, operation, right = details

            wires.setdefault(left, -1)
            wires.setdefault(right, -1)
            wires.setdefault(res, -1)

            connections[res] = [left, operation, right]

        return wires, connections

    def first_star(self) -> datatypes.DayResult:
        wires, connections = self.parse()

        for wire in wires:
            _send_signal(wires, connections, wire)

        code = [str(wires[key]) for key in sorted(wires, reverse=True) if key.startswith('z')]

        return int(''.join(code), 2)

    def second_star(self) -> datatypes.DayResult:
        wires, connections = self.parse()
        # ---- thank you reddit for the help
        swaps = _reddit_magic(wires, connections)

        return ','.join(sorted(swaps))


def _reddit_magic(wires: WiresType, connections: ConnectionsType) -> datatypes.StrsSet:  # noqa:WPS231
    swaps = set()

    highest_z = max(wires)
    smallest_x = 'x00'

    for wire, (left, operation, right) in connections.items():
        if wire.startswith('z') and operation != _xor and wire != highest_z:
            swaps.add(wire)

        if operation == _xor and all(map(_uninteresting_wire, (wire, left, right))):
            swaps.add(wire)

        if operation == 'AND' and smallest_x not in {left, right}:
            swaps.update(_compare_connections(wire, connections, lambda opera: opera != 'OR'))

        if operation == _xor:
            swaps.update(_compare_connections(wire, connections, lambda opera: opera == 'OR'))

    return swaps


def _compare_connections(
    target: str, connections: ConnectionsType, predicate: typing.Callable[[str], bool]
) -> datatypes.StrsSet:
    swaps = set()

    for left, operation, right in connections.values():
        if (target == left or target == right) and predicate(operation):
            swaps.add(target)

    return swaps


def _uninteresting_wire(wire: str) -> bool:
    return wire[0] not in {'x', 'y', 'z'}


def _send_signal(wires: WiresType, connections: ConnectionsType, wire: str) -> int:
    if wires[wire] != -1:
        return wires[wire]

    left, operation, right = connections[wire]

    from_left = wires[left]
    while from_left == -1:
        from_left = _send_signal(wires, connections, left)

    from_right = wires[right]
    while from_right == -1:
        from_right = _send_signal(wires, connections, right)

    wires[wire] = _connect(from_left, operation, from_right)

    return wires[wire]


def _connect(left: int, operation: str, right: int) -> int:
    match operation:
        case 'AND':
            return left & right
        case 'OR':
            return left | right
        case _:
            return left ^ right


def _parse_wires(indata: str) -> dict[str, int]:
    wires = {}

    for line in funcs.split(indata):
        name, signal = line.split(': ')
        wires[name] = int(signal)

    return wires
