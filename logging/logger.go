// Copyright (c) 2020-2021, Yuriy Semevsky <semevskiy@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package logging

import (
	"sync"
	"time"
)

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

type Logger struct {
	category string
}

func (l *Logger) Log(level int, msg string) {
	m := message{timestamp: time.Now(), level: level, category: l.category, message: msg}
	dataChannel <- &m
}

// Note: slow logging, use with care
func (l *Logger) LogSync(level int, msg string) {
	var wg sync.WaitGroup
	wg.Add(1)
	m := message{timestamp: time.Now(),level: level,category: l.category,message: msg, wg: &wg}
	dataChannel <- &m
	wg.Wait()
}

func (l *Logger) Fatal(message string) {
	l.Log(FATAL, message)
}

func (l *Logger) Error(message string) {
	l.Log(ERROR, message)
}

func (l *Logger) Warn(message string) {
	l.Log(WARN, message)
}

func (l *Logger) Info(message string) {
	l.Log(INFO, message)
}

func (l *Logger) Debug(message string) {
	l.Log(DEBUG, message)
}

func (l *Logger) Trace(message string) {
	l.Log(TRACE, message)
}
