package levels

type Level int

const (
	// LevelFatal 严重
	LevelFatal Level = iota
	// LevelError 错误
	LevelError
	// LevelInfo 消息
	LevelInfo
	// LevelWarning 警告
	LevelWarning
	// LevelDebug 调试
	LevelDebug
)

func (l Level) String() string {
	return [...]string{"fatal", "error", "info", "warning", "debug"}[l]
}
