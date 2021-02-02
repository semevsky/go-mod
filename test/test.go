package main

import (
	"fmt"
	"os"
	"strconv"

	"pkg.mytest.ru/go-mod/lipsum"
	"pkg.mytest.ru/go-mod/logging"
)

func main() {
	defer logging.RemoveAll()
	fmt.Println(lipsum.Lipsum(2))

	logger := logging.GetLogger("test")

	logging.AddAppender("time", logging.TRACE, "", 0)
	logger.LogSync(logging.INFO, "Just a test")
	logging.SetTimeFormat("06-1-2 3:4:5.999 -07")
	logger.LogSync(logging.INFO, "Just a test")
	logging.ResetTimeFormat()
	logger.LogSync(logging.INFO, "Just a test")
	logging.RemoveAppender("time")

	logName := "_tmp/log.log"
	backupName := "_tmp/log.bak"

	logging.AddAppender("test", logging.TRACE, logName, 1024)

	logger.Info("started")

	count := 0

	for ; count < 1000; count++ {
		logger.Info("just a test " + strconv.Itoa(count))
	}
	// note: this doesn't work on Windows
	err := os.Rename(logName, backupName)
	fmt.Println(err)

	logger.LogSync(logging.INFO, "reopening")
	logging.ReopenAll()
	logger.LogSync(logging.INFO, "reopened")

	for ; count < 1100; count++ {
		logger.Info("just a test " + strconv.Itoa(count))
	}
	logger.LogSync(logging.INFO, "closing")
}
