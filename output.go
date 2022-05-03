package alerts

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"strings"
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

func Success(message string) error {
	color, err := CreateColor("white", "green", []string{"bold"})

	if err != nil {
		return err
	}

	createBlock(message, "OK", color)

	return nil
}

func Warning(message string) error {
	color, err := CreateColor("black", "yellow", []string{"bold"})

	if err != nil {
		return err
	}

	createBlock(message, "WARNING", color)

	return nil
}

func Error(message string) error {
	color, err := CreateColor("white", "red", []string{"bold"})

	if err != nil {
		return err
	}

	createBlock(message, "ERROR", color)

	return nil
}

func Note(message string) error {
	color, err := CreateColor("yellow", "default", []string{"bold"})

	if err != nil {
		return err
	}

	createBlock(message, "NOTE", color)

	return nil
}

func createBlock(message string, name string, color *Color) {
	name = fmt.Sprintf(DefaultPrefix+"[%s]"+DefaultPrefix, name)
	indentLength := len(name)
	lineIndentation := strings.Repeat(" ", indentLength)

	fmt.Println("")
	fmt.Println(color.Apply(strings.Repeat(" ", maxLineLength)))
	lines := strings.Split(WrapString(message, maxLineLength-indentLength), "\n")

	for i, value := range lines {
		if i == 0 {
			message = name + value + strings.Repeat(" ", maxLineLength-len(value+name))
			fmt.Println(color.Apply(message))
			continue
		}

		message = lineIndentation + value + strings.Repeat(" ", maxLineLength-len(value+lineIndentation))
		fmt.Println(color.Apply(message))
	}

	fmt.Println(color.Apply(strings.Repeat(" ", maxLineLength)))
}
