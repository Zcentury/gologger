package formatter

import "github.com/Zcentury/gologger/levels"

type Formatter interface {
	Format(event *LogEvent) ([]byte, error)
}

type LogEvent struct {
	Message  string
	Level    levels.Level
	Metadata map[string]string
}
