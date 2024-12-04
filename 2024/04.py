import argparse
import sys


def search(needle: str, haystack: list[str], directions: list[tuple[int, int]]) -> int:
    result = []
    for row, line in enumerate(haystack):
        for col, char in enumerate(line):
            if char == needle[0]:
                for dr, dc in directions:
                    found = True
                    for i, char in enumerate(needle):
                        try:
                            if haystack[row+(i*dr)][col+(i*dc)] != needle[i]:
                                found = False
                                break
                        except (IndexError):
                            found = False
                            break
                    if found:
                        result += [(row+dr, col+dc)]
    return result


def main(args: argparse.Namespace) -> int:
    file = open(args.input, 'r').readlines()
    if args.part == 1:
        directions = [(1, -1), (1, 0), (1, 1), (0, 1)]
        backward = len(search('XMAS', file, directions))
        forward = len(search('SAMX', file, directions))
        print(forward + backward)
    else:
        nw_se = set(search('MAS', file, [(+1, +1)]) +
                    search('SAM', file, [(+1, +1)]))
        ne_sw = set(search('MAS', file, [(+1, -1)]) +
                    search('SAM', file, [(+1, -1)]))
        print(len(nw_se.intersection(ne_sw)))


if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('--part', type=int, default=1)
    parser.add_argument('--input', type=str, default='test04.txt')
    sys.exit(main(parser.parse_args()))
