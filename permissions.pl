#!/usr/bin/perl
use strict;
use warnings;
use v5.30.0;
use Cwd qw( abs_path getcwd );
use File::Basename qw( dirname );
use experimental qw( switch );
use Term::ANSIColor;

sub main {
    my $backlight_file = "/sys/class/backlight/intel_backlight/brightness";
    my $keyboard_file  = "/sys/class/leds/smc::kbd_backlight/brightness";

    my $fail_count = 1;

    while ( $fail_count < 10 ) {
        if ( -e $backlight_file ) {
            if ( system("chmod a+w $backlight_file") != 0 ) {
                print("Failed to chmod $backlight_file\n");
                $fail_count++;
            } else {
                last;
            }
        }
		sleep(1);
    }
    $fail_count = 0;
    while ( $fail_count < 10 ) {
        if ( -e $keyboard_file ) {
            if ( system("chmod a+w $keyboard_file") != 0 ) {
                print("Failed to chmod $keyboard_file\n");
                $fail_count++;
            } else {
                last;
            }
        }
		sleep(1);
    }
    print("Done\n");
}

main();
