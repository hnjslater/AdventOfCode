import sys
import argparse
import collections


def main(args: argparse.Namespace) -> int:
    f = open(args.input, "r")
    result = 0
    dial = 50
    for l in f.readlines():
        if l[0] == "L":
            d = -1
        else:
            d = 1
        n = int(l[1:].strip()) * d

        prev = dial
        dial = dial + n

        if args.part == 2:
            if n > 0:
                result += max(0, (dial - 1) // 100 - (prev - 1) // 100)
            elif n < 0:
                result += max(0, (prev) // 100 - (dial) // 100)

        dial = dial % 100
        if args.part == 1 and dial == 0:
            result += 1

    print("=", result)


if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument("--part", type=int, default=1)
    parser.add_argument("--input", type=str, default="test01.txt")
    parser.add_argument("--verbose", type=bool, default=False)
    sys.exit(main(parser.parse_args()))
