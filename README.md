# atomicwriter

Atomic file writer for Go. It uses a temporary file for writing, and when close
is performed - the temporary file is renamed back to the original file name.

Rename is supposed to be atomic (as it is on most platforms).

Example:

``` go
package main

import (
	"log"

	"github.com/zserge/atomicwriter"
)

func main() {
	f, err := atomicwriter.NewWriter("file.txt")
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()

	f.Write([]byte("Hello"))
	f.Write([]byte("world"))
	f.Write([]byte("\n"))
}
```
