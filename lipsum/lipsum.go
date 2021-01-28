// Copyright (c) 2021, Yuriy Semevsky <semevskiy@gmail.com>
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file. 
package lipsum

import "strings"

const source = "Lorem ipsum dolor sit amet est eu velit eligendi moderatius ne vis vitae nonumy deserunt Agam mucius nominavi vim ut at erroribus maluisset  referrentur vis Has reque vulputate consequuntur in ex ius vivendo oporteat  torquatos Ne soluta definitionem vim consul integre erroribus ne est Choro  vitae invidunt eu mel populo expetenda efficiantur ut nec veri nullam  senserit id vim"
var array = strings.Fields(source)

func Lipsum(count int) string {
    if count > len(array) {
        count = len(array)
    }
    tmp := array[:count]
    result := strings.Join(tmp, " ")
    return result
}
