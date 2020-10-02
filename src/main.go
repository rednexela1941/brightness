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

func getMax() int {
	return getBrightnessValue(MAX_BRIGHTNESS_PATH)
}

func getCurrent() int {
	return getBrightnessValue(BRIGHTNESS_PATH)
}

func setBrightness(max int, percent int, up bool) {
	percent_increment := 8
	remainder := percent % percent_increment
	whole := percent / percent_increment
	new_val := percent_increment * (whole + int(math.Round(float64(remainder)/float64(percent_increment))))

	if up {
		new_val += percent_increment
	} else {
		new_val -= percent_increment
	}

	new_val = int(math.Max(20, float64(new_val))) // Set minimum brightness to 20%
	new_val = int(math.Min(100, float64(new_val)))
	new_percent := new_val
	new_val = int((float64(new_val) / 100) * float64(max))

	str_value := strconv.Itoa(new_val)
	file, err := os.OpenFile(BRIGHTNESS_PATH, os.O_WRONLY, 02)
	defer file.Close()
	_, err = file.WriteString(str_value)
	check(err)
	file.Sync()
	fmt.Printf("Set brightness to: %d%%\n", new_percent)
	return
}

func main() {
	args := os.Args[1:]
	current_brightness := getCurrent()
	max_brightness := getMax()
	percent := int((float64(current_brightness) / float64(max_brightness)) * 100)
	fmt.Printf("Current brightness: %d%%\n", percent)
	if len(args) > 0 {
		cmd := args[0]
		if strings.ToUpper(cmd) == "UP" {
			setBrightness(max_brightness, percent, true)
		} else if strings.ToUpper(cmd) == "DOWN" {
			setBrightness(max_brightness, percent, false)
		} else {
			fmt.Printf("Invalid command: %s\n", cmd)
		}
	} else {

		fmt.Println("No arguments supplied: up, or down.")
	}

}
