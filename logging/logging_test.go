package logging

import (
	"testing"
)

func init() {
	AddAppender("test", TRACE, "/dev/null", 0)
	logger = GetLogger("test")
}

var logger Logger

func BenchmarkLog(b *testing.B) {
	for i := 0; i < b.N; i++ {
		logger.Log(INFO, "just a test")
	}
}

func BenchmarkLogSync(b *testing.B) {
	for i := 0; i < b.N; i++ {
		logger.LogSync(INFO, "just a test")
	}
}
