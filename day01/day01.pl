#!/usr/bin/perl
use v5.36;
use strict;
use warnings;

my @calories = ();
my $sum = 0;
while (<<>>) {
    chomp;
    if (!$_) {
        push(@calories, $sum);
        $sum = 0;
    } else {
        $sum += $_;
    }
}
@calories = sort { $b <=> $a } @calories;
say $calories[0];
say $calories[0]+$calories[1]+$calories[2];
