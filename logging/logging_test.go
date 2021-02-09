package logging

import (
	"testing"
)

func init() {
	AddAppender("test", TRACE, "/dev/null", 0)
	l = GetLogger("test")
}

var l Logger

func BenchmarkLog(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l.Log(INFO, "just a test")
	}
}

func BenchmarkLogSync(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l.LogSync(INFO, "just a test")
	}
}
