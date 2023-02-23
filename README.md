# misc-go

A small Go library providing some useful functions.

## Installation

```bash
go get github.com/boatware/misc-go
```

## Usage

> WIP

```go
package main

import (
	"fmt"
	"github.com/boatware/misc-go"
)

func main() {
	// Round a float.
	sampleFLoat := 1.23456789
	rounded := misc_go.Round(sampleFloat, 0.01)
	fmt.Println(rounded) // 1.23
}
```