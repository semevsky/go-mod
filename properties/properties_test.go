// Copyright (c) 2020-2021, Yuriy Semevsky <semevskiy@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package properties

import (
	"testing"
)

var data = `# Test config
string=Lorem ipsum
int=-123
uint=123
float=123.45
bool=true
boolNeg=off
last=last`

var config Config

func init() {
	config = LoadString(data)
}

func TestConfig_GetBool(t *testing.T) {
	type args struct {
		name         string
		defaultValue bool
	}
	tests := []struct {
		name  string
		c     Config
		args  args
		wantB bool
	}{
		{name: "true", c: config, args: args{name: "bool", defaultValue: false}, wantB: true},
		{name: "false", c: config, args: args{name: "boolNeg", defaultValue: true}, wantB: false},
		{name: "default", c: config, args: args{name: "none", defaultValue: true}, wantB: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotB := tt.c.GetBool(tt.args.name, tt.args.defaultValue); gotB != tt.wantB {
				t.Errorf("GetBool() = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}

func TestConfig_GetInt(t *testing.T) {
	type args struct {
		name         string
		defaultValue int
	}
	tests := []struct {
		name    string
		c       Config
		args    args
		wantI   int
		wantErr bool
	}{
		{name: "positive", c: config, args: args{name: "int", defaultValue: 0}, wantI: -123, wantErr: false},
		{name: "default", c: config, args: args{name: "none", defaultValue: 0}, wantI: 0, wantErr: false},
		{name: "negative", c: config, args: args{name: "string", defaultValue: 0}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotI, err := tt.c.GetInt(tt.args.name, tt.args.defaultValue)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotI != tt.wantI {
				t.Errorf("GetInt() gotI = %v, want %v", gotI, tt.wantI)
			}
		})
	}
}

func TestConfig_GetString(t *testing.T) {
	type args struct {
		name         string
		defaultValue string
	}
	tests := []struct {
		name string
		c    Config
		args args
		want string
	}{
		{name: "positive", c: config, args: args{name: "string", defaultValue: "Lorem ipsum"}, want: "Lorem ipsum"},
		{name: "default", c: config, args: args{name: "none", defaultValue: "none"}, want: "none"},
		{name: "eof", c: config, args: args{name: "last", defaultValue: ""}, want: "last"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.GetString(tt.args.name, tt.args.defaultValue); got != tt.want {
				t.Errorf("GetString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_GetUint(t *testing.T) {
	type args struct {
		name         string
		defaultValue uint
	}
	tests := []struct {
		name    string
		c       Config
		args    args
		wantU   uint
		wantErr bool
	}{
		{name: "positive", c: config, args: args{name: "uint", defaultValue: 0}, wantU: 123, wantErr: false},
		{name: "default", c: config, args: args{name: "none", defaultValue: 0}, wantU: 0, wantErr: false},
		{name: "negative", c: config, args: args{name: "string", defaultValue: 0}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotU, err := tt.c.GetUint(tt.args.name, tt.args.defaultValue)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotU != tt.wantU {
				t.Errorf("GetUint() gotU = %v, want %v", gotU, tt.wantU)
			}
		})
	}
}
