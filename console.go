package console

import (
	"fmt"
	"os"
)

var handlerLogger = NewLogger()

func SetLogger(logger *Logger) {
	handlerLogger = logger
}

func SetFormatter(formatter Formatter) {
	handlerLogger.Formatter = formatter
}

func SetFlags(flags int) {
	handlerLogger.Flags = flags
}

func SetHook(hook Hook) {
	handlerLogger.Hook = hook
}

func SetInfoColor(color Color) {
	handlerLogger.InfoColor = color
}

func SetWarningColor(color Color) {
	handlerLogger.WarningColor = color
}

func SetDebugColor(color Color) {
	handlerLogger.DebugColor = color
}

func SetErrorColor(color Color) {
	handlerLogger.ErrorColor = color
}

func Exit(v ...interface{}) {
	handlerLogger.Info(v...)
	os.Exit(0)
}

func Println(v ...interface{}) {
	fmt.Println(v...)
}

func Printf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}

func OneLine(format string, v ...interface{}) {
	fmt.Printf("\r"+format, v...)
}

func Info(v ...interface{}) {
	handlerLogger.Info(v...)
}

func Debug(v ...interface{}) {
	handlerLogger.Debug(v...)
}

func Warning(v ...interface{}) {
	handlerLogger.Warning(v...)
}

func Error(v ...interface{}) {
	handlerLogger.Error(v...)
}

func Infof(format string, v ...interface{}) {
	handlerLogger.Infof(format, v...)
}

func Warningf(format string, v ...interface{}) {
	handlerLogger.Warningf(format, v...)
}

func Debugf(format string, v ...interface{}) {
	handlerLogger.Debugf(format, v...)
}

func Errorf(format string, v ...interface{}) {
	handlerLogger.Errorf(format, v...)
}

func AssertError(v ...interface{}) {
	if len(v) == 0 {
		return
	}
	if isNil(v[len(v)-1]) {
		return
	}
	handlerLogger.Errorf("%v\n", v[len(v)-1])
}
