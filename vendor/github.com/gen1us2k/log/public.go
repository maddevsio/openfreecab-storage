package log

import (
	"io"
)

// NewLogger creates new logger instance
func NewLogger(name string) Logger {
	me.Lock()
	defer me.Unlock()

	log, ok := channels[name]
	if !ok {
		log = newChannel(name, defaultLevel)
		channels[name] = log
	}

	return log
}

// NewCustomLogger creates new logger instance with custom out stream
func NewCustomLogger(name string, out io.Writer) CustomLogger {
	me.Lock()
	defer me.Unlock()

	log, ok := customChannels[name]

	if !ok {
		log = newCustomChannel(name, defaultLevel, out)
		customChannels[name] = log
	} else {
		log.SetOutput(out)
	}

	return log
}

// RemoveLogger removes logger
func RemoveLogger(name string) {
	me.Lock()
	defer me.Unlock()
	delete(channels, name)
}

// RemoveLogger removes logger
func RemoveCustomLogger(name string) {
	me.Lock()
	defer me.Unlock()
	delete(customChannels, name)
}

// Panic is equivalent to Print() followed by a call to panic().
func Panic(v ...interface{}) {
	logger.Panic(v...)
}

// Fatal is equivalent to Print() followed by a call to os.Exit(1).
func Fatal(v ...interface{}) {
	logger.Fatal(v...)
}

// Error is equivalent to Print() followed by error log level
func Error(v ...interface{}) {
	logger.Error(v...)
}

// Warning is equivalent to Print() followed by warning log level
func Warning(v ...interface{}) {
	logger.Warning(v...)
}

// Info is equivalent to Print() followed by info log level
func Info(v ...interface{}) {
	logger.Info(v...)
}

// Debug is equivalent to Print() followed by debug log level
func Debug(v ...interface{}) {
	logger.Debug(v...)
}

func Panicf(format string, v ...interface{}) {
	logger.Panicf(format, v...)
}

func Fatalf(format string, v ...interface{}) {
	logger.Fatalf(format, v...)
}

func Errorf(format string, v ...interface{}) {
	logger.Errorf(format, v...)
}

func Warningf(format string, v ...interface{}) {
	logger.Warningf(format, v...)
}

func Infof(format string, v ...interface{}) {
	logger.Infof(format, v...)
}

func Debugf(format string, v ...interface{}) {
	logger.Debugf(format, v...)
}

// SetLevel sets level of logger
func SetLevel(level int) {
	me.Lock()
	defer me.Unlock()

	defaultLevel = level

	for _, log := range channels {
		log.SetLevel(level)
	}

	for _, log := range customChannels {
		log.SetLevel(level)
	}
}
