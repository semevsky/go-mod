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

func GetLogger(category string) Logger {
	return Logger{category}
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

// set access mode to log directories
// may be 0750 or 0700
func SetDirAccess(mode os.FileMode) {
	dirAccess = mode
}

// set access mode to log files
// may be 0640 or 0600
func SetFileAccess(mode os.FileMode) {
	fileAccess = mode
}

// reset to default values, 0755 - for dir, 0644 - for file
func ResetDirFileAccess() {
	dirAccess = DEFAULT_DIR_ACCESS
	fileAccess = DEFAULT_FILE_ACCESS
}

// name
func AddAppender(name string, level int, filename string, size int) {
	var wg sync.WaitGroup
	wg.Add(1)
	c := command{cmd: ADD, name: name, level: level, filename: filename, size: size, wg: &wg}
	commandChannel <- &c
	wg.Wait()
}

func Remove(name string) {
	var wg sync.WaitGroup
	wg.Add(1)
	c := command{cmd: REMOVE, name: name, wg: &wg}
	commandChannel <- &c
	wg.Wait()
}

func Reopen(name string) {
	var wg sync.WaitGroup
	wg.Add(1)
	c := command{cmd: REOPEN, name: name, wg: &wg}
	commandChannel <- &c
	wg.Wait()
}

// reopen all existing appenders
func ReopenAll() {
	var wg sync.WaitGroup
	wg.Add(len(appenders))
	for name := range appenders {
		c := command{cmd: REOPEN, name: name, wg: &wg}
		commandChannel <- &c
	}
	wg.Wait()
}

func RemoveAll() {
	var wg sync.WaitGroup
	wg.Add(len(appenders))
	for name := range appenders {
		c := command{cmd: REMOVE, name: name, wg: &wg}
		commandChannel <- &c
	}
	wg.Wait()
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

			c.done()

			if err != nil {
				println(err.Error())
			}
		}
	}
}