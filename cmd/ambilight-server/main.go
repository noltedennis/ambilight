package main

import "fmt"
import "github.com/noltedennis/ambilight"
import "github.com/noltedennis/ambilight/serverlib"

func main() {
  fmt.Println(ambilight.Config())
  fmt.Println(serverlib.Hello())
}
