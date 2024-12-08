import argparse
import sys
import collections
import itertools


def printcity(width, height, antinodes):
    for r in range(height):
        for c in range(width):
            if (r, c) in antinodes:
                print('#', end='')
            else:
                print('.', end='')
        print()


def main(args: argparse.Namespace) -> int:
    antennas = collections.defaultdict(list)
    width, height = 0, 0
    with open(args.input, 'r') as file:
        lines = file.readlines()
        width = len(lines[0].strip())
        height = len(lines)
        for row, line in enumerate(lines):
            for col, char in enumerate(line.strip()):
                if char != '.':
                    antennas[char].append((row, col))

    antinodes = set()

    for freqlist in antennas.values():
        for a1, a2 in itertools.combinations(freqlist, r=2):

            dr = a2[0] - a1[0]
            dc = a2[1] - a1[1]

            if args.part == 1:
                antinodes.add((a1[0] - dr, a1[1] - dc))
                antinodes.add((a2[0] + dr, a2[1] + dc))
            else:
                curr = a1
                while 0 <= curr[0] < height and 0 <= curr[1] < width:
                    antinodes.add(curr)
                    curr = (curr[0] - dr, curr[1] - dc)

                curr = a1
                while 0 <= curr[0] < height and 0 <= curr[1] < width:
                    antinodes.add(curr)
                    curr = (curr[0] + dr, curr[1] + dc)

    antinodes = {a for a in antinodes if 0 <=
                 a[0] < height and 0 <= a[1] < width}

    if args.verbose:
        printcity(width, height, antinodes)

    print(len(antinodes))


if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('--part', type=int, default=1)
    parser.add_argument('--input', type=str, default='test08.txt')
    parser.add_argument('--verbose', type=bool, default=False)
    sys.exit(main(parser.parse_args()))
