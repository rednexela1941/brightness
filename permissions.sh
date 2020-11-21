#!/bin/bash

# Set brightness permissions.
chmod a+w /sys/class/backlight/gmux_backlight/brightness 
chmod a+w /sys/class/leds/spi::kbd_backlight/brightness 
