# go-color
Simple package that provide colorable output into terminal;

## Installation

```
$ go get github.com/NovelCorpse/go-color
```

## Test

```
$ go test
```

## Usage

```go
package main

import (
	"fmt"
	"os"

	color "github.com/NovelCorpse/go-color"
)

func main() {
	c := color.NewColor()
	c.Add("red").Println("Hello red world!")
	c.Add("yellow").Printf("Hello yellow world!\n")
	c.Add("cyan").Fprintf(os.Stdout, "Hello cyan world!\n")
	magenta := c.Add("magenta").Sprintf("Hello magenta world!\n")
	fmt.Printf("colorable string is: %s", magenta)

	color.Printf("red", "red;\n")
	color.Println("yellow", "yellow;")
	color.Fprintf(os.Stdout, "cyan", "cyan;\n")
	blue := color.Sprintf("blue", "blue;\n")
	fmt.Printf("colorable string is: %s", blue)

	color.Red("It's red;")
	color.Green("It's green;")
	color.Blue("It's blue;")
	color.Cyan("It's cyan;")
	color.Magenta("It's magenta")

	color.Greenf("It's formating green;\n")
	color.Redf("It's formating red;\n")
	color.Yellowf("It's formating yellow;\n")
	color.Bluef("It's formating blue;\n")
	color.Cyanf("It's formating cyan;\n")
	color.Magentaf("It's formating magenta;\n")
	color.Whitef("It's formating white;\n")
	color.Blackf("It's formating black;\n")

	red := color.SRed("It's stringified red;")
	fmt.Printf("%s\n", red)

	color.FGreenf(os.Stdout, "It's formating green to file;\n")
  color.FYellow(os.Stdout, "It's yellow to file;")
}
```



## License

GPL-3.0