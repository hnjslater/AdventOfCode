import argparse
import sys


def main(args: argparse.Namespace) -> int:
    file = open(args.input, 'r')
    rules = []
    result = 0

    while line := file.readline():
        line = line.strip()
        if line == '':
            break
        sline = line.split('|')
        rules.append((int(sline[0]), int(sline[1])))

    while line := file.readline():
        correct = True
        update = {}

        for position, page in enumerate(
                [int(x) for x in line.strip().split(',')]):
            update[page] = position
            changed = True
            while changed:
                changed = False
                for r in rules:
                    p1, p2 = r
                    if p1 in update and p2 in update and update[p1] > update[p2]:
                        correct = False
                        if args.part == 2:
                            changed = True
                            update = {page: position * 2 for page,
                                      position in update.items()}
                            update[p1] = update[p2] - 1

        if (correct and args.part == 1) or (not correct and args.part == 2):
            raw_update = sorted([(position, page)
                                for page, position in update.items()])
            result += raw_update[int((len(raw_update) - 1) / 2)][1]

    print(result)


if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('--part', type=int, default=1)
    parser.add_argument('--input', type=str, default='input05.txt')
    sys.exit(main(parser.parse_args()))
