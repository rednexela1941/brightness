package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const MAX_BRIGHTNESS_PATH = "/sys/class/backlight/intel_backlight/max_brightness"
const BRIGHTNESS_PATH = "/sys/class/backlight/intel_backlight/brightness"
const MAX_KEYBOARD_BRIGHTNESS_PATH = "/sys/class/leds/smc::kbd_backlight/max_brightness"
const KEYBOARD_BRIGHTNESS_PATH = "/sys/class/leds/smc::kbd_backlight/brightness"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getBrightnessValue(path string) int {
	file, err := os.Open(path)
	check(err)
	s := bufio.NewScanner(file)
	var line string
	for s.Scan() {
		line = s.Text()
		break
	}
	val, err := strconv.Atoi(line)
	check(err)
	return val
}

func getMax(path string) int {
	return getBrightnessValue(path)
}

func getCurrent(path string) int {
	return getBrightnessValue(path)
}

func getCurrentBrightness(path string, max_path string) (int, int) { // Return percentage, max.
	current_brightness := getCurrent(path)
	max_brightness := getMax(max_path)
	percent := int((float64(current_brightness) / float64(max_brightness)) * 100)
	return percent, max_brightness
}

func setBrightness(backlight bool, up bool) {
	var percent int
	var max int
	var path string
	var min float64

	if backlight {
		min = 20
		path = BRIGHTNESS_PATH
		percent, max = getCurrentBrightness(BRIGHTNESS_PATH, MAX_BRIGHTNESS_PATH)
	} else {
		min = 0
		path = KEYBOARD_BRIGHTNESS_PATH
		percent, max = getCurrentBrightness(KEYBOARD_BRIGHTNESS_PATH, MAX_KEYBOARD_BRIGHTNESS_PATH)
	}

	percent_increment := 8
	remainder := percent % percent_increment
	whole := percent / percent_increment
	new_val := percent_increment * (whole + int(math.Round(float64(remainder)/float64(percent_increment))))

	if up {
		new_val += percent_increment
	} else {
		new_val -= percent_increment
	}

	new_val = int(math.Max(min, float64(new_val))) // Set minimum brightness to 20%
	new_val = int(math.Min(100, float64(new_val)))
	new_percent := new_val
	new_val = int((float64(new_val) / 100) * float64(max))
	str_value := strconv.Itoa(new_val)

	file, err := os.OpenFile(path, os.O_WRONLY, 02)
	defer file.Close()
	_, err = file.WriteString(str_value)
	check(err)
	file.Sync()

	if backlight {
		fmt.Printf("Set brightness to: %d%%\n", new_percent)
	} else {
		fmt.Printf("Set keyboard brightness to: %d%%\n", new_percent)
	}
	return
}

func main() {
	args := os.Args[1:]
	percent, _ := getCurrentBrightness(BRIGHTNESS_PATH, MAX_BRIGHTNESS_PATH)
	percent_kbd, _ := getCurrentBrightness(KEYBOARD_BRIGHTNESS_PATH, MAX_KEYBOARD_BRIGHTNESS_PATH)
	fmt.Printf("Current brightness: %d%%\n", percent)
	fmt.Printf("Current keyboard brightness: %d%%\n", percent_kbd)
	if len(args) > 0 {
		cmd := args[0]
		if strings.ToUpper(cmd) == "UP" {
			setBrightness(true, true)
		} else if strings.ToUpper(cmd) == "DOWN" {
			setBrightness(true, false)
		} else if strings.ToUpper(cmd) == "--KBD" && len(args) > 1 {
			cmd2 := args[1]
			if strings.ToUpper(cmd2) == "UP" {
				setBrightness(false, true)
			} else if strings.ToUpper(cmd2) == "DOWN" {
				setBrightness(false, false)
			} else {
				fmt.Printf("Invalid command: %s %s\n", cmd, cmd2)
			}
		} else {
			fmt.Printf("Invalid command: %s\n", cmd)
		}
	} else {

		fmt.Println("No arguments supplied: up, or down.")
	}

}
