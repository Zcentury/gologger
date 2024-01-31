package formatter

import (
	"bytes"
	"github.com/Zcentury/gologger/levels"
	"github.com/fatih/color"
)

type CLI struct {
	UseColors bool
}

func NewCLI(useColors bool) *CLI {
	return &CLI{
		UseColors: useColors,
	}
}

func (c *CLI) Format(event *LogEvent) ([]byte, error) {
	c.colorizeLabel(event)

	buffer := &bytes.Buffer{}
	buffer.Grow(len(event.Message))

	label, ok := event.Metadata["label"]
	if label != "" && ok {
		buffer.WriteRune('[')
		buffer.WriteString(label)
		buffer.WriteRune(']')
		buffer.WriteRune(' ')
		delete(event.Metadata, "label")
	}
	timestamp, ok := event.Metadata["timestamp"]
	if timestamp != "" && ok {
		buffer.WriteRune('[')
		buffer.WriteString(color.GreenString(timestamp))
		buffer.WriteRune(']')
		buffer.WriteRune(' ')
		delete(event.Metadata, "timestamp")
	}
	buffer.WriteString(event.Message)

	for k, v := range event.Metadata {
		buffer.WriteRune(' ')
		buffer.WriteString(k)
		buffer.WriteRune('=')
		buffer.WriteString(v)
	}
	data := buffer.Bytes()
	return data, nil
}

func (c *CLI) colorizeLabel(event *LogEvent) {
	label := event.Metadata["label"]
	if label == "" || !c.UseColors {
		return
	}
	switch event.Level {
	case levels.LevelInfo:
		event.Metadata["label"] = color.New(color.BgGreen, color.FgWhite).Sprint(" " + label + " ")
	case levels.LevelFatal:
		event.Metadata["label"] = color.New(color.BgMagenta, color.FgWhite).Sprint(" " + label + " ")
	case levels.LevelError:
		event.Metadata["label"] = color.New(color.BgRed, color.FgWhite).Sprint(" " + label + " ")
	case levels.LevelDebug:
		event.Metadata["label"] = color.New(color.BgBlue, color.FgWhite).Sprint(" " + label + " ")
	case levels.LevelWarning:
		event.Metadata["label"] = color.New(color.BgYellow, color.FgWhite).Sprint(" " + label + " ")
	}
}
