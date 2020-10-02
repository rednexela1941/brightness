### Commands

* `./main [up|down]` Increase|Decrease screen brightness
* `./main --kbd [up|down]` Increase|Decrease keyboard brightness

### Note
* `chmod a+w` both of the following paths: 

BRIGHTNESS_PATH = "/sys/class/backlight/intel_backlight/brightness"
KEYBOARD_BRIGHTNESS_PATH = "/sys/class/leds/smc::kbd_backlight/brightness"

### Install
* `sudo GOBIN=/usr/local/bin/ go install src/brightness.go`

## i3 Plugin
* Add the following lines to i3 configuration.

```
# Screen brightness controls
bindsym XF86MonBrightnessUp exec brightness up		#increase screen brightness
bindsym XF86MonBrightnessDown exec brightness down 	#decrease screen brightness

# Keyboard backlight controls
bindsym XF86KbdBrightnessUp exec brightness --kbd up		#increase Kbd brightness
bindsym XF86KbdBrightnessDown exec brightness --kbd down	#decrease Kbd brightness
```
