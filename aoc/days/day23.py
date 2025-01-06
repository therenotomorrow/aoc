import collections
import dataclasses
import typing

from aoc.library import datatypes, funcs

GraphType: typing.TypeAlias = dict[str, datatypes.StrsSet]


@dataclasses.dataclass
class Day(datatypes.Day):
    def parse(self) -> list[datatypes.StrsTuple]:
        connections: list[datatypes.StrsTuple] = []

        for line in funcs.split(self.indata):
            left, right = line.split('-')
            connections.append((left, right))

        return connections

    def first_star(self) -> datatypes.DayResult:
        graph = self._collect_graph()
        triples = set()

        for node, candidates in graph.items():
            if not node.startswith('t'):
                continue

            for candidate in candidates:
                triples |= {
                    tuple(sorted((node, candidate, neighbor)))
                    for neighbor in graph[candidate]
                    if neighbor in candidates
                }

        return len(triples)

    def second_star(self) -> datatypes.DayResult:
        graph = self._collect_graph()

        largest_len = 0
        largest_set = set()
        cliques = _find_connected(set(), set(graph), set(), graph)

        for clique in cliques:
            if len(clique) > largest_len:
                largest_set = clique
                largest_len = len(clique)

        return ','.join(sorted(largest_set))

    def _collect_graph(self) -> GraphType:
        connections = self.parse()
        graph = collections.defaultdict(set)

        for left, right in connections:
            graph[left].add(right)
            graph[right].add(left)

        return graph


def _find_connected(
    connected: datatypes.StrsSet,
    nodes: datatypes.StrsSet,
    excluded: datatypes.StrsSet,
    graph: GraphType,
) -> list[datatypes.StrsSet]:
    # uses Bron–Kerbosch algorithm
    if not nodes and not excluded:
        return [connected]

    cliques = []

    for vertex in list(nodes):
        cliques.extend(
            _find_connected(
                connected | {vertex},
                nodes & graph[vertex],
                excluded & graph[vertex],
                graph,
            )
        )

        nodes.remove(vertex)
        excluded.add(vertex)

    return cliques
