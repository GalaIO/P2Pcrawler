package misc

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	PANIC
	TRACE
)

type LoggerFactory interface {
	GetDefaultName() string
	NewLogger(name string, level LogLevel) Logger
}

type Logger interface {
	SetPrefix(prefix string) Logger
	Debug(msg string, ctx Dict)
	Info(msg string, ctx Dict)
	Warn(msg string, ctx Dict)
	Error(msg string, ctx Dict)
	Panic(msg string, ctx Dict)
	Trace(msg string, ctx Dict)
	Flush()
}

type bufferLogger struct {
	sync.Mutex
	bufferedLoggers map[string]Logger
	level           LogLevel
}

// global var
var bufferedLoggers = &bufferLogger{
	bufferedLoggers: make(map[string]Logger, 10),
	level:           DEBUG,
}

var defaultLoggerFactory LoggerFactory = NewConsoleLoggerFactory()

func (b *bufferLogger) mapLogger(name string) Logger {
	if v, ok := b.bufferedLoggers[name]; ok {
		return v
	}
	b.Lock()
	defer b.Unlock()
	if v, ok := b.bufferedLoggers[name]; ok {
		return v
	}
	logger := defaultLoggerFactory.NewLogger(name, b.level)
	b.bufferedLoggers[name] = logger
	return logger
}

func (b *bufferLogger) flush() error {
	for _, v := range b.bufferedLoggers {
		v.Flush()
	}
	return nil
}

// using this function when programe init
func InitLogger(factory LoggerFactory, level LogLevel) {
	defaultLoggerFactory = factory
	bufferedLoggers.level = level
}

// different name whit different log file, you can add prefix to set [gin] module log
func GetLogger(names ...string) Logger {
	loggerName := defaultLoggerFactory.GetDefaultName()
	if len(names) > 0 {
		loggerName = names[0]
	}
	return bufferedLoggers.mapLogger(loggerName)
}

func Close() error {
	return bufferedLoggers.flush()
}

// logger default console

type consoleLoggerFactory struct {
}

func NewConsoleLoggerFactory() *consoleLoggerFactory {
	return &consoleLoggerFactory{}
}

func (f *consoleLoggerFactory) GetDefaultName() string {
	return ""
}

func (f *consoleLoggerFactory) NewLogger(name string, level LogLevel) Logger {
	return &consoleLogger{
		prefix: "",
		level:  level,
	}
}

type consoleLogger struct {
	prefix string
	level  LogLevel
}

func (c *consoleLogger) SetPrefix(prefix string) Logger {
	c.prefix = prefix
	return c
}

func (c *consoleLogger) Debug(msg string, ctx Dict) {
	if c.level > DEBUG {
		return
	}
	c.print("Debug", msg, ctx)
}

func (c *consoleLogger) print(level string, msg string, ctx Dict) {
	timeStr := time.Now().Format("2006-01-02 15:04:05.000")
	msg = appendCtx(msg, ctx)
	if len(c.prefix) > 0 {
		fmt.Printf("%s %6s: [%s]%s\r\n", timeStr, level, c.prefix, msg)
		return
	}
	fmt.Printf("%s %6s: %s\r\n", timeStr, level, msg)
}

func (c *consoleLogger) Info(msg string, ctx Dict) {
	if c.level > INFO {
		return
	}
	c.print("Info", msg, ctx)
}

func (c *consoleLogger) Warn(msg string, ctx Dict) {
	if c.level > WARN {
		return
	}
	c.print("Warn", msg, ctx)
}

func (c *consoleLogger) Error(msg string, ctx Dict) {
	if c.level > ERROR {
		return
	}
	c.print("Error", msg, ctx)
}

func (c *consoleLogger) Panic(msg string, ctx Dict) {
	if c.level > PANIC {
		return
	}
	c.print("Panic", msg, ctx)
	panic(msg)
}

func (c *consoleLogger) Trace(msg string, ctx Dict) {
	if c.level > TRACE {
		return
	}
	c.print("Trace", msg, ctx)
}

func (c *consoleLogger) Flush() {
	return
}

func appendCtx(msg string, ctx Dict) string {
	if ctx == nil {
		return msg
	}

	i := 0
	var builder strings.Builder
	builder.Grow(2 * len(msg))
	builder.WriteString(msg)
	builder.WriteString(", ")
	for k, v := range ctx {
		if i > 0 {
			builder.WriteByte('|')
		}
		builder.WriteString(k)
		builder.WriteByte(':')
		builder.WriteString(ToString(v))
		i++
	}
	return builder.String()
}
