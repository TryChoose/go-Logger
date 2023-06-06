package Logger

import "testing"

func TestFileLogger_Info(t *testing.T) {
	log := NewFileLogger("info", "./", 10*1024)
	for i := 0; i < 100; i++ {
		log.Info("这是一条Info日志")
		log.Error("这是一条error日志")
		log.Debug("这是一条Debug日志")
		log.Fatal("这是一条Fatal日志")
	}

}
