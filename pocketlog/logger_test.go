package pocketlog_test

import (
	"fmt"
	"os"
	"pocket-logger/pocketlog"
	"testing"
)

func TestLoggerDebugf(t *testing.T) {
	debugLogger := pocketlog.New(pocketlog.LevelDebug, pocketlog.WithOutput(os.Stdout))
	debugLogger.Debugf("Hello, %s", "world")
}

const (
	debugMessage = "Why write I still all one, ever the same"
	infoMessage  = "And keep invention in a noted weed"
	errorMessage = "That every word doth almost tell my name"
)

type mockWriter struct {
	contents string
}

func (tw *mockWriter) Write(p []byte) (n int, err error) {
	tw.contents = tw.contents + string(p)
	fmt.Printf(tw.contents)
	return len(p), nil
}

func TestLogger_DebugfInfofErrorf(t *testing.T) {

	type testCase struct {
		level    pocketlog.Level
		expected string
	}

	tt := map[string]testCase{
		"Debug": {level: pocketlog.LevelDebug, expected: debugMessage},
		"Info":  {level: pocketlog.LevelInfo, expected: infoMessage},
		"Error": {level: pocketlog.LevelError, expected: errorMessage},
	}

	for name, tc := range tt {

		t.Run(name, func(t *testing.T) {
			tw := &mockWriter{}
			testLogger := pocketlog.New(tc.level, pocketlog.WithOutput(tw))
			testLogger.Debugf(tc.expected)
			testLogger.Infof(tc.expected)
			testLogger.Errorf(tc.expected)

			if tw.contents != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, tw.contents)
			}

		})
	}
}
