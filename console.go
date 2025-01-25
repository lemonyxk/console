package console

import (
	"fmt"
	"github.com/rs/zerolog"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

type Logger struct {
	*zerolog.Event
}

func (l *Logger) Log(msg ...any) {
	l.Event.Msg(fmt.Sprint(msg...))
}

func (l *Logger) Logf(format string, v ...any) {
	l.Event.Msgf(format, v...)
}

func NewLogger(event *zerolog.Event) *Logger {
	return &Logger{Event: event}
}

var defaultLogger = zerolog.New(os.Stdout).With().Caller().Logger()

var Info = NewLogger(defaultLogger.Info())
var Debug = NewLogger(defaultLogger.Debug())
var Warn = NewLogger(defaultLogger.Warn())
var Error = NewLogger(defaultLogger.Error())

func New(w io.Writer) zerolog.Logger {
	return zerolog.New(w)
}

func init() {
	zerolog.CallerFieldName = "file"
	zerolog.TimeFieldFormat = time.DateTime
	zerolog.TimestampFieldName = "time"
	zerolog.CallerSkipFrameCount = 3
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		var paths = strings.Split(file, string(os.PathSeparator))
		if len(paths) > 3 {
			file = strings.Join(paths[len(paths)-3:], string(os.PathSeparator))
		}
		return file + ":" + strconv.Itoa(line)
	}
}

func Exit(v ...any) {
	fmt.Println(v...)
	os.Exit(0)
}

func Println(v ...any) {
	fmt.Println(v...)
}

func Printf(format string, v ...any) {
	fmt.Printf(format, v...)
}

func OneLine(format string, v ...any) {
	fmt.Printf("\r"+format, v...)
}
