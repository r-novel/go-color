package color

import (
	"fmt"
	"io"
)

func Printf(v string, format string, a ...interface{}) (int, error) {
	it := NewColor()
	it.Add(v)
	it.set(nil)
	defer it.reset(nil)

	return fmt.Fprintf(Out, format, a...)
}

func Fprintf(w io.Writer, v string, format string, a ...interface{}) (int, error) {
	it := NewColor()
	it.Add(v)
	it.set(w)
	defer it.reset(w)

	return fmt.Fprintf(w, format, a...)
}

func Println(v string, a ...interface{}) (int, error) {
	it := NewColor()
	it.Add(v)
	it.set(nil)
	defer it.reset(nil)

	return fmt.Fprintln(Out, a...)
}

func Sprintf(v string, format string, a ...interface{}) string {
	it := NewColor()
	it.Add(v)
	return (it.fmt() + fmt.Sprintf(format, a...) + it.unfmt())
}
