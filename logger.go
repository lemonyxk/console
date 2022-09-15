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

	"github.com/lemonyxk/caller"
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
		Colorful:     true,
		Fields:       make(map[string]any),
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
	Colorful     bool
	Fields       map[string]any
}

func (log *Logger) Errorf(format string, args ...any) {
	var str = log.getLevelStringf(ERR, format, args...)

	if !log.Colorful {
		errWrite(str)
		return
	}

	errWrite(log.ErrorColor.Sprint(str))
}

func (log *Logger) Warningf(format string, args ...any) {
	var str = log.getLevelStringf(WAR, format, args...)

	if !log.Colorful {
		write(str)
		return
	}

	write(log.WarningColor.Sprint(str))
}

func (log *Logger) Infof(format string, args ...any) {
	var str = log.getLevelStringf(INF, format, args...)

	if !log.Colorful {
		write(str)
		return
	}

	write(log.InfoColor.Sprint(str))
}

func (log *Logger) Debugf(format string, args ...any) {
	var str = log.getLevelStringf(DEB, format, args...)

	if !log.Colorful {
		write(str)
		return
	}

	write(log.DebugColor.Sprint(str))
}

func (log *Logger) Error(args ...any) {
	var str = log.getLevelStringln(ERR, args...)

	if !log.Colorful {
		errWrite(str)
		return
	}

	errWrite(log.ErrorColor.Sprint(str))
}

func (log *Logger) Warning(args ...any) {
	var str = log.getLevelStringln(WAR, args...)

	if !log.Colorful {
		write(str)
		return
	}

	write(log.WarningColor.Sprint(str))
}

func (log *Logger) Info(args ...any) {
	var str = log.getLevelStringln(INF, args...)

	if !log.Colorful {
		write(str)
		return
	}

	write(log.InfoColor.Sprint(str))
}

func (log *Logger) Debug(args ...any) {
	var str = log.getLevelStringln(DEB, args...)

	if !log.Colorful {
		write(str)
		return
	}

	write(log.DebugColor.Sprint(str))
}

func (log *Logger) AddField(key string, value any) {
	log.Fields[key] = value
}

func (log *Logger) SetFields(fields map[string]any) {
	log.Fields = fields
}

func (log *Logger) getLevelStringf(level Level, format string, args ...any) string {
	return log.levelSprintf(level, format, args...)
}

func (log *Logger) getLevelStringln(level Level, args ...any) string {
	var msg = joinInterface(args, " ")
	return log.levelSprintf(level, "%s\n", msg)
}

func (log *Logger) levelSprintf(level Level, format string, args ...any) string {
	var entry = &Entry{}

	if log.Flags&LEVEL != 0 {
		entry.Level = level
	}

	if log.Flags&TIME != 0 {
		entry.Time = time.Now()
	}

	if log.Flags&FILE != 0 {
		file, line := caller.Auto(packageName)
		entry.File = file
		entry.Line = line
	}

	entry.Format = format
	entry.Fields = log.Fields
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
