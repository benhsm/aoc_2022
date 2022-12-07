#!/usr/bin/awk -f
# version: GNU awk 5.2.1

BEGIN {
    FS = ","
    contains = 0
    overlaps = 0
}

{
    split($1, s1, "-")
    split($2, s2, "-")
    # let left pair assignment be s1, let right pair assignment be s2

    if (s1[1] >= s2[1] && s1[2] <= s2[2]) {
    # case: s2 fully contains s1
        contains++
    }
    else if (s2[1] >= s1[1] && s2[2] <= s1[2]) {
    # case: s1 fully contains s2
        contains++
    }

    if (s1[1] >= s2[1] && s2[2] >= s1[1]) {
    # case: pair assignment 2 starts ahead of pair assignment 1
        overlaps++
    } else if (s2[1] >= s1[1] && s1[2] >= s2[1]) {
    # case: pair assignment 1 starts ahead of pair assignment 2
        overlaps++
    }
}

END {
    print "The number of assignment pairs in which one range fully contains the other is: " contains
    print "The number of assignment pairs with any amount of overlap is: " overlaps
}
