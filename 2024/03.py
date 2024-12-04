import re
import sys

part = 1
if len(sys.argv) > 1 and sys.argv[1] == '2':
    part = 2

f = open('03.txt', 'r')
result = 0
do = True

for line in f.readlines():
    for x in re.findall("mul\\(([0-9]{1,3}),([0-9]{1,3})\\)|(do)\\(\\)|(don't)\\(\\)", line):
        if x[2] == 'do':
            do = True
        elif x[3] == "don't":
            do = False
        elif do or part == 1:
            result += int(x[0]) * int(x[1])

print(result)
