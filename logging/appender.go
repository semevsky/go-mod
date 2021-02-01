// Copyright (c) 2020-2021, Yuriy Semevsky <semevskiy@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package logging

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

const (
	DEFAULT_DIR_ACCESS = 0755
	DEFAULT_FILE_ACCESS = 0644
)

var dirAccess os.FileMode = DEFAULT_DIR_ACCESS
var fileAccess os.FileMode = DEFAULT_FILE_ACCESS

type Appender struct {
	name     string
	filename string
	size     int

	level int

	file   *os.File
	buffer *bufio.Writer

	writer io.Writer
}

func (a *Appender) append(m *message) {
	if m.level <= a.level {
		s := m.String()
		if !strings.HasSuffix(s, "\n") {
			s += "\n"
		}
		// flush intentionally
		if a.buffer != nil && len(s) > a.buffer.Available() {
			err := a.buffer.Flush()
			if err != nil {
				println(err.Error())
			}
		}

		_, err := a.writer.Write([]byte(s))
		if err != nil {
			println(err.Error())
		}
	}
}

func (a *Appender) close() (err error) {
	if a.buffer != nil {
		err = a.buffer.Flush()
		for count := 0; count < 10; count++ {
			if err == nil {
				break
			}
		}
	}
	if a.file != nil {
		for count := 0; count < 10; count++ {
			err = a.file.Close()
			if err == nil {
				break
			}
		}
	}
	return
}

func (a *Appender) open() (err error) {
	a.file, a.buffer, a.writer, err = openFile(a.filename, a.size)
	return
}

func (a *Appender) reopen() (err error) {
	err = a.close()
	if err == nil {
		err = a.open()
	}
	return
}

func openFile(filename string, size int) (file *os.File, buffer *bufio.Writer, writer io.Writer, err error) {
	if filename == "" {
		writer = os.Stdout
	} else if filename == "/dev/null" {
		writer = ioutil.Discard
	} else {
		dir := path.Dir(filename)
		_, err = os.Stat(dir)
		if err != nil {
			err = os.MkdirAll(dir, dirAccess)
		}
		if err == nil {
			file, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, fileAccess)
		}
		if err == nil {
			if size > 0 {
				buffer = bufio.NewWriterSize(file, size)
				writer = buffer
			} else {
				writer = file
			}
		}
	}
	return
}
