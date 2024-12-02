import sys
import collections


def is_safe(line):
    increasing = line[1] > line[0]
    decreasing = line[0] > line[1]
    for i, curr in enumerate(line[1:], 1):
        diff = curr - line[i-1]
        if diff == 0:
            return (False, i)
        if increasing and diff < 0:
            return (False, i)
        if decreasing and diff > 0:
            return (False, i)
        if abs(diff) > 3:
            return (False, i)

    return (True, -1)


part = 1
if len(sys.argv) > 1 and sys.argv[1] == '2':
    part = 2

f = open('02.txt', 'r')
result = 0

for l in f:
    line = [int(x) for x in l.split()]

    safe, bad_idx = is_safe(line)

    if safe:
        result += 1
    elif part == 2:
        for i in range(len(line)):
            line2 = line.copy()
            del line2[i]
            safe, bad_idx = is_safe(line2)

            if safe:
                result += 1
                break

print(result)
