package main

import (
	"fmt"

	"github.com/noltedennis/ambilight"
	"github.com/noltedennis/ambilight/internal/clientlib"
)

func main() {
	fmt.Println(ambilight.Config())
	fmt.Println(clientlib.Hello())
}
