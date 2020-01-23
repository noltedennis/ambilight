package main

import (
	"fmt"

	"github.com/noltedennis/ambilight"
	"github.com/noltedennis/ambilight/serverlib"
)

func main() {
	fmt.Println(ambilight.Config())
	fmt.Println(serverlib.Hello())
}
