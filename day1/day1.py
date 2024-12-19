import os

def part1(lines):
    sum = 0
    rnums = []
    lnums = []
    for line in lines:
        nums = line.split("   ")
        lnum = int(nums[0])
        rnum = int(nums[1])
        lnums.append(lnum)
        rnums.append(rnum)

    rnums.sort()
    lnums.sort()
    
    for i, lnum in enumerate(lnums):
        sum += abs(lnum - rnums[i])
    print(sum)

def part2(lines):
    sum = 0
    rnumFreq = {}

    for line in lines:
        rnum = int(line.split("   ")[1])
        rnumFreq[rnum] = rnumFreq.get(rnum, 0) + 1
    
    for line in lines:
        nums = line.split("   ")
        lnum = int(nums[0])
        rnum = int(nums[1]) 
        if lnum in rnumFreq:
            sum += lnum * rnumFreq[lnum]

    print(sum)

with open("input.txt", 'r') as inFile:
    lines = inFile.read().split('\n')
    part1(lines)
    part2(lines)

