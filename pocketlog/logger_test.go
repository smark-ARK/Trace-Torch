package pocketlog_test

import "github.com/smark-ark/trace-torch/pocketlog"

func ExampleLogger_Debugf() {
	debug_logger := pocketlog.New(pocketlog.LevelDebug)
	debug_logger.Debugf("Hello World!")
	// Output: Hello World!
}
