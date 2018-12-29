package color

import (
	"fmt"
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
	fmt.Printf("BGB White: dec=%d, hex=%x\n", AttributeBGBrightColorWhite, AttributeBGBrightColorWhite)

}

func TestColorablePrintfWithInit(t *testing.T) {
	NewColor().Add(AttributeFGColorRed).Printf("initialized Red\n")
	NewColor().Add(AttributeFGColorMagenta).Printf("initialized Magenta\n")
	NewColor().Add(AttributeFGColorYellow).Printf("initialized Yellow\n")
	NewColor().Add(AttributeFGColorCyan).Printf("initialized Cyan\n")
}

func TestColorablePrintf(t *testing.T) {
	Printf(AttributeFGColorRed, "Red\n")
	Printf(AttributeFGColorMagenta, "Magenta\n")
	Printf(AttributeFGColorYellow, "Yellow\n")
	Printf(AttributeFGColorCyan, "Cyan\n")
}

func TestColorablePrinlnWithInit(t *testing.T) {
	NewColor().Add(AttributeFGColorRed).Println("initialized Println Red")
	NewColor().Add(AttributeFGColorMagenta).Println("initialized Println Magenta")
	NewColor().Add(AttributeFGColorYellow).Println("initialized Println Yellow")
	NewColor().Add(AttributeFGColorCyan).Println("initialized Println Cyan")
}

func TestColorablePrintln(t *testing.T) {
	Println(AttributeFGColorRed, "Println Red")
	Println(AttributeFGColorMagenta, "Println Magenta")
	Println(AttributeFGColorYellow, "Println Yellow")
	Println(AttributeFGColorCyan, "Println Cyan")

}
