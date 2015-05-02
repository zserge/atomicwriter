package atomicwriter

import (
	"io/ioutil"
	"os"
	"testing"
)

func write(t *testing.T, name, s string) {
	w, err := NewWriter(name)
	if err != nil {
		t.Error(err)
	}
	if _, err := w.Write([]byte(s)); err != nil {
		t.Error(err)
	}
	if err := w.Close(); err != nil {
		t.Error(err)
	}
}

func check(t *testing.T, name, s string) {
	if b, err := ioutil.ReadFile(name); err != nil {
		t.Error(err)
	} else if string(b) != s {
		t.Error("Unexpected contents:", b, []byte(s))
	}
}

// Write string to file, then read it
func TestWriterSimple(t *testing.T) {
	write(t, "test.txt", "test string")
	check(t, "test.txt", "test string")
	os.Remove("test.txt")
}

// Start one writer, make another write operation in the middle, finish the
// first writer.
func TestWriterPartialWrite(t *testing.T) {
	w, _ := NewWriter("test.txt")
	w.Write([]byte("first part\n"))
	write(t, "test.txt", "another write")
	check(t, "test.txt", "another write")
	w.Write([]byte("second part\n"))
	w.Close()
	check(t, "test.txt", "first part\nsecond part\n")
	os.Remove("test.txt")
}
