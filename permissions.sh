#!/bin/bash

# Set brightness permissions.
chmod a+w /sys/class/backlight/intel_backlight/brightness
chmod a+w /sys/class/leds/smc::kbd_backlight/brightness
