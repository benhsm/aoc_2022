#!/usr/bin/python

with open("input10.txt") as f:
    input = f.read()

lines = input.splitlines()

cycle = 0
sum_signal_strengths = 0
x = 1

crt = [[False for _ in range(40)] for _ in range(6)]

def display_crt():
    for row in crt:
        for pixel in row:
            if pixel == True:
                print("#", end="")
            else:
                print(".", end="")
        print("")

def inc_cycle():
    global crt
    global cycle
    row = cycle // 40
    col = cycle % 40
    if x-1 <= col <= x+1:
        crt[row][col] = True
#        display_crt()
    cycle += 1
    if (cycle - 20) % 40 == 0:
        print(f"On cycle {cycle}, x*cycle is {x*cycle}")
        global sum_signal_strengths
        sum_signal_strengths += x * cycle


def run():
    global lines
    global x
    for l in lines:
        if l[:4] == "noop":
            inc_cycle()
        elif l[:4] == "addx":
            inc_cycle()
            inc_cycle()
            x += int(l.split(" ")[1])

run()
print(f"The sum of the signal strengths is {sum_signal_strengths}")

display_crt()
