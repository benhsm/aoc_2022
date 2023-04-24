#!/usr/bin/python

with open('input.txt') as f:
    input = f.readlines()

field = [list(line[:-1]) for line in input]

starting_point = 0,0
ending_point = 0,0

for y, line in enumerate(field):
    for x, c in enumerate(line):
        if c == 'S':
            starting_point = y,x
            field[y][x] = ord('a')
        elif c == 'E':
            ending_point = y,x
            field[y][x] = ord('z')
        else:
            field[y][x] = ord(c)

directions = {
    "up": (-1, 0),
    "down": (1, 0),
    "left": (0, -1),
    "right": (0, 1),
}

def get_traversable_neighbours(point):
    res = []
    for d in directions.values():
        dy, dx = point[0] + d[0], point[1] + d[1]
        if dy < 0 or dx < 0 or dy >= len(field) or dx >= len(field[0]):
            continue
        py, px = point[0], point[1]
        if field[dy][dx] - field[py][px] <= 1:
            res.append((dy, dx))
    return res

def min_steps(start):
    to_visit = get_traversable_neighbours(start)
    visited = set()
    depth = 0
    while True:
        depth += 1 
        visit_next = []
        for n in to_visit:
            if n == ending_point:
                return(depth)
            for e in get_traversable_neighbours(n):
                if not e in visited:
                    visited.add(e)
                    visit_next.append(e)
        if not visit_next:
            return(-1)
        to_visit = visit_next

print(min_steps(starting_point))

candidates = []
for y, line in enumerate(field):
    for x, c in enumerate(line):
        if c == 97:
            candidates.append(min_steps((y,x)))
candidates = [x for x in candidates if x != -1]
print(min(candidates))
