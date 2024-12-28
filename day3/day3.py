import os
import re

mulRegex = r'mul\(\d{0,3},\d{0,3}\)'
doRegex = r'do\(\)'
dontRegex = r'don\'t\(\)' 

def calcMuls(text):
    res = 0
    for match in re.findall(mulRegex, text):
        match = match[4:]
        match = match[:len(match)-1]
        num1 = int(match.split(',')[0])
        num2 = int(match.split(',')[1])
        res += num1 * num2
    return res

def part1(text):
    calcMuls(text)

def part2(text):
    validText = ""

    dontMatch = re.search(dontRegex, text)
    validText += text[:dontMatch.start()]
    text = text[dontMatch.end():]
    dontMatch = re.search(dontRegex, text)
    # print(validText)
    # print("---")
    # print(text)

    while dontMatch:
        doMatch = re.search(doRegex, text)
        if doMatch:
            validText += text[doMatch.end():dontMatch.start()]
            text = text[dontMatch.end():]
        dontMatch = re.search(dontRegex, text)

    doMatch = re.search(doRegex, text)
    if doMatch:
        validText += text[doMatch.end():]
    print("Part 2 result: " + str(calcMuls(validText)))
    # print(text)

with open("input.txt") as inFile:
    text = inFile.read()
    part1(text)
    part2(text)
