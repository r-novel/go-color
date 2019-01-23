# go-color
Simple package that provide colorable output into terminal;

## Installation

```
$ go get github.com/NovelCorpse/go-color
```

## Test

```
go test
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
}
```



## License

GPL-3.0