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

func isColorable() bool {
	if tty.IsTTY() {
		return true
	}
	return false
}

func Set(v ...Attribute) *Color {
	it := NewColor()
	it.Add(v...)
	if !isColorable() {
		return it
	}
	fmt.Fprintf(Out, it.format())
	return it
}

func (it *Color) set() *Color {
	if !isColorable() {
		return it
	}
	fmt.Fprintf(Out, it.format())
	return it
}

func (it *Color) Printf(format string, a ...interface{}) (n int, err error) {
	it.set()

	return fmt.Fprintf(Out, format, a...)
}

func NewColor() *Color {
	it := new(Color)
	it.params = make([]Attribute, 0)
	return it
}
