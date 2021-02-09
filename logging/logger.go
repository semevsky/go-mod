// Copyright (c) 2021, Yuriy Semevsky <semevskiy@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package logging

const (
	FATAL = iota
	ERROR
	WARN
	INFO
	DEBUG
	TRACE
)

var levelName = [6]string{
	"FATAL",
	"ERROR",
	"WARN",
	"INFO",
	"DEBUG",
	"TRACE",
}

var levelByName = map[string]int{
	"FATAL": FATAL,
	"ERROR": ERROR,
	"WARN":  WARN,
	"INFO":  INFO,
	"DEBUG": DEBUG,
	"TRACE": TRACE,
}

type Logger interface {
	Log(level int, msg string)
	// Note: slow logging, use with care
	LogSync(level int, msg string)
	Fatal(message string)
	Error(message string)
	Warn(message string)
	Info(message string)
	Debug(message string)
	Trace(message string)
}
