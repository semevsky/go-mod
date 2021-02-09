// Copyright (c) 2020-2021, Yuriy Semevsky <semevskiy@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package logging

import (
	"sync"
	"time"
)

type logger struct {
	category string
}

func (l *logger) Log(level int, msg string) {
	m := message{timestamp: time.Now(), level: level, category: l.category, message: msg}
	dataChannel <- &m
}

func (l *logger) LogSync(level int, msg string) {
	var wg sync.WaitGroup
	wg.Add(1)
	m := message{timestamp: time.Now(), level: level, category: l.category, message: msg, wg: &wg}
	dataChannel <- &m
	wg.Wait()
}

func (l *logger) Fatal(message string) {
	l.Log(FATAL, message)
}

func (l *logger) Error(message string) {
	l.Log(ERROR, message)
}

func (l *logger) Warn(message string) {
	l.Log(WARN, message)
}

func (l *logger) Info(message string) {
	l.Log(INFO, message)
}

func (l *logger) Debug(message string) {
	l.Log(DEBUG, message)
}

func (l *logger) Trace(message string) {
	l.Log(TRACE, message)
}
