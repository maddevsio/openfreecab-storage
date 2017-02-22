package log

import (
	"fmt"
	"io"
	"os"
	"sync"
)

type (
	ch struct {
		mu        sync.Mutex
		name      string
		level     int
		out       io.Writer
		formatter Formatter
	}
)

func newChannel(name string, level int) Logger {
	return &ch{
		name:      name,
		level:     level,
		out:       os.Stderr,
		formatter: newFormatter(),
	}
}

func newCustomChannel(name string, level int, out io.Writer) CustomLogger {
	return &ch{
		name:      name,
		level:     level,
		out:       out,
		formatter: newFormatter(),
	}
}

func (c *ch) SetOutput(out io.Writer) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.out = out
}

func (c *ch) Panic(v ...interface{}) {
	c.log(LevelPanic, v...)
	panic(fmt.Sprint(v...))
}

func (c *ch) Fatal(v ...interface{}) {
	c.log(LevelFatal, v...)
	os.Exit(1)
}

func (c *ch) Error(v ...interface{}) {
	c.log(LevelError, v...)
}

func (c *ch) Warning(v ...interface{}) {
	c.log(LevelWarning, v...)
}

func (c *ch) Info(v ...interface{}) {
	c.log(LevelInfo, v...)
}

func (c *ch) Debug(v ...interface{}) {
	c.log(LevelDebug, v...)
}

func (c *ch) Panicf(format string, v ...interface{}) {
	c.logf(LevelPanic, format, v...)
	panic(fmt.Sprintf(format, v...))
}

func (c *ch) Fatalf(format string, v ...interface{}) {
	c.logf(LevelFatal, format, v...)
	os.Exit(1)
}

func (c *ch) Errorf(format string, v ...interface{}) {
	c.logf(LevelError, format, v...)
}

func (c *ch) Warningf(format string, v ...interface{}) {
	c.logf(LevelWarning, format, v...)
}

func (c *ch) Infof(format string, v ...interface{}) {
	c.logf(LevelInfo, format, v...)
}

func (c *ch) Debugf(format string, v ...interface{}) {
	c.logf(LevelDebug, format, v...)
}

func (c *ch) SetLevel(level int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.level = level
}

func (c *ch) Level() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.level
}

func (c *ch) SetName(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.name = name
}

func (c *ch) Name() string {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.name
}

func (c *ch) SetFormatter(formatter Formatter) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.formatter = formatter
}

func (c *ch) Formatter() Formatter {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.formatter
}

func (c *ch) log(level int, v ...interface{}) {
	if c.level < level {
		return
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.formatter.Format(c.out, level, c.name, fmt.Sprintln(v...))
}

func (c *ch) logf(level int, format string, v ...interface{}) {
	if c.level < level {
		return
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.formatter.Format(c.out, level, c.name, fmt.Sprintf(format, v...))
}
