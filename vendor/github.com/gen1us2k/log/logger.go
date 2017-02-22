package log

import (
	"io"
	"sync"
)

type (
	Logger interface {
		Panic(v ...interface{})
		Fatal(v ...interface{})
		Error(v ...interface{})
		Warning(v ...interface{})
		Info(v ...interface{})
		Debug(v ...interface{})

		Panicf(format string, v ...interface{})
		Fatalf(format string, v ...interface{})
		Errorf(format string, v ...interface{})
		Warningf(format string, v ...interface{})
		Infof(format string, v ...interface{})
		Debugf(format string, v ...interface{})

		SetLevel(level int)
		Level() int

		SetName(name string)
		Name() string

		SetFormatter(formatter Formatter)
		Formatter() Formatter
	}

	CustomLogger interface {
		Logger
		SetOutput(io.Writer)
	}

	Formatter interface {
		Format(out io.Writer, level int, channel string, msg string)
	}
)

var (
	defaultLevel = LevelWarning
	defaultName  = "main"
	logger       = newChannel(defaultName, defaultLevel)
	channels     = map[string]Logger{
		defaultName: logger,
	}
	customChannels = map[string]CustomLogger{}

	me sync.Mutex
)
