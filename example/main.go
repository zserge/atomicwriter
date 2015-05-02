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
