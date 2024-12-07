import argparse
import operator
import sys


def eval(test_value, nums, operators, total, i):
    if total > test_value:
        return False
    if i == len(nums):
        return total == test_value

    for op in operators:
        if eval(test_value, nums, operators, op(total,nums[i]), i + 1):
            return True

    return False


def main(args: argparse.Namespace) -> int:
    file = open(args.input, 'r')
    result = 0

    while line := file.readline():
        test_value_str, nums = line.split(':')
        test_value = int(test_value_str)
        nums = [int(x) for x in nums.strip().split(' ')]

        operators = [operator.add, operator.mul]
        if args.part == 2:
            operators += [lambda x,y: int(str(x) + str(y))]

        if eval(test_value, nums, operators, nums[0], 1):
            result += test_value

    print(result)


if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('--part', type=int, default=1)
    parser.add_argument('--input', type=str, default='test07.txt')
    sys.exit(main(parser.parse_args()))
