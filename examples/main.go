package main

import (
	"fmt"

	"github.com/tinystack/tsrand"
)

func main() {
	fmt.Printf("rand int32 is %d\n", rand.Int32())
	fmt.Printf("rand int64 is %d\n", rand.Int64())
	fmt.Printf("rand int is %d\n", rand.Int())

	fmt.Printf("rand normal strings is %s\n", rand.String(40))
	fmt.Printf("rand visible strings is %s\n", rand.VisibleString(40))
	fmt.Printf("uuid is %s\n", rand.UUID())
}
