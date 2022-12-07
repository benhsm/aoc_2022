#!/usr/bin/perl
use v5.36;
use strict;
use warnings;

my @stacks;
my @stacks2;
my @tmp_stack;

while (<<>>) {
    if (/\[/) {
        my @fields = unpack('A4 A4 A4 A4 A4 A4 A4 A4 A4');
        for my $i ( 0 .. $#fields ) {
            if ($fields[$i] ne "") {
                unshift @{$stacks[$i]}, "$fields[$i]";
                unshift @{$stacks2[$i]}, "$fields[$i]";
            }
        }
    }

    if (/move (\d+) from (\d+) to (\d+)/) {
        my $src = $2 - 1;
        my $dest = $3 - 1;
        for my $i (1..$1) {
            push (@{$stacks[$dest]}, pop(@{$stacks[$src]}));
        }

        my @tmp_stack = ();
        for my $i (1..$1) {
            unshift (@tmp_stack, pop(@{$stacks2[$src]}));
        }
        push (@{$stacks2[$dest]}, @tmp_stack);
    }

}

for my $i (0..$#stacks) {
    print $stacks[$i][-1];
}
print "\n";

for my $i (0..$#stacks) {
    print $stacks2[$i][-1];
}
print "\n";
