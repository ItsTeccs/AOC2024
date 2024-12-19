import os

def safe(nums):
    if not nums:
        return False

    valid = True
    is_monotonic_incr = True
    is_monotonic_decr = True

    for i in range(len(nums)-1):
        diff = abs(nums[i] - nums[i+1])
        if not (1 <= diff and diff <= 3):
            valid = False
            break

        if nums[i] < nums[i+1]:
            is_monotonic_incr = False
        elif nums[i] > nums[i+1]:
            is_monotonic_decr = False

    return (is_monotonic_incr or is_monotonic_decr) and valid

def part1(lines):
    sum = 0
    for line in lines:
        if line:
            nums = list(map(int, line.split(" ")))
            if safe(nums):
                sum += 1
    print(sum)

def part2(lines):
    ans = 0
    for line in lines:
        if not line:
            continue

        nums = list(map(int, line.split(" ")))

        if safe(nums):
            ans += 1
        else:
            for i in range(len(nums)):
                subarr = nums.copy()
                subarr.pop(i)
                if safe(subarr):
                    ans += 1
                    break
    print(ans)

with open("input.txt") as inFile:
    lines = inFile.read().split('\n')
    part1(lines)
    part2(lines)

