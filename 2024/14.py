import argparse
import sys
import re
import collections
import os
import itertools

def main(args: argparse.Namespace) -> int:
    quads = [0,0,0,0]
    robots = []

    with open(args.input, "r") as file:
        while line := file.readline():
            robot = [int(x) for x in re.findall("p=(\\-?[0-9]+),(\\-?[0-9]+) v=(\\-?[0-9]+),(\\-?[0-9]+)", line.strip())[0]]
            robots.append(robot)


    if args.part == 1:
        for robot in robots:
            x,y,dx,dy = robot
            x2 = (x + 100*dx) % args.width
            y2 = (y + 100*dy) % args.height

            bottom = y2>(args.height/2)
            right = x2>(args.width/2)

            if x2 != int(args.width/2) and y2 != int(args.height/2):
                if bottom and right:
                    quads[3] += 1
                elif bottom:
                    quads[2] += 1
                elif right:
                    quads[1] += 1
                else:
                    quads[0] += 1
        print(quads[0] * quads[1] * quads[2] * quads[3])
    else:
        for steps in itertools.count(start=1):
            robotlocs = collections.defaultdict(int)
            
            for i,robot in enumerate(robots):
                x,y,dx,dy = robot
                robot =  ((x+dx) % args.width,(y+dy) % args.height,dx,dy)
                robots[i] = robot
                robotlocs[(robot[0], robot[1])] += 1
            
            if max(robotlocs.values()) == 1:
                break
        if args.verbose: 
            for j in range(args.height):
                for i in range(args.width):
                    c = (i,j)

                    if robotlocs[c] == 0:
                        print(".", end="")
                    else:
                        print(robotlocs[c], end="")
                print()

        print(steps)



if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('--part', type=int, default=1)
    parser.add_argument('--input', type=str, default="input14.txt")
    parser.add_argument('--verbose', type=bool, default=False)
    parser.add_argument('--width', type=int, default=101)
    parser.add_argument('--height', type=int, default=103)
    sys.exit(main(parser.parse_args()))
