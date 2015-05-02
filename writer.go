package atomicwriter

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"sync/atomic"
)

type atomicWriter struct {
	name string
	f    *os.File
}

var nextID uint32 = 0

//
// suffix() generates unique suffix for the file. Suffix includes the md5 hash
// of the filename, PID and number of atomic writes that have been done so far.
// This allows to safely write into the same file from two different
// processes/goroutines and be sure that only the last write will be stored.
func suffix(name string) string {
	h := md5.New()
	h.Write([]byte(name))
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, uint32(os.Getpid()))
	h.Write(buf)
	binary.LittleEndian.PutUint32(buf, atomic.AddUint32(&nextID, 1))
	h.Write(buf)
	return fmt.Sprintf(".%x", h.Sum(nil)) + ".tmp"
}

// NewWriter() returns a new file writer for the given filename.  The writer
// creates some unique temporary file for writing, and once write is finished -
// it renames the temporary file to the original given filename (rename is
// atomic on most systems).
func NewWriter(name string) (io.WriteCloser, error) {
	f, err := os.Create(name + suffix(name))
	if err != nil {
		return nil, err
	}
	return &atomicWriter{name, f}, nil
}

// Write() writes data into the file. Bytes will be written to the temporary file.
func (aw *atomicWriter) Write(b []byte) (int, error) {
	return aw.f.Write(b)
}

// Close() flushes the temporary file, closes it and renames to the original
// filename.
func (aw *atomicWriter) Close() error {
	var ret error
	if err := aw.f.Sync(); err != nil {
		ret = err
	}
	if err := aw.f.Close(); err != nil {
		ret = err
	}
	if err := os.Rename(aw.f.Name(), aw.name); err != nil {
		ret = err
	}
	return ret
}
