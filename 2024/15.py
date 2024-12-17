import argparse
import sys
from typing import Dict, Tuple


def push(world: Dict[Tuple[int, int], chr], robot: Tuple[int,
         int], x: int, y: int, dx: int, dy: int, dryrun=True):

    if (x, y) not in world:
        return True
    if world[(x, y)] == '#':
        return False
    can_push = False

    if world[(x, y)] == '[' and dy != 0:
        can_push = push(
            world,
            robot,
            x + dx,
            y + dy,
            dx,
            dy,
            dryrun) and push(
            world,
            robot,
            x + dx + 1,
            y + dy,
            dx,
            dy,
            dryrun)
    elif world[(x, y)] == ']' and dy != 0:
        can_push = push(
            world,
            robot,
            x + dx,
            y + dy,
            dx,
            dy,
            dryrun) and push(
            world,
            robot,
            x + dx - 1,
            y + dy,
            dx,
            dy,
            dryrun)
    elif push(world, robot, x + dx, y + dy, dx, dy, dryrun):
        can_push = True
    if can_push:
        if not dryrun and (x, y) in world:
            if world[(x, y)] == '[':
                world[(x + dx, y + dy)] = world[(x, y)]
                del world[(x, y)]
                if (x + 1, y) in world:
                    world[(x + dx + 1, y + dy)] = world[(x + 1, y)]
                    del world[(x + 1, y)]
            elif world[(x, y)] == ']':
                world[(x + dx, y + dy)] = world[(x, y)]
                del world[(x, y)]
                if (x - 1, y) in world:
                    world[(x + dx - 1, y + dy)] = world[(x - 1, y)]
                    del world[(x - 1, y)]
            else:
                world[(x + dx, y + dy)] = world[(x, y)]
                del world[(x, y)]
        return True
    else:
        return False


def printworld(world, robot):
    width, height = max(world.keys())

    for y in range(height + 1):
        for x in range(width + 1):
            if (x, y) == robot:
                print('@', end='')
            elif (x, y) not in world:
                print('.', end='')
            else:
                print(world[(x, y)], end='')
        print()


def main(args: argparse.Namespace) -> int:
    world = {}
    robot = (0, 0)
    y = 0
    with open(args.input, 'r') as file:
        while line := file.readline():
            if line.strip() == '':
                break

            for x, char in enumerate(line.strip()):
                match char:
                    case '#':
                        world[(x, y)] = '#'
                    case 'O':
                        world[(x, y)] = 'O'
                    case '@':
                        world[(x, y)] = '@'
                        robot = (x, y)

            y += 1

        if args.part == 2:
            world2 = {}
            for k, v in world.items():
                x, y = k
                match v:
                    case 'O':
                        world2[(2 * x, y)] = '['
                        world2[(2 * x + 1, y)] = ']'
                    case '#':
                        world2[(2 * x, y)] = '#'
                        world2[(2 * x + 1, y)] = '#'
                    case _:
                        world2[(2 * x, y)] = '@'
            robot = (2 * robot[0], robot[1])
            world = world2

        while line := file.readline():
            for cmd in line.strip():
                dx, dy = 0, 0
                match cmd:
                    case '^':
                        dy = -1
                    case '>':
                        dx = +1
                    case 'v':
                        dy = +1
                    case '<':
                        dx = -1

                if push(world, robot, robot[0], robot[1], dx, dy, dryrun=True):
                    push(
                        world,
                        robot,
                        robot[0],
                        robot[1],
                        dx,
                        dy,
                        dryrun=False)
                    robot = (robot[0] + dx, robot[1] + dy)

    if args.verbose:
        printworld(world, robot)

    print(sum([100 * y + x for (x, y) in world.keys()
          if world[(x, y)] == 'O' or world[(x, y)] == '[']))


if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('--part', type=int, default=1)
    parser.add_argument('--verbose', type=bool, default=False)
    parser.add_argument('--input', type=str, default='test15.txt')
    sys.exit(main(parser.parse_args()))
