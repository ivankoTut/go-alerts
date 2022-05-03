package alerts

import (
	"fmt"
	"strings"
)

var colors = map[string]int{
	"black":   0,
	"red":     1,
	"green":   2,
	"yellow":  3,
	"blue":    4,
	"magenta": 5,
	"cyan":    6,
	"white":   7,
	"default": 9,
}

var brightColors = map[string]int{
	"gray":           0,
	"bright-red":     1,
	"bright-green":   2,
	"bright-yellow":  3,
	"bright-blue":    4,
	"bright-magenta": 5,
	"bright-cyan":    6,
	"bright-white":   7,
}

var availableOptions = map[string]map[string]int{
	"bold":       {"set": 1, "unset": 22},
	"underscore": {"set": 4, "unset": 24},
	"blink":      {"set": 5, "unset": 25},
	"reverse":    {"set": 7, "unset": 27},
	"conceal":    {"set": 8, "unset": 28},
}

type Color struct {
	foreground string
	background string
	options    map[string]map[string]int
}

func (c *Color) Apply(text string) string {

	return fmt.Sprintf("%s%s%s", c.set(), text, c.unset())
}

func (c *Color) set() string {
	var codes []string

	if c.foreground != "" {
		codes = append(codes, c.foreground)
	}

	if c.background != "" {
		codes = append(codes, c.background)
	}

	for _, val := range c.options {
		codes = append(codes, fmt.Sprintf("%d", val["set"]))
	}

	if len(codes) == 0 {
		return ""
	}

	return fmt.Sprintf("\033[%sm", strings.Join(codes, ";"))
}

func (c *Color) unset() string {
	var codes []string

	if c.foreground != "" {
		codes = append(codes, "39")
	}

	if c.background != "" {
		codes = append(codes, "49")
	}

	for _, val := range c.options {
		codes = append(codes, fmt.Sprintf("%d", val["unset"]))
	}

	if len(codes) == 0 {
		return ""
	}

	return fmt.Sprintf("\033[%sm", strings.Join(codes, ";"))
}

func CreateColor(foreground string, background string, options []string) (*Color, error) {
	color := Color{
		foreground: parseColor(foreground, false),
		background: parseColor(background, true),
	}
	color.options = make(map[string]map[string]int)

	for _, value := range options {
		if _, ok := availableOptions[value]; ok == false {
			return nil, fmt.Errorf("invalid option specified: %s", value)
		}

		color.options[value] = availableOptions[value]
	}

	return &color, nil
}

func parseColor(color string, isBackground bool) string {
	if color == "" {
		return ""
	}

	if color[0:1] == "#" {

	}

	if value, ok := colors[color]; ok {
		colorFirstCode := 3
		if isBackground {
			colorFirstCode = 4
		}
		return fmt.Sprintf("%d%d", colorFirstCode, value)
	}

	if value, ok := brightColors[color]; ok {
		colorFirstCode := 9
		if isBackground {
			colorFirstCode = 10
		}
		return fmt.Sprintf("%d%d", colorFirstCode, value)
	}

	return ""
}
