package console

import (
	"fmt"
	"os"
)

var DefaultLogger = NewLogger()

func Exit(v ...any) {
	DefaultLogger.Info(v...)
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
	DefaultLogger.Info(v...)
}

func Debug(v ...any) {
	DefaultLogger.Debug(v...)
}

func Warning(v ...any) {
	DefaultLogger.Warning(v...)
}

func Error(v ...any) {
	DefaultLogger.Error(v...)
}

func Infof(format string, v ...any) {
	DefaultLogger.Infof(format, v...)
}

func Warningf(format string, v ...any) {
	DefaultLogger.Warningf(format, v...)
}

func Debugf(format string, v ...any) {
	DefaultLogger.Debugf(format, v...)
}

func Errorf(format string, v ...any) {
	DefaultLogger.Errorf(format, v...)
}
