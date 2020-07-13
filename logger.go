/**
* @program: lemo
*
* @description:
*
* @author: lemo
*
* @create: 2020-07-11 18:20
**/

package console

import (
	"fmt"
	"time"
)

const (
	NONE  = 0
	LEVEL = 1 << iota
	TIME
	FILE
)

func NewLogger() *Logger {
	return &Logger{Formatter: NewTextFormatter(), Flags: LEVEL | TIME | FILE, Color: true}
}

type Logger struct {
	Hook      Hook
	Formatter Formatter
	Flags     int
	Color     bool
}

func (log *Logger) SetHook(hook Hook) {
	log.Hook = hook
}

func (log *Logger) SetColor(flag bool) {
	log.Color = flag
}

func (log *Logger) SetFormatter(formatter Formatter) {
	log.Formatter = formatter
}

func (log *Logger) SetFlags(flags int) {
	log.Flags = flags
}

func (log *Logger) Errorf(format string, args ...interface{}) {
	var level Level
	var color Color
	if log.Flags&LEVEL != 0 {
		level = ERR
		color = FgRed
	}
	log.doPrint(level, color, false, format, args...)
}

func (log *Logger) Warningf(format string, args ...interface{}) {
	var level Level
	var color Color
	if log.Flags&LEVEL != 0 {
		level = WAR
		color = FgYellow
	}
	log.doPrint(level, color, false, format, args...)
}

func (log *Logger) Infof(format string, args ...interface{}) {
	var level Level
	var color Color
	if log.Flags&LEVEL != 0 {
		level = INF
		color = Bold
	}
	log.doPrint(level, color, false, format, args...)
}

func (log *Logger) Debugf(format string, args ...interface{}) {
	var level Level
	var color Color
	if log.Flags&LEVEL != 0 {
		level = DEB
		color = FgBlue
	}
	log.doPrint(level, color, false, format, args...)
}

func (log *Logger) Error(args ...interface{}) {
	var level Level
	var color Color
	if log.Flags&LEVEL != 0 {
		level = ERR
		color = FgRed
	}
	var msg = joinInterface(args, " ")
	log.doPrint(level, color, true, "%s", msg)
}

func (log *Logger) Warning(args ...interface{}) {
	var level Level
	var color Color
	if log.Flags&LEVEL != 0 {
		level = WAR
		color = FgYellow
	}
	var msg = joinInterface(args, " ")
	log.doPrint(level, color, true, "%s", msg)
}

func (log *Logger) Info(args ...interface{}) {
	var level Level
	var color Color
	if log.Flags&LEVEL != 0 {
		level = INF
		color = Bold
	}
	var msg = joinInterface(args, " ")
	log.doPrint(level, color, true, "%s", msg)
}

func (log *Logger) Debug(args ...interface{}) {
	var level Level
	var color Color
	if log.Flags&LEVEL != 0 {
		level = DEB
		color = FgBlue
	}
	var msg = joinInterface(args, " ")
	log.doPrint(level, color, true, "%s", msg)
}

func (log *Logger) doPrint(level Level, color Color, ln bool, format string, args ...interface{}) {
	var entry = &Entry{Level: level}

	if log.Flags&TIME != 0 {
		entry.Time = time.Now()
	}

	if log.Flags&FILE != 0 {
		file, line := caller(3)
		entry.File = file
		entry.Line = line
	}

	entry.Format = format
	entry.Args = args

	var str string
	if log.Formatter == nil {
		str = fmt.Sprintf(format, args...)
	} else {
		str = log.Formatter.Format(entry)
	}

	if !log.Color {
		if ln {
			fmt.Printf("%s\n", str)
		} else {
			fmt.Printf("%s", str)
		}
	} else {
		if ln {
			color.Printf("%s\n", str)
		} else {
			color.Printf("%s", str)
		}
	}

	if log.Hook != nil {
		log.Hook.Fire(entry)
	}
}
