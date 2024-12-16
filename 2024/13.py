import argparse
import re
import sys


def main(args: argparse.Namespace) -> int:
    result = 0
    with open(args.input, 'r') as file:
        while btnA := file.readline():
            btnB = file.readline()
            prize = file.readline()
            file.readline()

            btnre = 'Button .: X\\+([0-9]+), Y\\+([0-9]+)'
            prizere = 'Prize: X=([0-9]+), Y=([0-9]+)'

            ax, ay = [int(x) for x in re.findall(btnre, btnA.strip())[0]]
            bx, by = [int(x) for x in re.findall(btnre, btnB.strip())[0]]
            px, py = [int(x) for x in re.findall(prizere, prize.strip())[0]]

            if args.part == 2:
                px = px + 10000000000000
                py = py + 10000000000000

            ap = (px * by - py * bx) / (ax * by - ay * bx)
            bp = (ax * py - ay * px) / (ax * by - ay * bx)

            if ap == int(ap) and bp == int(bp):
                result += ap * 3 + bp

    print(result)


if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('--part', type=int, default=1)
    parser.add_argument('--input', type=str, default='input13.txt')
    sys.exit(main(parser.parse_args()))
