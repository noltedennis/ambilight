package main

import (
	"fmt"
	"log"

	"github.com/noltedennis/ambilight"
	"github.com/noltedennis/ambilight/serverlib"
	"github.com/stianeikeland/go-rpio/v4"
)

func main() {
	fmt.Println(ambilight.Config())
	fmt.Println(serverlib.Hello())

	err := rpio.Open()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("RPIO device opened")
	rpio.Close()
}
