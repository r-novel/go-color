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
	color.Println(color.AttributeFGColorWhite, "Hello white world!!!")
	color.Printf(color.AttributeFGColorYellow, "Hello yellow world!!!\n")
}

```



## License

GPL-3.0