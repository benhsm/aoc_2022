#!/usr/bin/python

DISK_SPACE = 70000000
NEEDED_FOR_UPDATE = 30000000

with open('input7.txt') as f:
    data = f.read()

dirstack = []
dirs = {}

def add_size_to_all_parents(size):
    for dir in dirstack:
        if dir in dirs:
            dirs[dir] += size
        else:
            dirs[dir] = size

for line in data[0:-1].split("\n"):
    t = line.split()
    if t[0] == '$':
        if t[1] == 'cd':
            if t[2] == '..':
                dirstack.pop()
            else:
                dirstack.append("".join(dirstack)+t[2])
        elif t[1] == 'ls':
            pass # we don't need ls for anything
    elif t[0] == 'dir':
        pass # we don't need dir for anything
    else:
        add_size_to_all_parents(int(t[0]))

deletion_candidates = []
for d in dirs.values():
    if d < 100000:
        deletion_candidates.append(d)

# part 1
print(sum(deletion_candidates))

# part 2
for d in sorted(dirs.items(), key=lambda x:x[1]):
    used_space = dirs['/']
    free_space = DISK_SPACE - used_space
    must_be_freed = NEEDED_FOR_UPDATE - free_space
    if d[1] > must_be_freed:
        print(d[1])
        break

