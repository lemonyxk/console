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
	"io"
	"os"
	"sync"
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
		Hook:         nil,
		Formatter:    NewTextFormatter(),
		Flags:        LEVEL | TIME | FILE,
		InfoColor:    Bold,
		WarningColor: FgYellow,
		DebugColor:   FgGreen,
		ErrorColor:   FgRed,
		Colorful:     false,
		Fields:       make(map[string]any),
		Stdout:       os.Stdout,
		Stderr:       os.Stderr,
		Deep:         4,
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
	Stdout       io.Writer
	Stderr       io.Writer
	Deep         int

	mux sync.Mutex
}

func (l *Logger) Errorf(format string, args ...any) {
	var str = l.getLevelStringf(ERR, format, args...)

	if l.Colorful {
		str = l.ErrorColor.Sprint(str)
	}

	l.stderr(str)
}

func (l *Logger) Warningf(format string, args ...any) {
	var str = l.getLevelStringf(WAR, format, args...)

	if l.Colorful {
		str = l.WarningColor.Sprint(str)
	}

	l.stdout(str)
}

func (l *Logger) Infof(format string, args ...any) {
	var str = l.getLevelStringf(INF, format, args...)

	if l.Colorful {
		str = l.InfoColor.Sprint(str)
	}

	l.stdout(str)
}

func (l *Logger) Debugf(format string, args ...any) {
	var str = l.getLevelStringf(DEB, format, args...)

	if l.Colorful {
		str = l.DebugColor.Sprint(str)
	}

	l.stdout(str)
}

func (l *Logger) Error(args ...any) {
	var str = l.getLevelStringln(ERR, args...)

	if l.Colorful {
		str = l.ErrorColor.Sprint(str)
	}

	l.stderr(str)
}

func (l *Logger) Warning(args ...any) {
	var str = l.getLevelStringln(WAR, args...)

	if l.Colorful {
		str = l.WarningColor.Sprint(str)
	}

	l.stdout(str)
}

func (l *Logger) Info(args ...any) {
	var str = l.getLevelStringln(INF, args...)

	if l.Colorful {
		str = l.InfoColor.Sprint(str)
	}

	l.stdout(str)
}

func (l *Logger) Debug(args ...any) {
	var str = l.getLevelStringln(DEB, args...)

	if l.Colorful {
		str = l.DebugColor.Sprint(str)
	}

	l.stdout(str)
}

func (l *Logger) AddField(key string, value any) {
	l.Fields[key] = value
}

func (l *Logger) SetFields(fields map[string]any) {
	l.Fields = fields
}

func (l *Logger) getLevelStringf(level Level, format string, args ...any) string {
	return l.levelSprintf(level, format, args...)
}

func (l *Logger) getLevelStringln(level Level, args ...any) string {
	var msg = joinInterface(args, " ")
	return l.levelSprintf(level, "%s\n", msg)
}

func (l *Logger) levelSprintf(level Level, format string, args ...any) string {
	var entry = &Entry{}

	if l.Flags&LEVEL != 0 {
		entry.Level = level
	}

	if l.Flags&TIME != 0 {
		entry.Time = time.Now()
	}

	if l.Flags&FILE != 0 {
		info := caller.Deep(l.Deep)
		entry.File = info.File
		entry.Line = info.Line
	}

	entry.Format = format
	entry.Fields = l.Fields
	entry.Args = args

	var str string
	if l.Formatter == nil {
		str = fmt.Sprintf(format, args...)
	} else {
		str = l.Formatter.Format(entry)
	}

	if l.Hook != nil {
		l.Hook.Fire(entry)
	}

	return str
}

func (l *Logger) stdout(str string) {
	l.mux.Lock()
	defer l.mux.Unlock()
	panicIfError(l.Stdout.Write([]byte(str)))
}

func (l *Logger) stderr(str string) {
	l.mux.Lock()
	defer l.mux.Unlock()
	panicIfError(l.Stderr.Write([]byte(str)))
}
