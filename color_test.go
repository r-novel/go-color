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
	NewColor().Add("red").Printf("initialized Red\n")
	NewColor().Add("magenta").Printf("initialized Magenta\n")
	NewColor().Add("yellow").Printf("initialized Yellow\n")
	NewColor().Add("cyan").Printf("initialized Cyan\n")
}

func TestColorablePrintf(t *testing.T) {
	Printf("red", "Red\n")
	Printf("magenta", "Magenta\n")
	Printf("yellow", "Yellow\n")
	Printf("cyan", "Cyan\n")
}

func TestColorableBackgroundPrintf(t *testing.T) {
	Printf("b-red", "Background red\n")
	Printf("b-magenta", "Background magenta\n")
	Printf("b-yellow", "Background yellow\n")
	Printf("b-cyan", "Background cyan\n")
}

func TestColorablePrinlnWithInit(t *testing.T) {
	NewColor().Add("red").Println("initialized Println Red")
	NewColor().Add("magenta").Println("initialized Println Magenta")
	NewColor().Add("yellow").Println("initialized Println Yellow")
	NewColor().Add("cyan").Println("initialized Println Cyan")
}

func TestColorablePrintln(t *testing.T) {
	Println("red", "Println Red")
	Println("magenta", "Println Magenta")
	Println("yellow", "Println Yellow")
	Println("cyan", "Println Cyan")

}
