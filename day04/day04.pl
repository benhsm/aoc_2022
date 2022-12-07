#!/usr/bin/perl
use v5.36;
use strict;
use warnings;

my $contains = 0;
my $overlaps = 0;

while (<<>>) {
   /(\d+)-(\d+),(\d+)-(\d+)/;
   if (($1 >= $3 && $2 <= $4) || ($3 >= $1 && $4 <= $2)) {
       $contains++
   }
   if (($1 >= $3 && $4 >= $1) || ($3 >= $1 && $2 >= $3)) {
       $overlaps++
   }
}
say "$contains $overlaps"
