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

func SetFields(fields map[string]any) {
	handlerLogger.Fields = fields
}

func AddField(key string, value any) {
	handlerLogger.Fields[key] = value
}

func Colorful(flag bool) {
	handlerLogger.Colorful = flag
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

func Exit(v ...any) {
	handlerLogger.Info(v...)
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

func Info(v ...any) {
	handlerLogger.Info(v...)
}

func Debug(v ...any) {
	handlerLogger.Debug(v...)
}

func Warning(v ...any) {
	handlerLogger.Warning(v...)
}

func Error(v ...any) {
	handlerLogger.Error(v...)
}

func Infof(format string, v ...any) {
	handlerLogger.Infof(format, v...)
}

func Warningf(format string, v ...any) {
	handlerLogger.Warningf(format, v...)
}

func Debugf(format string, v ...any) {
	handlerLogger.Debugf(format, v...)
}

func Errorf(format string, v ...any) {
	handlerLogger.Errorf(format, v...)
}

// func IfNotNil(v ...any) {
// 	if len(v) == 0 {
// 		return
// 	}
// 	if isNil(v[len(v)-1]) {
// 		return
// 	}
// 	handlerLogger.Errorf("%v\n", v[len(v)-1])
// }
