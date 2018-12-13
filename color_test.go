package color

import (
	"fmt"
	"testing"
)

func TestColor(t *testing.T) {
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
