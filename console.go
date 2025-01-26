package console

import (
	"fmt"
	"github.com/rs/zerolog"
	"io"
	"os"
	"strconv"
	"time"
)

type Logger struct {
	logger *zerolog.Logger
	Level  zerolog.Level
}

func (l *Logger) Log(v any) {
	switch l.Level {
	case zerolog.InfoLevel:
		l.logger.Info().Msg(fmt.Sprint(v))
	case zerolog.DebugLevel:
		l.logger.Debug().Msg(fmt.Sprint(v))
	case zerolog.WarnLevel:
		l.logger.Warn().Msg(fmt.Sprint(v))
	case zerolog.ErrorLevel:
		l.logger.Error().Msg(fmt.Sprint(v))
	default:
	}
}

func (l *Logger) Logf(format string, v ...any) {
	switch l.Level {
	case zerolog.InfoLevel:
		l.logger.Info().Msgf(format, v...)
	case zerolog.DebugLevel:
		l.logger.Debug().Msgf(format, v...)
	case zerolog.WarnLevel:
		l.logger.Warn().Msgf(format, v...)
	case zerolog.ErrorLevel:
		l.logger.Error().Msgf(format, v...)
	default:
	}
}

func NewLogger(logger *zerolog.Logger, level zerolog.Level) *Logger {
	return &Logger{logger, level}
}

var defaultLogger = zerolog.New(os.Stdout).With().Caller().Timestamp().Logger()

var Info = NewLogger(&defaultLogger, zerolog.InfoLevel)
var Debug = NewLogger(&defaultLogger, zerolog.DebugLevel)
var Warn = NewLogger(&defaultLogger, zerolog.WarnLevel)
var Error = NewLogger(&defaultLogger, zerolog.ErrorLevel)

func Log() zerolog.Logger {
	return defaultLogger
}

func New(w io.Writer) zerolog.Logger {
	return zerolog.New(w)
}

func CallerMarshalFunc(fn func(pc uintptr, file string, line int) string) {
	zerolog.CallerMarshalFunc = fn
}

func CallerSkipFrameCount(count int) {
	zerolog.CallerSkipFrameCount = count
}

func TimeFieldFormat(format string) {
	zerolog.TimeFieldFormat = format
}

func TimestampFieldName(name string) {
	zerolog.TimestampFieldName = name
}

func CallerFieldName(name string) {
	zerolog.CallerFieldName = name
}

func InterfaceMarshalFunc(fn func(v interface{}) ([]byte, error)) {
	zerolog.InterfaceMarshalFunc = fn
}

func ErrorMarshalFunc(fn func(err error) interface{}) {
	zerolog.ErrorMarshalFunc = fn
}

func LevelFieldName(name string) {
	zerolog.LevelFieldName = name
}

func LevelFieldMarshalFunc(fn func(l zerolog.Level) string) {
	zerolog.LevelFieldMarshalFunc = fn
}

func MessageFieldName(name string) {
	zerolog.MessageFieldName = name
}

func LevelDebugValue(value string) {
	zerolog.LevelDebugValue = value
}

func LevelInfoValue(value string) {
	zerolog.LevelInfoValue = value
}

func LevelWarnValue(value string) {
	zerolog.LevelWarnValue = value
}

func LevelErrorValue(value string) {
	zerolog.LevelErrorValue = value
}

func init() {
	zerolog.LevelFieldName = "level"
	zerolog.CallerFieldName = "file"
	zerolog.TimeFieldFormat = time.DateTime
	zerolog.TimestampFieldName = "time"
	zerolog.CallerSkipFrameCount = 3
	zerolog.MessageFieldName = "log"
	zerolog.LevelDebugValue = "DEB"
	zerolog.LevelInfoValue = "INF"
	zerolog.LevelWarnValue = "WAR"
	zerolog.LevelErrorValue = "ERR"
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		var index = -1
		var count = 0
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == os.PathSeparator {
				count++
				if count == 2 {
					index = i
					break
				}
			}
		}

		if index == -1 {
			return file + ":" + strconv.Itoa(line)
		}
		return file[index+1:] + ":" + strconv.Itoa(line)
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
