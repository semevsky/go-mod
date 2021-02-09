// Copyright (c) 2020-2021, Yuriy Semevsky <semevskiy@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package logging

import (
	"errors"
	"os"
	"sync"
)

var dataChannel = make(chan *message, 64)
var commandChannel = make(chan *command, 16)

var appenders = make(map[string]*Appender)

func init() {
	go execute()
}

// Creates new logger
// if there are no appenders - creates new one with empty name, TRACE, stdout, no buffer
func GetLogger(category string) Logger {
	return &logger{category}
}

func GetLevel(name string) (level int) {
	level = levelByName[name]
	return
}

// Use format symbols as descripbed at https://golang.org/pkg/time/#Time.Format
func SetTimeFormat(fmt string) {
	layout = fmt
}

func ResetTimeFormat() {
	layout = DEFAULT_LAYOUT
}

// Set file/directory access
// fMode - file acces, e.g. 0644 or 0640
// dMode - directory access, e.g. 0750 or 0700
func SetMode(fMode, dMode os.FileMode) {
	fileAccess = fMode
	dirAccess = dMode
}

// file/directory access to default values: 0644/0755
func ResetMode() {
	dirAccess = DEFAULT_DIR_ACCESS
	fileAccess = DEFAULT_FILE_ACCESS
}

// name - ID
// level - log no more than this log level
// filename - path, empty string means stdout
// size - buffer size, <= 0 means no buffering, write to file immediately
func AddAppender(name string, level int, filename string, size int) {
	var wg sync.WaitGroup
	wg.Add(1)
	c := command{cmd: ADD, name: name, level: level, filename: filename, size: size, wg: &wg}
	commandChannel <- &c
	wg.Wait()

	if c.err != nil {
		println(c.err)
	}
}

func RemoveAppender(name string) {
	var wg sync.WaitGroup
	wg.Add(1)
	c := command{cmd: REMOVE, name: name, wg: &wg}
	commandChannel <- &c
	wg.Wait()

	if c.err != nil {
		println(c.err)
	}
}

// reopen particular appender
func ReopenAppender(name string) {
	var wg sync.WaitGroup
	wg.Add(1)
	c := command{cmd: REOPEN, name: name, wg: &wg}
	commandChannel <- &c
	wg.Wait()

	if c.err != nil {
		println(c.err)
	}
}

// reopen all existing appenders
func ReopenAll() {
	broadcast(REOPEN)
}

// closes all appenders
func RemoveAll() {
	broadcast(REMOVE)
}

func broadcast(cmd int) {
	var cmds []*command
	var wg sync.WaitGroup
	wg.Add(len(appenders))
	for name := range appenders {
		c := command{cmd: cmd, name: name, wg: &wg}
		commandChannel <- &c
		cmds = append(cmds, &c)
	}
	wg.Wait()

	for _, c := range cmds {
		if c.err != nil {
			println(c.err)
		}
	}
}

// logging goroutine executing commands and writing data to log
// handles all the appenders in single thread
func execute() {
	for {
		select {
		// append the message
		case m := <-dataChannel:
			for _, a := range appenders {
				a.append(m)
			}
			m.done()
			// execute the command
		case c := <-commandChannel:
			var err error

			switch c.cmd {
			case ADD:
				_, is := appenders[c.name]
				if !is {
					a := Appender{name: c.name, level: c.level, filename: c.filename, size: c.size}
					err = a.open()
					if err == nil {
						appenders[c.name] = &a
					}
				} else {
					err = errors.New(c.name + " - appender already exists")
				}
			case REMOVE:
				a, is := appenders[c.name]
				if is {
					err = a.close()
				}
				if err == nil {
					delete(appenders, c.name)
				}
			case REOPEN:
				a, is := appenders[c.name]
				if is {
					err = a.reopen()
				}
			}
			c.done(err)
		}
	}
}
