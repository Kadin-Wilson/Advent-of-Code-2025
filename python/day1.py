from sys import stdin

dial = 50
zeros = 0

for line in stdin:
    rot = int(line[1:-1])

    dir = 1
    if line[0] == 'L':
        dir = -1

    while rot:
        dial += dir

        if dial < 0:
            dial += 100
        if dial > 99:
            dial -= 100

        if dial == 0:
            zeros += 1

        rot -= 1

print(zeros)
