package color

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	tty "github.com/NovelCorpse/go-tty"
)

var Out, _ = NewColorableDevice(OutStreamStdout, nil)

type Color struct {
	params []Attribute
}

func (it *Color) Add(in ...string) *Color {
	for _, k := range in {
		if v, ok := stringAttribute[k]; ok {
			it.params = append(it.params, v)
		}
	}

	return it
}

func (it *Color) sequence() string {
	format := make([]string, len(it.params))
	for i, v := range it.params {
		format[i] = strconv.Itoa(int(v))
	}

	return strings.Join(format, ";")
}

func (it *Color) fmt() string {
	return fmt.Sprintf("%s[%sm", oct_escape, it.sequence())
}

func (it *Color) unfmt() string {
	return fmt.Sprintf("%s[%dm", hex_escape, AttributeFormatReset)
}

func (it *Color) wrap(in string) string {
	return fmt.Sprintf("%s%s%s", it.fmt(), in, it.unfmt())
}

func (it *Color) set(w io.Writer) *Color {
	if !isColorable() {
		return it
	}
	if w != nil {
		fmt.Fprintf(w, it.fmt())
	}

	fmt.Fprintf(Out, it.fmt())
	return it
}

func (it *Color) reset(w io.Writer) {
	if !isColorable() {
		return
	}
	if w != nil {
		fmt.Fprintf(w, it.unfmt())
	}
	fmt.Fprintf(Out, it.unfmt())
}

func isColorable() bool { return tty.IsTTY() }

func (it *Color) Printf(format string, a ...interface{}) (int, error) {
	it.set(nil)
	defer it.reset(nil)

	return fmt.Fprintf(Out, format, a...)
}

func (it *Color) Println(a ...interface{}) (int, error) {
	it.set(nil)
	defer it.reset(nil)

	return fmt.Fprintln(Out, a...)
}

func (it *Color) Fprintf(w io.Writer, format string, a ...interface{}) (int, error) {
	it.set(w)
	defer it.reset(w)

	return fmt.Fprintf(w, format, a...)
}

func (it *Color) Fprintln(w io.Writer, a ...interface{}) (int, error) {
	it.set(w)
	defer it.reset(w)

	return fmt.Fprintln(w, a...)
}

func (it *Color) Sprintf(format string, a ...interface{}) string {
	return it.wrap(fmt.Sprintf(format, a...))
}

func (it *Color) Sprintln(a ...interface{}) string {
	return it.wrap(fmt.Sprintln(a...))
}

func NewColor() *Color {
	it := new(Color)
	it.params = make([]Attribute, 0)
	return it
}
