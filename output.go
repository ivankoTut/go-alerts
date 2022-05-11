package alerts

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"strings"
	"unicode/utf8"
)

const (
	paddingTop    = "padding-top"
	paddingBottom = "padding-bottom"
)

const DefaultPrefix = "  "

var maxLineLength int

func init() {
	weight, _, err := terminal.GetSize(0)
	maxLineLength = weight
	if err != nil {
		maxLineLength = 120
	}
}

func Success(message string) {
	color, _ := CreateColor("white", "green", []string{"bold"})

	CreateBlock(message, "OK", color)
}

func Warning(message string) {
	color, _ := CreateColor("black", "yellow", []string{"bold"})

	CreateBlock(message, "WARNING", color)
}

func Error(message string) {
	color, _ := CreateColor("white", "red", []string{"bold"})

	CreateBlock(message, "ERROR", color)
}

func Note(message string) {
	color, _ := CreateColor("yellow", "default", []string{"bold"})

	CreateBlock(message, "NOTE", color)
}

func CreateBlock(message string, name string, color *Color) {
	if name != "" {
		name = fmt.Sprintf(DefaultPrefix+"[%s]"+DefaultPrefix, name)
	} else {
		name = DefaultPrefix + name
	}

	indentLength := utf8.RuneCountInString(name)
	lineIndentation := strings.Repeat(" ", indentLength)

	if color.newLine {
		fmt.Println("")
	}

	createPaddingBlock(color, paddingTop)

	lines := strings.Split(WrapString(message, maxLineLength-indentLength), "\n")

	for i, value := range lines {
		if i == 0 {
			message = name + value + strings.Repeat(" ", maxLineLength-utf8.RuneCountInString(value+name))
			fmt.Println(color.Apply(message))
			continue
		}

		message = lineIndentation + value + strings.Repeat(" ", maxLineLength-utf8.RuneCountInString(value+lineIndentation))
		fmt.Println(color.Apply(message))
	}

	createPaddingBlock(color, paddingBottom)
}

func createPaddingBlock(color *Color, paddingMode string) {
	str := ""
	if paddingMode == paddingTop {
		if color.paddingTop == false {
			return
		}

		if color.paddingTopColor != "" {
			color.createColorMode = fillColorPaddingTop
			str = color.Apply(strings.Repeat(" ", maxLineLength))
			color.createColorMode = fillColorBody
		} else {
			str = color.Apply(strings.Repeat(" ", maxLineLength))
		}
	}

	if paddingMode == paddingBottom {
		if color.paddingBottom == false {
			return
		}

		if color.paddingBottomColor != "" {
			color.createColorMode = fillColorPaddingBottom
			str = color.Apply(strings.Repeat(" ", maxLineLength))
			color.createColorMode = fillColorBody
		} else {
			str = color.Apply(strings.Repeat(" ", maxLineLength))
		}
	}

	fmt.Println(str)
}
