package test5

import "testing"

func TestLog(t *testing.T) {
	logger.Debug("debug")
	logger.Info("info")
	logger.Warn("warn")
	logger.Error("error")
}
