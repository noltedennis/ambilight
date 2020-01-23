package main

import "fmt"
import "github.com/noltedennis/ambilight"
import "github.com/noltedennis/ambilight/clientlib"

func main() {
  fmt.Println(ambilight.Config())
  fmt.Println(clientlib.Hello())
}
