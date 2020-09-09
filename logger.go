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
	return &Logger{
		Formatter:    NewTextFormatter(),
		Flags:        LEVEL | TIME | FILE,
		InfoColor:    Bold,
		WarningColor: FgYellow,
		DebugColor:   FgGreen,
		ErrorColor:   FgRed,
		DisableColor: false,
	}
}

type Logger struct {
	Hook         Hook
	Formatter    Formatter
	Flags        int
	InfoColor    Color
	WarningColor Color
	DebugColor   Color
	ErrorColor   Color
	DisableColor bool
}

func (log *Logger) GetLevelStringf(level Level, format string, args ...interface{}) string {
	if log.Flags&LEVEL == 0 {
		level = ""
	}
	return log.Sprintf(level, format, args...)
}

func (log *Logger) GetLevelStringln(level Level, format string, args ...interface{}) string {
	if log.Flags&LEVEL == 0 {
		level = ""
	}
	var msg = joinInterface(args, " ")
	return log.Sprintf(level, "%s", msg)
}

func (log *Logger) Errorf(format string, args ...interface{}) {
	var str = log.GetLevelStringf(ERR, format, args...)

	if log.DisableColor {
		fmt.Printf("%s", str)
		return
	}

	log.ErrorColor.Printf("%s", str)
}

func (log *Logger) Warningf(format string, args ...interface{}) {
	var str = log.GetLevelStringf(WAR, format, args...)

	if log.DisableColor {
		fmt.Printf("%s", str)
		return
	}

	log.WarningColor.Printf("%s", str)
}

func (log *Logger) Infof(format string, args ...interface{}) {
	var str = log.GetLevelStringf(INF, format, args...)

	if log.DisableColor {
		fmt.Printf("%s", str)
		return
	}

	log.InfoColor.Printf("%s", str)
}

func (log *Logger) Debugf(format string, args ...interface{}) {
	var str = log.GetLevelStringf(DEB, format, args...)

	if log.DisableColor {
		fmt.Printf("%s", str)
		return
	}

	log.DebugColor.Printf("%s", str)
}

func (log *Logger) Error(args ...interface{}) {
	var str = log.GetLevelStringln(ERR, "%s", args...)

	if log.DisableColor {
		fmt.Printf("%s\n", str)
		return
	}

	log.ErrorColor.Printf("%s\n", str)
}

func (log *Logger) Warning(args ...interface{}) {
	var str = log.GetLevelStringln(WAR, "%s", args...)

	if log.DisableColor {
		fmt.Printf("%s\n", str)
		return
	}

	log.WarningColor.Printf("%s\n", str)
}

func (log *Logger) Info(args ...interface{}) {
	var str = log.GetLevelStringln(INF, "%s", args...)

	if log.DisableColor {
		fmt.Printf("%s\n", str)
		return
	}

	log.InfoColor.Printf("%s\n", str)
}

func (log *Logger) Debug(args ...interface{}) {
	var str = log.GetLevelStringln(DEB, "%s", args...)

	if log.DisableColor {
		fmt.Printf("%s\n", str)
		return
	}

	log.DebugColor.Printf("%s\n", str)
}

func (log *Logger) Sprintf(level Level, format string, args ...interface{}) string {
	var entry = &Entry{Level: level}

	if log.Flags&TIME != 0 {
		entry.Time = time.Now()
	}

	if log.Flags&FILE != 0 {
		file, line := caller(4)
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

	if log.Hook != nil {
		log.Hook.Fire(entry)
	}

	return str
}
