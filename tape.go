package main

import (
	"os"
)

type Tape struct {
	file *os.File // io.ReadWriteSeeker
}

func (t *Tape) Write(p []byte) (n int, e error) {
	t.file.Truncate(0)
	t.file.Seek(0, 0)
	return t.file.Write(p)
}
