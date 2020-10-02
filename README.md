### Commands

* `./main [up|down]` Increase|Decrease screen brightness
* `./main --kbd [up|down]` Increase|Decrease keyboard brightness

### Note
* `chmod a+w` both of the following paths: 
BRIGHTNESS_PATH = "/sys/class/backlight/intel_backlight/brightness"
KEYBOARD_BRIGHTNESS_PATH = "/sys/class/leds/smc::kbd_backlight/brightness"

### Install
* `sudo GOBIN=/usr/local/bin/ go install src/brightness.go`



