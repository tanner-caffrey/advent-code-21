# Author Tanner Caffrey
# file: main.py - Advent of code 2021 day 1

# Part 1

lines = []
with open('input.txt') as f:
    lines = f.readlines()
    f.close()

lines = list(map(lambda x: int(x), lines))

last = lines[0]
increased_count = 0
for n in lines:
    if n > last:
        increased_count += 1
    last = n

print("Part One Increases: ",increased_count)

# Part 2

index = 3
last = sum(lines[:3])
increased_count = 0
for n in range(len(lines)):
    if n <= 3: continue
    window = sum(lines[n-2:n+1])
    if window > last:
        increased_count += 1
    last = window

print("Part Two Increases: ",increased_count)
