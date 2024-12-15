import argparse
import sys
from collections import defaultdict
import itertools


def tadd(t1, t2):
    return (t1[0] + t2[0], t1[1] + t2[1])


def main(args: argparse.Namespace) -> int:
    garden = {}

    with open(args.input, 'r') as file:
        for row, line in enumerate(file.readlines()):
            for col, char in enumerate(line.strip()):
                garden[(row, col)] = char
    num_plots = 0
    plots = {}
    equivs = []
    for plot, plant in garden.items():
        up = (plot[0] - 1, plot[1])
        left = (plot[0], plot[1] - 1)
        up_plot = None
        left_plot = None

        if up in plots and garden[plot] == garden[up]:
            up_plot = plots[up]
        if left in plots and garden[plot] == garden[left]:
            left_plot = plots[left]

        if up_plot is not None and left_plot is not None and up_plot != left_plot:
            left_target = None
            up_target = None
            for s in equivs:
                if left_plot in s:
                    left_target = s
                if up_plot in s:
                    up_target = s
            if left_target is not None and up_target is not None and left_target != up_target:
                equivs.remove(left_target)
                equivs.remove(up_target)
                equivs.append(left_target | up_target)
            elif left_target is not None:
                left_target.add(up_plot)
            elif up_target is not None:
                up_target.add(left_plot)
            else:
                equivs.append(set([left_plot, up_plot]))
            plots[plot] = min(left_plot, up_plot)

        elif up_plot:
            plots[plot] = up_plot
        elif left_plot:
            plots[plot] = left_plot
        else:
            num_plots += 1
            plots[plot] = num_plots

    for pos, plot in plots.items():
        for s in equivs:
            if plot in s:
                plots[pos] = min(s)

    if args.part == 1:
        areas = defaultdict(int)
        perimeters = defaultdict(int)
        for plot, plant in plots.items():
            areas[plant] += 1

            neighbours = ((-1, 0), (+1, 0), (0, -1), (0, +1))

            for delta in neighbours:
                n = tadd(plot, delta)
                if n not in plots:
                    perimeters[plant] += 1
                elif plots[n] != plant:
                    perimeters[plant] += 1

        result = 0
        for plant in areas:
            result += areas[plant] * perimeters[plant]

        print(result)
    else:
        areas = defaultdict(int)
        corners = defaultdict(int)
        for plot, plant in plots.items():
            areas[plant] += 1
            neighbours = ((-1, 0), (0, +1), (+1, 0), (0, -1), (-1, 0))

            for d1, d2 in itertools.pairwise(neighbours):
                p0 = plant
                ns = [tadd(plot, d1), tadd(plot, d2), tadd(plot, tadd(d1, d2))]
                p1, p2, p3 = [-1 if n not in plots else plots[n] for n in ns]

                if p0 != p1 and p0 != p2:
                    corners[plant] += 1
                elif p0 == p1 and p0 == p2 and p0 != p3:
                    corners[plant] += 1

        result = 0
        for plant in areas:
            result += areas[plant] * corners[plant]

        print(result)


if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('--part', type=int, default=1)
    parser.add_argument('--input', type=str, default='test12.txt')
    sys.exit(main(parser.parse_args()))
