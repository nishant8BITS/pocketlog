package pocketlog_test

import (
	"pocket-logger/pocketlog"
	"testing"
)

func TestLoggerDebugf(t *testing.T) {
	debugLogger := pocketlog.New(pocketlog.LevelDebug)
	debugLogger.Debugf("Hello, %s", "world")
	t.Log("All is well")
}
