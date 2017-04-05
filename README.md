# Quantum Entanglement Generation Device
A Go library which will generate quantum statements ported from http://sebpearce.com/bullshit/.

## Example generation

> As you grow, you will enter into infinite balance that transcends understanding. Illusion is born in the gap where consciousness has been excluded. Reality has always been bursting with warriors whose dreams are opened by life-force. We must learn how to lead enlightened lives in the face of suffering.

## Usage

```go
package main

import (
	"fmt"

	"github.com/JaTochNietDan/quantum"
)

func main() {
	fmt.Println("Generated:", quantum.Generate(4))
}
```
