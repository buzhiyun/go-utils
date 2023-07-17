package log

import (
	"io"
	"os"
)

type logCfg struct {
	Color     bool
	AddOutput []io.Writer
	Level     string
}

var (
	zapCfg = logCfg{
		Color:     true,
		AddOutput: []io.Writer{os.Stdout},
		Level:     "info",
	}
	logger      = newZap(zapCfg)
	sugarLogger = logger.Sugar()
)

func (z *logCfg) Refresh() {
	logger = newZap(zapCfg)
	sugarLogger = logger.Sugar()
}

func DisableColor() {
	zapCfg.Color = false
	zapCfg.Refresh()
}

func Info(args ...interface{}) {
	sugarLogger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	sugarLogger.Infof(template, args...)
}

func Debug(args ...interface{}) {
	sugarLogger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	sugarLogger.Debugf(template, args...)
}

func Warn(args ...interface{}) {
	sugarLogger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	sugarLogger.Warnf(template, args...)
}

func Error(args ...interface{}) {
	sugarLogger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	sugarLogger.Errorf(template, args...)
}

func Fatal(args ...interface{}) {
	sugarLogger.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	sugarLogger.Fatalf(template, args...)
}

func AddOutput(writers ...io.Writer) {
	zapCfg.AddOutput = append(zapCfg.AddOutput, writers...)
	zapCfg.Refresh()
}

func SetLevel(level string) {
	zapCfg.Level = level
	zapCfg.Refresh()
}

func Sync() {
	_ = sugarLogger.Sync()
}

func RemoveAllOutput() {
	zapCfg.AddOutput = nil
	zapCfg.Refresh()
}
