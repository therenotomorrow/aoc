Advent Of Code
==============

How I spent my third and **last** time solving [Advent Of Code](https://adventofcode.com/) for 2024. 
This was complicated, huge thanks to:

- [Reddit](https://www.reddit.com/r/adventofcode/)
- [tomakado](https://github.com/tomakado)
- [alexchao26](https://github.com/alexchao26)
- [mgtezak](https://github.com/mgtezak)

System Requirements
-------------------

```shell
python --version 
# Python 3.13.1
```

Installation
------------

```shell
pip install -U pip poetry
poetry --version
# Poetry (version 2.0.0) 
poetry install --no-root
```

Development
-----------

After adding some code samples run [`code.sh`](./scripts/code.sh)

```shell
make code
```

You can automate linter action before each commit by calling [`pre-commit`](./.githooks/pre-commit)

```shell
# setup pre-commit and check integrity
git config --local core.hooksPath .githooks
```

Testing
-------

Test your changes with simple run [`test.sh`](./scripts/test.sh)

```shell
# fast unit tests to be sure that no regression was 
make test/smoke
# same as test/smoke but with creation coverage report
make test/cover
# ---------- coverage: platform darwin, python 3.13.1-final-0 ----------
# Name                                 Stmts   Miss  Cover
# --------------------------------------------------------
# ...
# --------------------------------------------------------
# TOTAL                                 1802      0   100%
```

Solutions
---------

| Problem                                                        | Solution                        |
|:---------------------------------------------------------------|:--------------------------------|
| [Historian Hysteria](https://adventofcode.com/2024/day/1)      | [day01.py](./aoc/days/day01.py) |
| [Red-Nosed Reports](https://adventofcode.com/2024/day/2)       | [day02.py](./aoc/days/day02.py) |
| [Mull It Over](https://adventofcode.com/2024/day/3)            | [day03.py](./aoc/days/day03.py) |
| [Ceres Search](https://adventofcode.com/2024/day/4)            | [day04.py](./aoc/days/day04.py) |
| [Print Queue](https://adventofcode.com/2024/day/5)             | [day05.py](./aoc/days/day05.py) |
| [Guard Gallivant](https://adventofcode.com/2024/day/6)         | [day06.py](./aoc/days/day06.py) |
| [Bridge Repair](https://adventofcode.com/2024/day/7)           | [day07.py](./aoc/days/day07.py) |
| [Resonant Collinearity](https://adventofcode.com/2024/day/8)   | [day08.py](./aoc/days/day08.py) |
| [Disk Fragmenter](https://adventofcode.com/2024/day/9)         | [day09.py](./aoc/days/day09.py) |
| [Hoof It](https://adventofcode.com/2024/day/10)                | [day10.py](./aoc/days/day10.py) |
| [Plutonian Pebbles](https://adventofcode.com/2024/day/11)      | [day11.py](./aoc/days/day11.py) |
| [Garden Groups](https://adventofcode.com/2024/day/12)          | [day12.py](./aoc/days/day12.py) |
| [Claw Contraption](https://adventofcode.com/2024/day/13)       | [day13.py](./aoc/days/day13.py) |
| [Restroom Redoubt](https://adventofcode.com/2024/day/14)       | [day14.py](./aoc/days/day14.py) |
| [Warehouse Woes](https://adventofcode.com/2024/day/15)         | [day15.py](./aoc/days/day15.py) |
| [Reindeer Maze](https://adventofcode.com/2024/day/16)          | [day16.py](./aoc/days/day16.py) |
| [Chronospatial Computer](https://adventofcode.com/2024/day/17) | [day17.py](./aoc/days/day17.py) |
| [RAM Run](https://adventofcode.com/2024/day/18)                | [day18.py](./aoc/days/day18.py) |
| [Linen Layout](https://adventofcode.com/2024/day/19)           | [day19.py](./aoc/days/day19.py) |
| [Race Condition](https://adventofcode.com/2024/day/20)         | [day20.py](./aoc/days/day20.py) |
| [Keypad Conundrum](https://adventofcode.com/2024/day/21)       | [day21.py](./aoc/days/day21.py) |
| [Monkey Market](https://adventofcode.com/2024/day/22)          | [day22.py](./aoc/days/day22.py) |
| [LAN Party](https://adventofcode.com/2024/day/23)              | [day23.py](./aoc/days/day23.py) |
| [Crossed Wires](https://adventofcode.com/2024/day/24)          | [day24.py](./aoc/days/day24.py) |
| [Code Chronicle](https://adventofcode.com/2024/day/25)         | [day25.py](./aoc/days/day25.py) |
