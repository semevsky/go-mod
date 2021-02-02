# Go-mod - general purpose Go module
Copyright (c) 2020-2021, Yuriy Semevsky <semevskiy@gmail.com>. All rights reserved.  
Use of the source code is governed by a BSD-style  
license that can be found in the LICENSE file.

Usage: `import "pkg.mytest.ru/go-mod/<folder>"`  
e.g.
```Go
    import (
    	"pkg.mytest.ru/go-mod/lipsum"
    )
    ...
    func ... {
        println(lipsum.Lipsum(5))
    }
```

**Folders**
- [lipsum](lipsum) - "Lorem ipsum" generator, was created just for fun and improve Go experience
- [logging](logging) - configurable asynchronous logger
- [properties](properties) - simplest properties file reader
- [signals](signals) - handling system signals, useful for *nixes
- [test](test) - stuff for non-standard testing of the libs

Go docs here: https://pkg.go.dev/pkg.mytest.ru/go-mod
