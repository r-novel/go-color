package color

import (
	"fmt"
	"strconv"
	"strings"

	tty "github.com/NovelCorpse/go-tty"
)

const escape = "\x1b"

var Out, _ = NewColorableDevice(OutStreamStdout, nil)

type Color struct {
	params []Attribute
}

func (it *Color) Add(v ...Attribute) *Color {
	it.params = append(it.params, v...)
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

func (it *Color) set() *Color {
	if !isColorable() {
		return it
	}
	fmt.Fprintf(Out, it.format())
	return it
}

func (it *Color) reset() {
	if !isColorable() {
		return
	}
	fmt.Fprintf(Out, "%s[%dm", escape, AttributeFormatReset)
}

func isColorable() bool { return tty.IsTTY() }

func Printf(v Attribute, format string, a ...interface{}) (int, error) {
	it := NewColor()
	it.Add(v)
	it.set()
	defer it.reset()

	return fmt.Fprintf(Out, format, a...)
}

func Println(v Attribute, a ...interface{}) (int, error) {
	it := NewColor()
	it.Add(v)
	it.set()
	defer it.reset()

	return fmt.Fprintln(Out, a...)
}

func (it *Color) Printf(format string, a ...interface{}) (int, error) {
	it.set()
	defer it.reset()

	return fmt.Fprintf(Out, format, a...)
}

func NewColor() *Color {
	it := new(Color)
	it.params = make([]Attribute, 0)
	return it
}
