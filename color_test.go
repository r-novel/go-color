package color

import (
	"fmt"
	"os"
	"testing"
)

func TestColorAttrs(t *testing.T) {
	fmt.Printf("Testing attribute enums \n")

	fmt.Printf("Reset: dec=%d, hex=%x\n", AttributeFormatReset, AttributeFormatReset)
	fmt.Printf("Bold: dec=%d, hex=%x\n", AttributeFormatBold, AttributeFormatBold)
	fmt.Printf("Faint: dec=%d, hex=%x\n", AttributeFormatFaint, AttributeFormatFaint)
	fmt.Printf("CrossedOut: dec=%d, hex=%x\n", AttributeFormatCrossedOut, AttributeFormatCrossedOut)

	fmt.Printf("FG Black: dec=%d, hex=%x\n", AttributeFGColorBlack, AttributeFGColorBlack)
	fmt.Printf("FG Red: dec=%d, hex=%x\n", AttributeFGColorRed, AttributeFGColorRed)
	fmt.Printf("FG Green: dec=%d, hex=%x\n", AttributeFGColorGreen, AttributeFGColorGreen)
	fmt.Printf("FG White: dec=%d, hex=%x\n", AttributeFGColorWhite, AttributeFGColorWhite)
	fmt.Printf("FG Yellow: dec=%d, hex=%x\n", AttributeFGColorYellow, AttributeFGColorYellow)

	fmt.Printf("BG Black: dec=%d, hex=%x\n", AttributeBGColorBlack, AttributeBGColorBlack)
	fmt.Printf("FG White: dec=%d, hex=%x\n", AttributeBGColorWhite, AttributeBGColorWhite)

	fmt.Printf("FGB Black: dec=%d, hex=%x\n", AttributeFGBrightColorBlack, AttributeFGBrightColorBlack)
	fmt.Printf("FGB White: dec=%d, hex=%x\n", AttributeFGBrightColorWhite, AttributeFGBrightColorWhite)

	fmt.Printf("BGB Black: dec=%d, hex=%x\n", AttributeBGBrightColorBlack, AttributeBGBrightColorBlack)
	fmt.Printf("BGB White: dec=%d, hex=%x\n\n", AttributeBGBrightColorWhite, AttributeBGBrightColorWhite)

}

func TestColorablePrintfWithInit(t *testing.T) {
	NewColor().Add("red").Printf("Initialized Printf red;\n")
	NewColor().Add("magenta").Printf("Initialized Printf magenta;\n")
	NewColor().Add("yellow").Printf("Initialized Printf yellow;\n")
	NewColor().Add("cyan").Printf("Initialized Printf cyan;\n\n")
}

func TestColorableFprinfWithInit(t *testing.T) {
	NewColor().Add("red").Fprintf(os.Stdout, "Initialized Fprintf red;\n")
	NewColor().Add("magenta").Fprintf(os.Stdout, "Initialized Fprintf magenta;\n")
	NewColor().Add("yellow").Fprintf(os.Stdout, "Initialized Fprintf yellow;\n")
	NewColor().Add("cyan").Fprintf(os.Stdout, "Initialized Fprintf cyan;\n\n")
}

func TestColorablePrinlnWithInit(t *testing.T) {
	NewColor().Add("red").Println("Initialized Println red;")
	NewColor().Add("magenta").Println("Initialized Println magenta;")
	NewColor().Add("yellow").Println("Initialized Println yellow;")
	NewColor().Add("cyan").Println("Initialized Println cyan;\n")
}

func TestColorableSprinfWithInit(t *testing.T) {
	red := NewColor().Add("red").Sprintf("Initialized Sprintf red;")
	fmt.Printf("colorable: %s\n", red)

	magenta := NewColor().Add("magenta").Sprintf("Initialized Sprintf magenta;")
	fmt.Printf("colorable: %s\n", magenta)

	yellow := NewColor().Add("yellow").Sprintf("Initialized Sprintf yellow;")
	fmt.Printf("colorable: %s\n", yellow)

	cyan := NewColor().Add("cyan").Sprintf("Initialized Sprintf cyan;\n")
	fmt.Printf("colorable: %s\n", cyan)
}

func TestColorablePrintf(t *testing.T) {
	Printf("red", "Printf red;\n")
	Printf("magenta", "Printf magenta\n")
	Printf("yellow", "Yellow\n")
	Printf("cyan", "Cyan\n\n")
}

func TestColorableFprintf(t *testing.T) {
	Fprintf(os.Stdout, "red", "Fprintf red;\n")
	Fprintf(os.Stdout, "magenta", "Fprintf magenta;\n")
	Fprintf(os.Stdout, "yellow", "Fprintf yellow;\n")
	Fprintf(os.Stdout, "cyan", "Fprintf cyan;\n\n")
}

func TestColorablePrintln(t *testing.T) {
	Println("red", "Println red;")
	Println("magenta", "Println magenta;")
	Println("yellow", "Println yellow;")
	Println("cyan", "Println cyan;\n")
}

func TestColorableSprintf(t *testing.T) {
	red := Sprintf("red", "Sprintf red;")
	fmt.Printf("colorable: %s\n", red)

	magenta := Sprintf("magenta", "Sprintf magenta;")
	fmt.Printf("colorable: %s\n", magenta)

	yellow := Sprintf("yellow", "Sprintf yellow;")
	fmt.Printf("colorable: %s\n", yellow)

	cyan := Sprintf("cyan", "Sprintf cyan;")
	fmt.Printf("colorable: %s\n", cyan)
}
