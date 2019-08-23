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

func TestColorableFprintfWithInit(t *testing.T) {
	NewColor().Add("red").Fprintf(os.Stdout, "Initialized Fprintf red;\n")
	NewColor().Add("magenta").Fprintf(os.Stdout, "Initialized Fprintf magenta;\n")
	NewColor().Add("yellow").Fprintf(os.Stdout, "Initialized Fprintf yellow;\n")
	NewColor().Add("cyan").Fprintf(os.Stdout, "Initialized Fprintf cyan;\n\n")
}

func TestColorableFprintlnWithInit(t *testing.T) {
	NewColor().Add("red").Fprintln(os.Stdout, "Initialized Fprintln red;")
	NewColor().Add("magenta").Fprintln(os.Stdout, "Initialized Fprintln magenta;")
	NewColor().Add("yellow").Fprintln(os.Stdout, "Initialized Fprintln yellow;")
	NewColor().Add("cyan").Fprintln(os.Stdout, "Initialized Fprintln cyan;\n")
}

func TestColorablePrintlnWithInit(t *testing.T) {
	NewColor().Add("red").Println("Initialized Println red;")
	NewColor().Add("magenta").Println("Initialized Println magenta;")
	NewColor().Add("yellow").Println("Initialized Println yellow;")
	NewColor().Add("cyan").Println("Initialized Println cyan;\n")
}

func TestColorableSprintfWithInit(t *testing.T) {
	red := NewColor().Add("red").Sprintf("Initialized Sprintf red;")
	fmt.Printf("colorable: %s\n", red)

	magenta := NewColor().Add("magenta").Sprintf("Initialized Sprintf magenta;")
	fmt.Printf("colorable: %s\n", magenta)

	yellow := NewColor().Add("yellow").Sprintf("Initialized Sprintf yellow;")
	fmt.Printf("colorable: %s\n", yellow)

	cyan := NewColor().Add("cyan").Sprintf("Initialized Sprintf cyan;")
	fmt.Printf("colorable: %s\n\n", cyan)
}

func TestColorableSprintlnWithInit(t *testing.T) {
	red := NewColor().Add("red").Sprintln("Initialized Sprintln red;")
	fmt.Printf("colorable: %s", red)

	magenta := NewColor().Add("magenta").Sprintln("Initialized Sprintln magenta;")
	fmt.Printf("colorable: %s", magenta)

	yellow := NewColor().Add("yellow").Sprintln("Initialized Sprintln yellow;")
	fmt.Printf("colorable: %s", yellow)

	cyan := NewColor().Add("cyan").Sprintln("Initialized Sprintln cyan;")
	fmt.Printf("colorable: %s\n", cyan)
}

func TestColorablePrintf(t *testing.T) {
	Printf("red", "Printf red;\n")
	Printf("magenta", "Printf magenta;\n")
	Printf("yellow", "Printf yellow;\n")
	Printf("cyan", "Printf cyan;\n\n")
}

func TestColorableFprintf(t *testing.T) {
	Fprintf(os.Stdout, "red", "Fprintf red;\n")
	Fprintf(os.Stdout, "magenta", "Fprintf magenta;\n")
	Fprintf(os.Stdout, "yellow", "Fprintf yellow;\n")
	Fprintf(os.Stdout, "cyan", "Fprintf cyan;\n\n")
}

func TestColorableFprintln(t *testing.T) {
	Fprintln(os.Stdout, "red", "Fprintln red;")
	Fprintln(os.Stdout, "magenta", "Fprintln magenta;")
	Fprintln(os.Stdout, "yellow", "Fprintln yellow;")
	Fprintln(os.Stdout, "cyan", "Fprintln cyan;\n")
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
	fmt.Printf("colorable: %s\n\n", cyan)
}

func TestColorableSprintln(t *testing.T) {
	red := Sprintln("red", "Sprintln red;")
	fmt.Printf("colorable: %s", red)

	magenta := Sprintln("magenta", "Sprintln magenta;")
	fmt.Printf("colorable: %s", magenta)

	yellow := Sprintln("yellow", "Sprintln yellow;")
	fmt.Printf("colorable: %s", yellow)

	cyan := Sprintln("cyan", "Sprintln cyan;")
	fmt.Printf("colorable: %s\n", cyan)
}

func TestColorablePrintlnByName(t *testing.T) {
	Green("It's green;")
	Red("It's red;")
	Yellow("It's yellow;")
	Blue("It's blue;")
	Cyan("It's cyan;")
	Magenta("It's magenta;")
	White("It's white;")
	Black("It's black;")
}

func TestColorablePrintfByName(t *testing.T) {
	Greenf("It's formating green;\n")
	Redf("It's formating red;\n")
	Yellowf("It's formating yellow;\n")
	Bluef("It's formating blue;\n")
	Cyanf("It's formating cyan;\n")
	Magentaf("It's formating magenta;\n")
	Whitef("It's formating white;\n")
	Blackf("It's formating black;\n")
}

func TestColorableFPrintfByName(t *testing.T) {
	FGreenf(os.Stdout, "It's formating green to file;\n")
	FRedf(os.Stdout, "It's formating red to file;\n")
	FYellowf(os.Stdout, "It's formating yellow to file;\n")
	FBluef(os.Stdout, "It's formating blue to file;\n")
	FCyanf(os.Stdout, "It's formating cyan to file;\n")
	FMagentaf(os.Stdout, "It's formating magenta to file;\n")
	FWhitef(os.Stdout, "It's formating white to file;\n")
	FBlackf(os.Stdout, "It's formating black to file;\n")
}

func TestColorableFPrintlnByName(t *testing.T) {
	FGreen(os.Stdout, "It's green to file;")
	FRed(os.Stdout, "It's red to file;")
	FYellow(os.Stdout, "It's yellow to file;")
	FBlue(os.Stdout, "It's blue to file;")
	FCyan(os.Stdout, "It's cyan to file;")
	FMagenta(os.Stdout, "It's magenta to file;")
	FWhite(os.Stdout, "It's white to file;")
	FBlack(os.Stdout, "It's black to file;")
}

func TestColorableSprintlnByName(t *testing.T) {
	red := SRed("Sprintln red;")
	fmt.Printf("colorable: %s", red)

	magenta := SMagenta("Sprintln magenta;")
	fmt.Printf("colorable: %s", magenta)

	yellow := SYellow("Sprintln yellow;")
	fmt.Printf("colorable: %s", yellow)

	cyan := SCyan("Sprintln cyan;")
	fmt.Printf("colorable: %s\n", cyan)
}

func TestColorableSprintfByName(t *testing.T) {
	red := SRedf("It's formatting sprint red;\n")
	fmt.Printf("colorable: %s", red)

	magenta := SMagentaf("It's formatting sprint magenta;\n")
	fmt.Printf("colorable: %s", magenta)

	yellow := SYellowf("It's formatting sprint yellow;\n")
	fmt.Printf("colorable: %s", yellow)

	cyan := SCyanf("It's formatting sprint cyan;\n")
	fmt.Printf("colorable: %s\n", cyan)
}
