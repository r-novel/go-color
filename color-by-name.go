package color

func Red(a ...interface{}) (int, error) {
	return Println("red", a...)
}

func Redf(format string, a ...interface{}) (int, error) {
	return Printf("red", format, a...)
}

func Green(a ...interface{}) (int, error) {
	return Println("green", a...)
}

func Greenf(format string, a ...interface{}) (int, error) {
	return Printf("green", format, a...)
}

func Yellow(a ...interface{}) (int, error) {
	return Println("yellow", a...)
}

func Yellowf(format string, a ...interface{}) (int, error) {
	return Printf("yellow", format, a...)
}

func Blue(a ...interface{}) (int, error) {
	return Println("blue", a...)
}

func Bluef(format string, a ...interface{}) (int, error) {
	return Printf("blue", format, a...)
}

func Cyan(a ...interface{}) (int, error) {
	return Println("cyan", a...)
}

func Cyanf(format string, a ...interface{}) (int, error) {
	return Printf("cyan", format, a...)
}

func Magenta(a ...interface{}) (int, error) {
	return Println("magenta", a...)
}

func Magentaf(format string, a ...interface{}) (int, error) {
	return Printf("magenta", format, a...)
}

func White(a ...interface{}) (int, error) {
	return Println("white", a...)
}

func Whitef(format string, a ...interface{}) (int, error) {
	return Printf("white", format, a...)
}

func Black(a ...interface{}) (int, error) {
	return Println("black", a...)
}

func Blackf(format string, a ...interface{}) (int, error) {
	return Printf("black", format, a...)
}
