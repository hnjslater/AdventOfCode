import argparse
import sys
from collections import defaultdict


def main(args: argparse.Namespace) -> int:
    book = defaultdict(lambda: -1)
    heads = set()

    with open(args.input, 'r') as file:
        for row, line in enumerate(file.readlines()):
            for col, char in enumerate(line.strip()):
                if char != '.':
                    book[(row, col)] = int(char)
                    if int(char) == 0:
                        heads.add((row, col))

    neighbours = ((0, 1), (0, -1), (1, 0), (-1, 0))

    result = []
    for h in heads:
        todo = set([(h,)])
        found_heads = set()
        found_tails = set()

        while len(todo) > 0:
            curr = todo.pop()
            for delta in neighbours:
                n = (curr[0][0] + delta[0], curr[0][1] + delta[1])
                if book[n] == book[curr[0]] + 1:
                    if book[n] == 9:
                        if args.part == 1:
                            found_heads.add(n)
                        else:
                            found_heads.add((n,) + curr)
                    else:
                        todo.add((n,) + curr)
        result += [len(found_heads)]
    print(sum(result))


if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('--part', type=int, default=1)
    parser.add_argument('--input', type=str, default='test10.txt')
    sys.exit(main(parser.parse_args()))
