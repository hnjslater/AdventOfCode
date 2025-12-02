import sys
import argparse
import collections
from textwrap import wrap


def main(args: argparse.Namespace) -> int:
    f = open(args.input, "r")
    result = 0
    for l in f.readlines():
        for entry in l.strip().split(","):
            if entry == "":
                continue
            start, end = (int(x) for x in entry.strip().split("-"))
            for x in range(start, end + 1):
                s = str(x)
                if args.part == 1:
                    if len(s) % 2 > 0:
                        continue
                    a, b = wrap(s, len(s) // 2)
                    if a == b:
                        result += x
                else:
                    for i in range(1, (len(s) // 2) + 1):
                        if len(s) % i > 0:
                            continue
                        items = wrap(s, i)

                        if all(a == items[0] for a in items):
                            result += x
                            break

    print(result)


if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument("--part", type=int, default=1)
    parser.add_argument("--input", type=str, default="test02.txt")
    parser.add_argument("--verbose", type=bool, default=False)
    sys.exit(main(parser.parse_args()))
