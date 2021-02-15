// Copyright (c) 2021, Yuriy Semevsky <semevskiy@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file. 
package main

import (
	"fmt"
	"pkg.mytest.ru/go-mod/lipsum"
)

func main() {
	fmt.Println(lipsum.Lipsum(5))
}
