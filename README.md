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

import color "github.com/NovelCorpse/go-color"

func main() {
	c := color.NewColor()
	c.Add("red").Println("Hello red world!")
	c.Add("yellow").Printf("Hello yellow world!\n")

	color.Printf("red", "Without color obj\n")
	color.Println("yellow", "Without color obj\n")
}

```



## License

GPL-3.0