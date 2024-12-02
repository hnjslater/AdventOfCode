import sys
import collections


def is_safe(line):
    increasing = line[1] > line[0]
    decreasing = line[0] > line[1]
    for i, curr in enumerate(line[1:], 1):
        diff = curr - line[i-1]
        if diff == 0:
            return False
        if increasing and diff < 0:
            return False
        if decreasing and diff > 0:
            return False
        if abs(diff) > 3:
            return False

    return (True, -1)


part = 1
if len(sys.argv) > 1 and sys.argv[1] == '2':
    part = 2

f = open('02.txt', 'r')
result = 0

for l in f:
    line = [int(x) for x in l.split()]

    if is_safe(line):
        result += 1
    elif part == 2:
        for i in range(len(line)):
            line2 = line.copy()
            del line2[i]
            if is_safe(line2):
                result += 1
                break

print(result)
