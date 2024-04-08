package logger

import (
	"sync/atomic"

	"go.uber.org/zap"
)

var defaultLogger atomic.Pointer[zap.Logger]

func init() {
	l, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(l)
	defaultLogger.Store(l)
}

func Default() *zap.Logger { return defaultLogger.Load() }

func SetDefault(l *zap.Logger) { defaultLogger.Store(l) }

func Debug(msg string, args ...interface{}) {
	Default().Sugar().Debug(msg)
}

func Warn(args ...interface{}) {
	Default().Sugar().Warn(args...)
}

func Warning(msg string, args ...interface{}) {
	Default().Sugar().Warnf(msg, args...)
}

func Warningf(tpl string, args ...interface{}) {
	Default().Sugar().Warnf(tpl, args...)
}

func SDebugf(tpl string, args ...interface{}) {
	Default().Sugar().Debugf(tpl, args...)
}

func SDebug(args ...interface{}) {
	Default().Sugar().Debug(args...)
}

func SInfof(tpl string, args ...interface{}) {
	Default().Sugar().Infof(tpl, args...)
}

func SFatal(args ...interface{}) {
	Default().Sugar().Fatal(args...)
}

func SFatalf(tpl string, args ...interface{}) {
	Default().Sugar().Fatalf(tpl, args...)
}

func SInfo(args ...interface{}) {
	Default().Sugar().Info(args...)
}

func SErrorf(tpl string, args ...interface{}) {
	Default().Sugar().Errorf(tpl, args...)
}

func SError(args ...interface{}) {
	Default().Sugar().Error(args...)
}

func SWarnf(tpl string, args ...interface{}) {
	Default().Sugar().Warnf(tpl, args...)
}

func SWarn(args ...interface{}) {
	Default().Sugar().Warn(args...)
}

func SPanicf(tpl string, args ...interface{}) {
	Default().Sugar().Panicf(tpl, args...)
}

func Sync() error {
	return Default().Sugar().Sync()
}
