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

	LoggerOptions *logger
)

func init() {
	LoggerOptions = &logger{}
	LoggerOptions.SetFormatter(formatter.NewCLI(true))
	LoggerOptions.SetWriter(writer.NewCLI())
}

type logger struct {
	writer    writer.Writer       // 输出接口
	formatter formatter.Formatter // 格式化接口
	timestamp bool                // 是否自动加入时间戳
}

func (l *logger) Log(event *event) {
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
func (l *logger) SetFormatter(formatter formatter.Formatter) {
	l.formatter = formatter
}

// SetWriter 设置输出方法
func (l *logger) SetWriter(writer writer.Writer) {
	l.writer = writer
}

// SetTimestamp 是否自动添加时间戳
func (l *logger) SetTimestamp(timestamp bool) {
	l.timestamp = timestamp
}

type event struct {
	logger   *logger
	level    levels.Level
	message  string
	metadata map[string]string
}

func newDefaultEventWithLevel(level levels.Level) *event {
	return newEventWithLevelAndLogger(level, LoggerOptions)
}

func newEventWithLevelAndLogger(level levels.Level, l *logger) *event {
	event := &event{
		logger:   l,
		level:    level,
		metadata: make(map[string]string),
	}
	if l.timestamp {
		event.TimeStamp()
	}
	return event
}

func (e *event) setLevelMetadata(level levels.Level) {
	e.metadata["label"] = labels[level]
}

func (e *event) Label(label string) *event {
	e.metadata["label"] = label
	return e
}

func (e *event) TimeStamp() *event {
	e.metadata["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	return e
}

func (e *event) Str(key, value string) *event {
	e.metadata[key] = value
	return e
}

func (e *event) Msg(message string) {
	e.message = message
	e.logger.Log(e)
}

func (e *event) Msgf(format string, args ...interface{}) {
	e.message = fmt.Sprintf(format, args...)
	e.logger.Log(e)
}

func Fatal() *event {
	event := newDefaultEventWithLevel(levels.LevelFatal)
	event.setLevelMetadata(levels.LevelFatal)
	return event
}

func Error() *event {
	event := newDefaultEventWithLevel(levels.LevelError)
	event.setLevelMetadata(levels.LevelError)
	return event
}

func Info() *event {
	event := newDefaultEventWithLevel(levels.LevelInfo)
	event.setLevelMetadata(levels.LevelInfo)
	return event
}

func Warning() *event {
	event := newDefaultEventWithLevel(levels.LevelWarning)
	event.setLevelMetadata(levels.LevelWarning)
	return event
}

func Debug() *event {
	event := newDefaultEventWithLevel(levels.LevelDebug)
	event.setLevelMetadata(levels.LevelDebug)
	return event
}
