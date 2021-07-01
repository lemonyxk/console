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

	"github.com/lemoyxk/caller"
)

const (
	NONE  = 0
	LEVEL = 1 << iota
	TIME
	FILE
	ID
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
	ID           string
}

func (log *Logger) GetLevelStringf(level Level, format string, args ...interface{}) string {
	if log.Flags&LEVEL == 0 {
		level = ""
	}
	return log.Sprintf(level, format, args...)
}

func (log *Logger) GetLevelStringln(level Level, args ...interface{}) string {
	if log.Flags&LEVEL == 0 {
		level = ""
	}
	var msg = joinInterface(args, " ")
	return log.Sprintf(level, "%s\n", msg)
}

func (log *Logger) Errorf(format string, args ...interface{}) {
	var str = log.GetLevelStringf(ERR, format, args...)

	if log.DisableColor {
		write(str)
		return
	}

	log.ErrorColor.Print(str)
}

func (log *Logger) Warningf(format string, args ...interface{}) {
	var str = log.GetLevelStringf(WAR, format, args...)

	if log.DisableColor {
		write(str)
		return
	}

	log.WarningColor.Print(str)
}

func (log *Logger) Infof(format string, args ...interface{}) {
	var str = log.GetLevelStringf(INF, format, args...)

	if log.DisableColor {
		write(str)
		return
	}

	log.InfoColor.Print(str)
}

func (log *Logger) Debugf(format string, args ...interface{}) {
	var str = log.GetLevelStringf(DEB, format, args...)

	if log.DisableColor {
		write(str)
		return
	}

	log.DebugColor.Print(str)
}

func (log *Logger) Error(args ...interface{}) {
	var str = log.GetLevelStringln(ERR, args...)

	if log.DisableColor {
		write(str)
		return
	}

	log.ErrorColor.Print(str)
}

func (log *Logger) Warning(args ...interface{}) {
	var str = log.GetLevelStringln(WAR, args...)

	if log.DisableColor {
		write(str)
		return
	}

	log.WarningColor.Print(str)
}

func (log *Logger) Info(args ...interface{}) {
	var str = log.GetLevelStringln(INF, args...)

	if log.DisableColor {
		write(str)
		return
	}

	log.InfoColor.Print(str)
}

func (log *Logger) Debug(args ...interface{}) {
	var str = log.GetLevelStringln(DEB, args...)

	if log.DisableColor {
		write(str)
		return
	}

	log.DebugColor.Print(str)
}

func (log *Logger) Sprintf(level Level, format string, args ...interface{}) string {
	var entry = &Entry{Level: level}

	if log.Flags&TIME != 0 {
		entry.Time = time.Now()
	}

	if log.Flags&ID != 0 {
		entry.ID = log.ID
	}

	if log.Flags&FILE != 0 {
		file, line := caller.Auto(packageName)

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
