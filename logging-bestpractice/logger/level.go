package logger

import (
	"fmt"
	"strings"
)

// Level is ...
type Level int8

type levels struct {
	Debug Level
	Info  Level
	Warn  Level
	Error Level
	Panic Level
	Fatal Level
}

// String is ...
func (l Level) String() string {
	switch l {
	case Levels.Debug:
		return "debug"
	case Levels.Info:
		return "info"
	case Levels.Warn:
		return "warn"
	case Levels.Error:
		return "error"
	case Levels.Panic:
		return "panic"
	case Levels.Fatal:
		return "fatal"
	default:
		return "unknown"
	}
}

// Levels is ...
var Levels = func() levels {
	const (
		Debug Level = iota + 1
		Info
		Warn
		Error
		Panic
		Fatal
	)
	return levels{
		Debug: Debug,
		Info:  Info,
		Warn:  Warn,
		Error: Error,
		Panic: Panic,
		Fatal: Fatal,
	}
}()

// ParseLevel is ...
func ParseLevel(level string) Level {
	switch strings.ToLower(level) {
	case "debug":
		return Levels.Debug
	case "info":
		return Levels.Info
	case "warn":
		return Levels.Warn
	case "error":
		return Levels.Error
	case "panic":
		return Levels.Panic
	case "fatal":
		return Levels.Fatal
	default:
		panic(fmt.Sprintf("unknown log level: %s", level))
	}
}
