// Copyright (c) 2020-2021, Yuriy Semevsky <semevskiy@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package logging

import "sync"

const (
	ADD = iota
	REMOVE
	REOPEN
)

var cmdName = [3]string{
	"ADD",
	"REMOVE",
	"REOPEN",
}

type command struct {
	cmd      int
	name     string
	level    int
	filename string
	size     int

	err error

	wg *sync.WaitGroup
}

func (c *command) String() string {
	// <cmd> <name> <LEVEL> "<filename>"
	return cmdName[c.cmd] + " " + c.name + " " + levelName[c.level] + " \"" + c.filename + "\""
}

func (c *command) done(err error) {
	c.err = err
	if c.wg != nil {
		c.wg.Done()
	}
}
