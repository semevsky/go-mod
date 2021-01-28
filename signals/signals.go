// Copyright (c) 2021, Yuriy Semevsky <semevskiy@gmail.com>
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file. 
package signals

import (
    "os"
    "os/signal"
    "syscall"
)

const (
    SIGUSR1 = syscall.Signal(0xa)
    SIGUSR2 = syscall.Signal(0xc)
)

func RegisterSignals(sigs ...os.Signal) (c chan os.Signal) {
    c = make(chan os.Signal, 1)
    signal.Notify(c, sigs...)
    return
}
