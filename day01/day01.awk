#!/usr/bin/awk
# version: GNU Awk 5.2.1

BEGIN { i = 0; sum = 0; }
{
    if (!NF) {
        data[i] = sum
        sum = 0
        i++
    } else {
        sum += $1
    }
}
END {
    n = asort(data)
    print data[n]
    print data[n]+data[n-1]+data[n-2]
}

