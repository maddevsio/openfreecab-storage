package log

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestLogWithoutFormat(t *testing.T) {
	expected := "1 2 3\n"
	var buf bytes.Buffer
	SetLevel(MustParseLevel("debug"))
	logger := NewCustomLogger("test", &buf)
	logger.Info(1, 2, 3)
	result := buf.String()
	fmt.Println(result)
	if !strings.HasSuffix(result, expected) {
		t.Errorf("Not ends with: [%s] actual: [%s]", expected, result)
	}
}

func TestLogWithFormat(t *testing.T) {
	expected := "log 1 test\n"
	var buf bytes.Buffer
	SetLevel(MustParseLevel("debug"))
	logger := NewCustomLogger("test2", &buf)
	logger.Infof("%s %d %s", "log", 1, "test")
	result := buf.String()
	fmt.Println(result)
	if !strings.HasSuffix(result, expected) {
		t.Errorf("Not ends with: [%s] actual: [%s]", expected, result)
	}
}

func TestLogWithMultipleSameInstance_ShouldUsedLastWriter(t *testing.T) {
	expected := "log 1 test\n"
	expected2 := "log2 2 test2\n"

	var buf bytes.Buffer
	var buf2 bytes.Buffer
	SetLevel(MustParseLevel("debug"))

	nameOfLogger := "test3"
	logger := NewCustomLogger(nameOfLogger, &buf)
	logger2 := NewCustomLogger(nameOfLogger, &buf2)

	logger.Infof("%s %d %s", "log", 1, "test")
	logger2.Infof("%s %d %s", "log2", 2, "test2")

	result := buf.String()
	result2 := buf2.String()

	if result != "" {
		t.Errorf("Second buffer should used for output")
	}

	if !strings.HasSuffix(result2, expected2) {
		t.Errorf("Not ends with: [%s] actual: [%s]", expected, result)
	}
}

func TestRemoveCustomLogger(t *testing.T) {
	var buf bytes.Buffer

	SetLevel(MustParseLevel("debug"))
	loggerForRemoveName := "removeTest1"
	logger := NewCustomLogger(loggerForRemoveName, &buf)
	RemoveCustomLogger(loggerForRemoveName)

	logger2 := NewCustomLogger(loggerForRemoveName, &buf)
	if logger == logger2 {
		t.Error("CustomLogger doesn't removed")
	}
}

func TestRemoveLogger(t *testing.T) {
	SetLevel(MustParseLevel("debug"))
	loggerForRemoveName := "removeTest2"
	logger := NewLogger(loggerForRemoveName)
	RemoveLogger(loggerForRemoveName)

	logger2 := NewLogger(loggerForRemoveName)
	if logger == logger2 {
		t.Error("Logger doesn't removed")
	}
}
