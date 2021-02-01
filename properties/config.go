// Copyright (c) 2020-2021, Yuriy Semevsky <semevskiy@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package properties

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Config map[string]string

func LoadFile(path string) (c Config, err error) {
	var buffer []byte
	buffer, err = ioutil.ReadFile(path)
	if err != nil {
		return
	}

	input := string(buffer)

	c = LoadString(input)
	return
}

func LoadString(input string) (c Config) {
	c = make(Config)

	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		line := scanner.Text()
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
