""" Module day08: https://adventofcode.com/2022/day/8"""
import pprint


def is_tree_visible(tree: int, side: list[int]) -> bool:
    for t in side:
        if t >= tree:
            return False

    return True


def count_visible_trees(trees: list[list[int]]) -> int:
    vert_len = len(trees) - 1
    hor_len = len(trees[0]) - 1

    visible = 0
    # on edges trees always visible
    visible += 2 * (hor_len + vert_len)
    # on sides maybe not
    for i in range(1, vert_len):
        for j in range(1, hor_len):
            tree = trees[i][j]

            left_raw = trees[i][:j]
            right_raw = trees[i][hor_len:j:-1]
            up_raw = [t[j] for t in trees[:i]]
            down_raw = [t[j] for t in trees[vert_len:i:-1]]

            if any(
                (
                    is_tree_visible(tree, left_raw),
                    is_tree_visible(tree, right_raw),
                    is_tree_visible(tree, up_raw),
                    is_tree_visible(tree, down_raw),
                )
            ):
                print(tree, i, j)
                visible += 1

    return visible


def calc_visible_score(tree: int, side: list[int]) -> int:
    score = 0

    for t in side:
        score += 1
        if t >= tree:
            break

    return score


def count_scenic_score(trees: list[list[int]]) -> int:
    highest = 0

    vert_len = len(trees)
    hor_len = len(trees[0])

    for i in range(vert_len):
        for j in range(hor_len):
            tree = trees[i][j]

            left_raw = trees[i][:j][::-1]
            right_raw = trees[i][j + 1 :]
            # print(f'{left_raw=}', f'{tree=}', f'{right_raw=}')

            up_raw = [t[j] for t in trees[:i]][::-1]
            down_raw = [t[j] for t in trees[i + 1 :]]
            # print(f'{up_raw=}', f'{tree=}', f'{down_raw=}')

            score = (
                calc_visible_score(tree, left_raw)
                * calc_visible_score(tree, right_raw)
                * calc_visible_score(tree, up_raw)
                * calc_visible_score(tree, down_raw)
            )
            if score > highest:
                highest = score
            #     print(tree, i, j)
            #     visible += 1

    return highest


if __name__ == '__main__':
    trees = []

    with open('input.txt') as f:
        for line in f:
            line = line.strip()
            trees.append([int(tree) for tree in line])

    pprint.pprint(trees)
    # print(count_visible_trees(trees))
    print(count_scenic_score(trees))
