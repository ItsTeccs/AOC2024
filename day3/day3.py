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
    print(calcMuls(text))

def part2(text):
    regexStr = r'(mul\(\d+,\d+\)|do\(\)|don\'t\(\))'
    curMuls = ""
    mulStrs = ""
    for match in re.findall(regexStr, text):
        if match.startswith("mul"):
            curMuls += match
            print(curMuls)
        elif match.startswith("don't"):
            mulStrs += curMuls
            curMuls = ""
        elif match.startswith("do"):
            curMuls = ""
    print(calcMuls(mulStrs))

with open("input.txt") as inFile:
    text = inFile.read()
    part1(text)
    part2(text)
