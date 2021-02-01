// Copyright (c) 2020-2021, Yuriy Semevsky <semevskiy@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package properties

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config map[string]string

func LoadFile(path string) (c Config, err error) {
	var file *os.File
	file, err = os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	c = load(reader)
	return
}

func LoadString(data string) (c Config) {
	reader := bufio.NewReader(strings.NewReader(data))
	c = load(reader)
	return
}

func load(reader *bufio.Reader) (c Config) {
	c = make(Config)
	for line, err1 := reader.ReadString('\n'); err1 == nil; line, err1 = reader.ReadString('\n') {
		err := c.parse(line)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
	return
}

func (this *Config) parse(line string) (err error) {
	line = strings.TrimSpace(line)
	if line == "" || strings.HasPrefix(line, "#") {
		return
	}
	index := strings.Index(line, "=")
	if index < 0 {
		err = errors.New("No \"=\" mark in line " + line)
		return
	}

	name := line[:index]
	value := line[index+1:]
	(*this)[name] = value

	return
}

func (c *Config) getString(name string) (value string, ok bool) {
	value, ok = (*c)[name]
	return
}

func (c *Config) getInt64(name string, bitSize int, defaultValue int64) (i int64, err error) {
	value, ok := c.getString(name)
	if ok {
		i, err = strconv.ParseInt(value, 10, bitSize)
	}else{
		i = defaultValue
	}
	return
}

func (c *Config) getUint64(name string, bitSize int, defaultValue uint64) (u uint64, err error) {
	value, ok := c.getString(name)
	if ok {
		u, err = strconv.ParseUint(value, 10, bitSize)
	}else{
		u = defaultValue
	}
	return
}

func (c *Config) getFloat64(name string, bitSize int, defaultValue float64) (f float64, err error) {
	value, ok := c.getString(name)
	if ok {
		f, err = strconv.ParseFloat(value, bitSize)
	}else{
		f = defaultValue
	}
	return
}
