package llog

import (
	"fmt"
	"log"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime)
}

func Debug(v ...interface{}) {
	lLogPrint("debug", v...)
}

func Info(v ...interface{}) {
	lLogPrint("info", v...)
}

func Warn(v ...interface{}) {
	lLogPrint("warn", v...)
}

func Error(v ...interface{}) {
	lLogPrint("error", v...)
}

func Fatal(v ...interface{}) {
	lLogFatal("fatal", v...)
}

func Panic(v ...interface{}) {
	lLogPanic("panic", v...)
}

func lLogPrint(level string, v ...interface{}) {
	msg := fmt.Sprintf("[%s] %s", level, fmt.Sprint(v...))
	log.Printf(msg)
}

func lLogFatal(level string, v ...interface{}) {
	msg := fmt.Sprintf("[%s] %s", level, fmt.Sprint(v...))
	log.Fatalf(msg)
}

func lLogPanic(level string, v ...interface{}) {
	msg := fmt.Sprintf("[%s] %s", level, fmt.Sprint(v...))
	log.Panicf(msg)
}
