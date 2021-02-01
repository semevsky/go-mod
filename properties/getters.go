// Copyright (c) 2020-2021, Yuriy Semevsky <semevskiy@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package properties

import "strings"

// Returns raw string
func (c *Config) GetString(name, defaultValue string) string {
	value, ok := (*c)[name]
	if !ok {
		value = defaultValue
	}
	return value
}

func (c *Config) GetInt(name string, defaultValue int) (i int, err error) {
	var i64 int64
	i64, err = c.getInt64(name, 32, int64(defaultValue))
	i = int(i64)
	return
}

func (c *Config) GetUint(name string, defaultValue uint) (u uint, err error) {
	var u64 uint64
	u64, err = c.getUint64(name, 32, uint64(defaultValue))
	u = uint(u64)
	return
}

// true values: "true", "yes" or "on" case insensitive
func (c *Config) GetBool(name string, defaultValue bool) (b bool) {
	s, ok := c.getString(name)
	if ok {
		s = strings.ToLower(s)
		b = (s == "true") || (s == "yes") || (s == "on")
	} else {
		b = defaultValue
	}
	return
}
