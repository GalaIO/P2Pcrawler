package misc

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// user define level
const (
	traceLevel zapcore.Level = iota + 10
)

type zapLogger struct {
	prefix string
	inner  *zap.Logger
}

func (l *zapLogger) Debug(msg string, ctx Dict) {
	if l.prefix != "" {
		msg = l.prefix + msg
	}
	l.inner.Debug(msg, convert2slice(ctx)...)
}

func (l *zapLogger) Info(msg string, ctx Dict) {
	if l.prefix != "" {
		msg = l.prefix + msg
	}
	l.inner.Info(msg, convert2slice(ctx)...)
}

func (l *zapLogger) Warn(msg string, ctx Dict) {
	if l.prefix != "" {
		msg = l.prefix + msg
	}
	l.inner.Warn(msg, convert2slice(ctx)...)
}

func (l *zapLogger) Error(msg string, ctx Dict) {
	if l.prefix != "" {
		msg = l.prefix + msg
	}
	l.inner.Error(msg, convert2slice(ctx)...)
}

func (l *zapLogger) Panic(msg string, ctx Dict) {
	if l.prefix != "" {
		msg = l.prefix + msg
	}
	l.inner.Panic(msg, convert2slice(ctx)...)
}

func (l *zapLogger) Trace(msg string, ctx Dict) {
	if l.prefix != "" {
		msg = l.prefix + msg
	}
	l.inner.Check(traceLevel, msg).Write(convert2slice(ctx)...)
}

func (l *zapLogger) Flush() {
	l.inner.Sync()
}

// get a new prefix logger
func (l *zapLogger) SetPrefix(prefix string) Logger {
	return &zapLogger{
		prefix: fmt.Sprintf("[%s]", prefix),
		inner:  l.inner,
	}
}

func NewZapLogger(name string, logLevel string) Logger {
	// lumberjack.Logger is already safe for concurrent use, so we don't need to
	// lock it.
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   fmt.Sprintf("logs/%s.log", name),
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
	})
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		w,
		parseLogLevel(logLevel),
	)
	logger := zap.New(core)
	return &zapLogger{
		inner: logger,
	}
}

func parseLogLevel(logLevel string) zapcore.Level {
	switch logLevel {
	case "DEBUG":
		return zapcore.DebugLevel
	case "INFO":
		return zapcore.InfoLevel
	case "ERROR":
		return zapcore.ErrorLevel
	case "PANIC":
		return zapcore.PanicLevel
	default:
		return zapcore.InfoLevel
	}
}

func convert2slice(ctx Dict) []zap.Field {
	if ctx == nil || len(ctx) == 0 {
		return []zap.Field{}
	}
	fs := make([]zap.Field, len(ctx))
	i := 0
	for k, v := range ctx {
		fs[i] = zap.Any(k, v)
		i++
	}
	return fs
}
