use v5.36;
use strict;
use warnings;

use constant WINDOW => 4;
use constant WINDOW2 => 14;

while (<>) {
    my $total_chars = WINDOW-1;
    my @chars = (split //);
    my $pos = 0;
    until ($pos + WINDOW > scalar(@chars)) {
        my %seen = ();
        for my $i ($pos..$pos+WINDOW-1) {
            $seen{$chars[$i]} = 1;
        }
        $total_chars++;
        last if keys %seen == WINDOW;
        undef %seen;
        $pos++;
    }
    say "start of packet: " . $total_chars;

    $total_chars = WINDOW2-1;
    $pos = 0;
    until ($pos + WINDOW2 > scalar(@chars)) {
        my %seen = ();
        for my $i ($pos..$pos+WINDOW2-1) {
            $seen{$chars[$i]} = 1;
        }
        $total_chars++;
        last if keys %seen == WINDOW2;
        undef %seen;
        $pos++;
    }
    say "start of message: " . $total_chars;
}
