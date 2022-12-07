#!/usr/bin/lua
-- lua version 5.4.4
-- usage: <script name> <input file path>

l = require("lpeg")
io.input(arg[1]);

contains = 0
overlaps = 0

d = l.C(l.R"09"^-2)
pat = l.Ct(d*l.P"-"*d)*l.P","*l.Ct(d*l.P"-"*d)

for line in io.lines() do

    p1, p2 = l.match(pat, line)

    p1[1] = tonumber(p1[1])
    p1[2] = tonumber(p1[2])
    p2[1] = tonumber(p2[1])
    p2[2] = tonumber(p2[2])

    if p1[1] >= p2[1] and p1[2] <= p2[2] then
        contains = contains + 1
    elseif p2[1] >= p1[1] and p2[2] <= p1[2] then
        contains = contains + 1
    end

    if p1[1] >= p2[1] and p2[2] >= p1[1] then
        overlaps = overlaps + 1
    elseif p2[1] >= p1[1] and p1[2] >= p2[1] then
        overlaps = overlaps + 1
    end
end

print(contains, overlaps)
