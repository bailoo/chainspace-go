package log

import (
	"fmt"
	"strconv"
)

// Logging levels.
const (
	DebugLevel level = iota + 1
	InfoLevel
	ErrorLevel
	FatalLevel
)

const maxLevel = FatalLevel + 1

const (
	blue   = 34
	red    = 31
	yellow = 33
)

var (
	minLevel     = maxLevel
	consoleLevel = maxLevel
	fileLevel    = maxLevel
	netLevel     = maxLevel
	textLevel    = maxLevel
)

var (
	level2color  = [...]string{"", color(yellow, "DEBUG"), color(blue, "INFO"), color(red, "ERROR"), color(red, "FATAL")}
	level2string = [...]string{"", rpad("DEBUG"), rpad("INFO"), rpad("ERROR"), rpad("FATAL")}
)

type level int8

func (l level) String() string {
	switch l {
	case DebugLevel:
		return "DEBUG"
	case ErrorLevel:
		return "ERROR"
	case FatalLevel:
		return "FATAL"
	case InfoLevel:
		return "INFO"
	default:
		panic(fmt.Errorf("log: unknown level: %d", l))
	}
}

func color(code int64, text string) string {
	return "\x1b[" + strconv.FormatInt(code, 10) + "m" + rpad(text) + "\x1b[0m"
}

func rpad(text string) string {
	for i := len(text); i < 8; i++ {
		text += " "
	}
	return text
}

func updateMinLevels() {
	minLevel = maxLevel
	textLevel = maxLevel
	if consoleLevel < minLevel {
		minLevel = consoleLevel
		textLevel = consoleLevel
	}
	if fileLevel < minLevel {
		minLevel = fileLevel
		textLevel = fileLevel
	}
	if netLevel < minLevel {
		minLevel = netLevel
	}
}

// AtDebug returns whether the current configuration will log at the DebugLevel.
func AtDebug() bool {
	return minLevel <= DebugLevel
}

// AtError returns whether the current configuration will log at the ErrorLevel.
func AtError() bool {
	return minLevel <= ErrorLevel
}

// AtInfo returns whether the current configuration will log at the InfoLevel.
func AtInfo() bool {
	return minLevel <= InfoLevel
}