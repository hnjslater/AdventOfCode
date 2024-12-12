import argparse
import sys
import functools


@functools.cache
def blink(stone: int, blinks_left: int) -> int:
    if blinks_left == 0:
        return 1
    stonestr = str(stone)
    if stone == 0:
        return blink(1, blinks_left - 1)
    elif len(stonestr) % 2 == 0:
        stlen = len(stonestr)
        return blink(int(stonestr[0:int(stlen / 2)]), blinks_left - 1) + \
            blink(int(stonestr[int(stlen / 2):]), blinks_left - 1)
    else:
        return blink(stone * 2024, blinks_left - 1)


def main(args: argparse.Namespace) -> int:
    with open(args.input, 'r') as file:
        line = [int(x) for x in file.readline().split()]

    if args.part == 1:
        print(sum([blink(s, 25) for s in line]))
    else:
        print(sum([blink(s, 75) for s in line]))


if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('--part', type=int, default=1)
    parser.add_argument('--input', type=str, default='test11.txt')
    sys.exit(main(parser.parse_args()))
