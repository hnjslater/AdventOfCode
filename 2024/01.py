import sys
import collections

f = open('01.txt', 'r')
lines = f.readlines()
lines = [l.split() for l in lines]

left = [int(x) for x, _ in lines]
right = [int(x) for _, x in lines]

if len(sys.argv) == 1 or sys.argv[1] == '1':
    print(sum([abs(y-x) for x, y in zip(sorted(left), sorted(right))]))
else:
    counts = collections.defaultdict(int)

    for x in right:
        counts[x] += 1

    print(sum([x * counts[x] for x in left]))
