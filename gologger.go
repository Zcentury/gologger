package gologger

import (
	"fmt"
	"github.com/Zcentury/gologger/formatter"
	"github.com/Zcentury/gologger/levels"
	"github.com/Zcentury/gologger/writer"
	"strings"
	"time"
)

var (
	labels = map[levels.Level]string{
		levels.LevelFatal:   "FTL",
		levels.LevelError:   "ERR",
		levels.LevelInfo:    "INF",
		levels.LevelWarning: "WRN",
		levels.LevelDebug:   "DBG",
	}

	LoggerOptions *Logger
)

func init() {
	LoggerOptions = &Logger{}
	LoggerOptions.SetFormatter(formatter.NewCLI(true))
	LoggerOptions.SetWriter(writer.NewCLI())
}

type Logger struct {
	writer    writer.Writer       // 输出接口
	formatter formatter.Formatter // 格式化接口
	timestamp bool                // 是否自动加入时间戳
}

func (l *Logger) Log(event *Event) {
	event.message = strings.TrimSuffix(event.message, "\n")
	data, err := l.formatter.Format(&formatter.LogEvent{
		Message:  event.message,
		Level:    event.level,
		Metadata: event.metadata,
	})
	if err != nil {
		return
	}
	l.writer.Write(data, event.level)
}

// SetFormatter 设置格式化方法
func (l *Logger) SetFormatter(formatter formatter.Formatter) {
	l.formatter = formatter
}

// SetWriter 设置输出方法
func (l *Logger) SetWriter(writer writer.Writer) {
	l.writer = writer
}

// SetTimestamp 是否自动添加时间戳
func (l *Logger) SetTimestamp(timestamp bool) {
	l.timestamp = timestamp
}

type Event struct {
	logger   *Logger
	level    levels.Level
	message  string
	metadata map[string]string
}

func newDefaultEventWithLevel(level levels.Level) *Event {
	return newEventWithLevelAndLogger(level, LoggerOptions)
}

func newEventWithLevelAndLogger(level levels.Level, l *Logger) *Event {
	event := &Event{
		logger:   l,
		level:    level,
		metadata: make(map[string]string),
	}
	if l.timestamp {
		event.TimeStamp()
	}
	return event
}

func (e *Event) setLevelMetadata(level levels.Level) {
	e.metadata["label"] = labels[level]
}

func (e *Event) Label(label string) *Event {
	e.metadata["label"] = label
	return e
}

func (e *Event) TimeStamp() *Event {
	e.metadata["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	return e
}

func (e *Event) Str(key, value string) *Event {
	e.metadata[key] = value
	return e
}

func (e *Event) Msg(message string) {
	e.message = message
	e.logger.Log(e)
}

func (e *Event) Msgf(format string, args ...interface{}) {
	e.message = fmt.Sprintf(format, args...)
	e.logger.Log(e)
}

func Fatal() *Event {
	event := newDefaultEventWithLevel(levels.LevelFatal)
	event.setLevelMetadata(levels.LevelFatal)
	return event
}

func Error() *Event {
	event := newDefaultEventWithLevel(levels.LevelError)
	event.setLevelMetadata(levels.LevelError)
	return event
}

func Info() *Event {
	event := newDefaultEventWithLevel(levels.LevelInfo)
	event.setLevelMetadata(levels.LevelInfo)
	return event
}

func Warning() *Event {
	event := newDefaultEventWithLevel(levels.LevelWarning)
	event.setLevelMetadata(levels.LevelWarning)
	return event
}

func Debug() *Event {
	event := newDefaultEventWithLevel(levels.LevelDebug)
	event.setLevelMetadata(levels.LevelDebug)
	return event
}
