package color

var Out, _ = NewColorableDevice(OutStreamStdout, nil)

type Color struct {
	params []Attribute
}

func (it *Color) Add(v ...Attribute) *Color {
	it.params = append(it.params, v...)
	return it
}

func New() {}

func NewColor() *Color {
	it := new(Color)
	it.params = make([]Attribute, 0)
	return it
}
