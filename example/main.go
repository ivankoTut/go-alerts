package main

import (
	"github.com/ivankoTut/go-alerts"
)

func main() {

	alerts.Success("test")
	alerts.Error("This is a long text! This is a long text! This is a long text! This is a long text! " +
		"This is a long text! This is a long text! This is a long text! This is a long text! This is a long text! " +
		"This is a long text! This is a long text! This is a long text!")

	// custom alerts
	color, err := alerts.CreateColor("default", "default", []string{"bold"})

	if err != nil {
		panic(err)
	}

	color.
		PrintPaddingBottom(true).
		PrintPaddingTop(true).
		PrintNewLine(true).
		SetPaddingTopColor("black").
		SetPaddingBottomColor("black")

	alerts.CreateBlock("", "", color) // without title and text

	alerts.CreateBlock("Text", "TITLE", color) // with title
}
