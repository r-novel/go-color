package color

import (
	"fmt"
	"io"
	"os"
)

type OutStream uint8

const (
	OutStreamUnknown = 0x0
	OutStreamStdout  = 0x1
	OutStreamStderr  = 0x2
	OutFile          = 0x3
)

func NewColorableDevice(kind OutStream, fd *os.File) (out io.Writer, e error) {
	switch kind {
	case OutStreamStdout:
		return os.Stdout, e
	case OutStreamStderr:
		return os.Stderr, e
	case OutStreamUnknown:
		return nil, fmt.Errorf("unknown output stream;")
	case OutFile:
		if fd != nil {
			return fd, e
		} else {
			return nil, fmt.Errorf("file not found;")
		}

	}
	return
}
