## Installation and Usage

Install using `go get github.com/ivankoTut/go-alerts`.

```go
alerts.Success("Test message")
alerts.Warning("Test message")
alerts.Error("Test message")
alerts.Note("Test message")
```

### Customize
```go
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
```

![alt text](https://github.com/ivankoTut/go-alerts/blob/master/example/1651582978036.jpg?raw=true)