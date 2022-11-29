package utils

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

type LogLevel int

const (
	INFO LogLevel = iota
	ERROR
	DEBUG
	WARNNING
)

type Logger struct {
	LogLevel    LogLevel
	Message     string
	Description string
	Error       error
}

func getTime() string {
	currentTime := time.Now()
	return currentTime.Format("2006-01-02 15:04:05")
}

func getError(err error, typ string) string {
	if err != nil {
		if typ == "error" {
			return "\n Error: " + err.Error()
		} else if typ == "warning" {
			return "\n Warning: " + err.Error()
		}
	}
	return ""
}

func (l *Logger) Log() {
	switch l.LogLevel {
	case INFO:
		message := fmt.Sprintf("[INFO] [%s] %s - %s", getTime(), l.Message, l.Description)
		color.Blue(message)
	case ERROR:
		message := fmt.Sprintf("[ERROR] [%s] %s - %s %s", getTime(), l.Message, l.Description, getError(l.Error, "error"))
		color.Red(message)
	case DEBUG:
		message := fmt.Sprintf("[DEBUG] [%s] %s - %s", getTime(), l.Message, l.Description)
		color.White(message)
	case WARNNING:
		message := fmt.Sprintf("[INFO] [%s] %s - %s %s", getTime(), l.Message, l.Description, getError(l.Error, "warning"))
		color.Yellow(message)
	default:
		message := fmt.Sprintf("[INFO] [%s] %s - %s", getTime(), l.Message, l.Description)
		color.Blue(message)
	}
}
