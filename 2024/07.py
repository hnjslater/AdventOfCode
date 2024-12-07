import argparse
import sys


def eval(test_value, nums, total, i, concat):
    if total > test_value:
        return False
    if i == len(nums):
        return total == test_value

    if eval(test_value, nums, total * nums[i], i + 1, concat):
        return True
    if eval(test_value, nums, total + nums[i], i + 1, concat):
        return True
    if concat and eval(test_value, nums, int(
            str(total) + str(nums[i])), i + 1, concat):
        return True

    return False


def main(args: argparse.Namespace) -> int:
    file = open(args.input, 'r')
    result = 0

    while line := file.readline():
        test_value_str, nums = line.split(':')
        test_value = int(test_value_str)
        nums = [int(x) for x in nums.strip().split(' ')]

        if eval(test_value, nums, nums[0], 1, args.part == 2):
            result += test_value

    print(result)


if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('--part', type=int, default=1)
    parser.add_argument('--input', type=str, default='test07.txt')
    sys.exit(main(parser.parse_args()))
