package main

import "io"

type Tape struct {
	file io.ReadWriteSeeker
}

func (t *Tape) Write(p []byte) (n int, e error) {
	t.file.Seek(0, 0)
	return t.file.Write(p)
}
