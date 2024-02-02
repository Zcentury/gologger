package formatter

import (
	"bytes"
	"github.com/Zcentury/gologger/levels"
	"github.com/logrusorgru/aurora/v4"
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

	if label, ok := event.Metadata["label"]; label != "" && ok {
		buffer.WriteRune('[')
		buffer.WriteString(label)
		buffer.WriteRune(']')
		buffer.WriteRune(' ')
		delete(event.Metadata, "label")
	}

	if timestamp, ok := event.Metadata["timestamp"]; timestamp != "" && ok {
		buffer.WriteRune('[')
		buffer.WriteString(aurora.Bold(aurora.Green(timestamp)).String())
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
		event.Metadata["label"] = aurora.Bold(aurora.BgGreen(" " + label + " ")).String()
	case levels.LevelFatal:
		event.Metadata["label"] = aurora.Bold(aurora.BgMagenta(" " + label + " ")).String()
	case levels.LevelError:
		event.Metadata["label"] = aurora.Bold(aurora.BgRed(" " + label + " ")).String()
	case levels.LevelDebug:
		event.Metadata["label"] = aurora.Bold(aurora.BgBlue(" " + label + " ")).String()
	case levels.LevelWarning:
		event.Metadata["label"] = aurora.Bold(aurora.BgYellow(" " + label + " ")).String()
	}
}
