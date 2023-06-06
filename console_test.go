package Logger

import "testing"

func TestConsole(t *testing.T) {
	log := NewConsoleLogger("info")
	log.Debug("这是一条Debug日志")
	log.Info("这是一条Info日志")
	log.Error("这是一条error日志")
	log.Fatal("这是一条Fatal日志")
}
