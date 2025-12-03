import sys
import argparse
import collections
import itertools


def main(args: argparse.Namespace) -> int:
    with open(args.input, "r") as f:
        result = 0
        for line in f.readlines():
            line = line.strip()
            removed = 0
            if args.part == 1:
                target = 2
            else:
                target = 12

            while len(line) > target:
                toremove = -1

                for i, (a, b) in enumerate(itertools.pairwise(line)):
                    if int(a) < int(b):
                        toremove = i
                        break
                # If there's no better option, LSD
                if toremove == -1:
                    toremove = len(line) - 1

                line = line[:toremove] + line[toremove + 1 :]
                removed += 1
            result += int(line)

        print(result)


if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument("--part", type=int, default=1)
    parser.add_argument("--input", type=str, default="test03.txt")
    parser.add_argument("--verbose", type=bool, default=False)
    sys.exit(main(parser.parse_args()))
