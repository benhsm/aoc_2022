#!/usr/bin/python
import numpy as np

with open('input8.txt') as f:
    input = f.readlines()

input = [list(line[:-1]) for line in input]
input = [[int(c) for c in line] for line in input]
grid = np.matrix(input)

trees_that_are_visible = 0
highest_scenic_score = 0

def visible(directions, height):
    sightlines  = []
    for d in directions:
        if d == []:
            sightlines.append(True)
        else:
            sightlines.append(visible_from_sightline(d, height))
        
    if True in sightlines:
        return True
    return False

def visible_from_sightline(direction, height):
    for tree in direction:
        if tree >= height:
            return False
    return True

def trees_in_sight(directions, height):
    sightlines = []
    for d in directions:
        if d == []:
            sightlines.append(0)
        else:
            sightlines.append(trees_in_sight_in_direction(d, height))
    res = 1
    for s in sightlines:
        res = res * s
    return res

def trees_in_sight_in_direction(direction, height):
    res = 0
    for tree in direction:
        if tree >= height:
            return res+1
        res += 1
    return res

with np.nditer(grid, order='K', flags=['multi_index']) as trees:
    for tree in trees:
        c, r = trees.multi_index[0], trees.multi_index[1]
        left = grid[c,:r]
        right = grid[c,r+1:]
        above = grid[:c,r]
        below = grid[c+1:,r]

        directions = [x.flatten().tolist()[0] for x in [left, right, above, below]]
        directions[0].reverse()
        directions[2].reverse()

        if visible(directions, trees[0]):
            trees_that_are_visible += 1

        sightline_score = trees_in_sight(directions, trees[0])
        if sightline_score > highest_scenic_score:
            highest_scenic_score = sightline_score

print(trees_that_are_visible)
print(highest_scenic_score)
