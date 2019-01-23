package color

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	tty "github.com/NovelCorpse/go-tty"
)

const escape = "\x1b"

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

func (it *Color) format() string {
	return fmt.Sprintf("%s[%sm", escape, it.sequence())
}

func (it *Color) set(w io.Writer) *Color {
	if !isColorable() {
		return it
	}
	if w != nil {
		fmt.Fprintf(w, it.format())
	}

	fmt.Fprintf(Out, it.format())
	return it
}

func (it *Color) reset(w io.Writer) {
	if !isColorable() {
		return
	}
	if w != nil {
		fmt.Fprintf(w, "%s[%dm", escape, AttributeFormatReset)
	}
	fmt.Fprintf(Out, "%s[%dm", escape, AttributeFormatReset)
}

func isColorable() bool { return tty.IsTTY() }

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
	return (it.sset() + fmt.Sprintf(format, a...) + it.sreset())
}

func (it *Color) sset() string {
	return fmt.Sprintf("%s[%sm", escape, it.sequence())
}
func (it *Color) sreset() string {
	return fmt.Sprintf("%s[%dm", escape, AttributeFormatReset)
}

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

func (it *Color) Sprintf(format string, a ...interface{}) string {
	return (it.sset() + fmt.Sprintf(format, a...) + it.sreset())
}

func NewColor() *Color {
	it := new(Color)
	it.params = make([]Attribute, 0)
	return it
}
