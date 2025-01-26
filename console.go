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

func init() {
	zerolog.CallerFieldName = "file"
	zerolog.TimeFieldFormat = time.DateTime
	zerolog.TimestampFieldName = "time"
	zerolog.CallerSkipFrameCount = 3
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
