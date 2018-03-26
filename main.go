package main

import (
    "log"

    // "gonum.org/v1/gonum/mat"
    "github.com/galvinma/agora/algorithm"
)

func main() {

  soybeans := hierdenc.InitSoybeanLarge()

  // Return a map of this data
  for key, value := range soybeans {
    log.Println("Key:", key, "Value:", value)
  }

}
