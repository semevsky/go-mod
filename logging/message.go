// Copyright (c) 2020-2021, Yuriy Semevsky <semevskiy@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package logging

import (
	"sync"
	"time"
)


// default: yyyy-MM-dd HH:mm:ss.SSS X(tz)
const DEFAULT_LAYOUT = "2006-01-02 15:04:05.000 -07:00"

var layout = DEFAULT_LAYOUT // weird format symbols.. :/

type message struct {
	timestamp time.Time
	level     int
	category  string
	message   string

	wg *sync.WaitGroup
}

func (m *message) String() string {
	return m.timestamp.Format(layout) + " " + levelName[m.level] + " " + m.category + " " + m.message + "\n"
}

func (m *message) done() {
	if m.wg != nil {
		m.wg.Done()
	}
}
