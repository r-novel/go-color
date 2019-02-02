package color

import "io"

func Red(a ...interface{}) (int, error)                 { return Println("red", a...) }
func Redf(format string, a ...interface{}) (int, error) { return Printf("red", format, a...) }
func SRed(a ...interface{}) string                      { return Sprintln("red", a...) }
func FRed(w io.Writer, a ...interface{}) (int, error)   { return Fprintln(w, "red", a...) }
func FRedf(w io.Writer, format string, a ...interface{}) (int, error) {
	return Fprintf(w, "red", format, a...)
}

func Green(a ...interface{}) (int, error)                 { return Println("green", a...) }
func Greenf(format string, a ...interface{}) (int, error) { return Printf("green", format, a...) }
func SGreen(a ...interface{}) string                      { return Sprintln("green", a...) }
func FGreen(w io.Writer, a ...interface{}) (int, error)   { return Fprintln(w, "green", a...) }
func FGreenf(w io.Writer, format string, a ...interface{}) (int, error) {
	return Fprintf(w, "green", format, a...)
}

func Yellow(a ...interface{}) (int, error)                 { return Println("yellow", a...) }
func Yellowf(format string, a ...interface{}) (int, error) { return Printf("yellow", format, a...) }
func SYellow(a ...interface{}) string                      { return Sprintln("yellow", a...) }
func FYellow(w io.Writer, a ...interface{}) (int, error)   { return Fprintln(w, "yellow", a...) }
func FYellowf(w io.Writer, format string, a ...interface{}) (int, error) {
	return Fprintf(w, "yellow", format, a...)
}

func Blue(a ...interface{}) (int, error)                 { return Println("blue", a...) }
func Bluef(format string, a ...interface{}) (int, error) { return Printf("blue", format, a...) }
func SBlue(a ...interface{}) string                      { return Sprintln("blue", a...) }
func FBlue(w io.Writer, a ...interface{}) (int, error)   { return Fprintln(w, "blue", a...) }
func FBluef(w io.Writer, format string, a ...interface{}) (int, error) {
	return Fprintf(w, "blue", format, a...)
}

func Cyan(a ...interface{}) (int, error)                 { return Println("cyan", a...) }
func Cyanf(format string, a ...interface{}) (int, error) { return Printf("cyan", format, a...) }
func SCyan(a ...interface{}) string                      { return Sprintln("cyan", a...) }
func FCyan(w io.Writer, a ...interface{}) (int, error)   { return Fprintln(w, "cyan", a...) }
func FCyanf(w io.Writer, format string, a ...interface{}) (int, error) {
	return Fprintf(w, "cyan", format, a...)
}

func Magenta(a ...interface{}) (int, error)                 { return Println("magenta", a...) }
func Magentaf(format string, a ...interface{}) (int, error) { return Printf("magenta", format, a...) }
func SMagenta(a ...interface{}) string                      { return Sprintln("magenta", a...) }
func FMagenta(w io.Writer, a ...interface{}) (int, error)   { return Fprintln(w, "magenta", a...) }
func FMagentaf(w io.Writer, format string, a ...interface{}) (int, error) {
	return Fprintf(w, "magenta", format, a...)
}

func White(a ...interface{}) (int, error)                 { return Println("white", a...) }
func Whitef(format string, a ...interface{}) (int, error) { return Printf("white", format, a...) }
func SWhite(a ...interface{}) string                      { return Sprintln("white", a...) }
func FWhite(w io.Writer, a ...interface{}) (int, error)   { return Fprintln(w, "white", a...) }
func FWhitef(w io.Writer, format string, a ...interface{}) (int, error) {
	return Fprintf(w, "white", format, a...)
}

func Black(a ...interface{}) (int, error)                 { return Println("black", a...) }
func Blackf(format string, a ...interface{}) (int, error) { return Printf("black", format, a...) }
func SBlack(a ...interface{}) string                      { return Sprintln("black", a...) }
func FBlack(w io.Writer, a ...interface{}) (int, error)   { return Fprintln(w, "black", a...) }
func FBlackf(w io.Writer, format string, a ...interface{}) (int, error) {
	return Fprintf(w, "black", format, a...)
}
