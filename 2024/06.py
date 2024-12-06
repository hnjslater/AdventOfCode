import argparse
import sys
from collections import defaultdict


def turn(guard):
    dr, dc = guard[2], guard[3]
    match (dr, dc):
        case (-1, 0):
            dr, dc = 0, 1
        case (0, 1):
            dr, dc = 1, 0
        case (1, 0):
            dr, dc = 0, -1
        case (0, -1):
            dr, dc = -1, 0
    return (guard[0], guard[1], dr, dc)


def step(guard):
    return (guard[0] + guard[2], guard[1] + guard[3], guard[2], guard[3])


def pos(guard):
    return (guard[0], guard[1])


def printworld(guard, world, height, width):
    for row in range(height + 1):
        for col in range(width + 1):
            if (row, col) in world:
                print(world[(row, col)], end='')
            elif row == guard[0] and col == guard[1]:
                print('G', end='')
            else:
                print(' ', end='')

        print()


def patrol(guard, world):
    past = set()
    places = set()
    while True:
        places.add(pos(guard))
        past.add(guard)
        next_guard = step(guard)
        if world[pos(next_guard)] == '#':
            guard = turn(guard)
        else:
            guard = next_guard
        if world[pos(guard)] == 'X':
            return ('escaped', places)
        elif guard in past:
            return ('looped', places)


def main(args: argparse.Namespace) -> int:
    file = open(args.input, 'r')

    world = defaultdict(str)

    height = 1
    guard = None
    width = 0
    while line := file.readline():
        width = len(line)
        for col, char in enumerate(line):
            match char:
                case '#':
                    world[(height, col + 1)] = '#'
                case '^':
                    guard = (height, col + 1, -1, 0)

        height += 1

    for i in range(width):
        world[0, i] = 'X'
        world[height, i] = 'X'
    for i in range(height + 1):
        world[i, 0] = 'X'
        world[i, width] = 'X'

    if args.verbose:
        printworld(guard, world, height, width)
    _, places = patrol(guard, world)

    if args.part == 1:
        print(len(places))
    elif args.part == 2:
        result = 0
        for place in places:
            if place == guard:
                next
            world[place] = '#'
            endstate, _ = patrol(guard, world)
            if endstate == 'looped':
                result += 1
            del world[place]
        print(result)


if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('--part', type=int, default=1)
    parser.add_argument('--input', type=str, default='test06.txt')
    parser.add_argument('--verbose', type=bool, default=False)
    sys.exit(main(parser.parse_args()))
