package main

import (
    "log"

    "github.com/galvinma/agora/algorithm"
)

func main() {

  soybeans := hierdenc.InitSoybeanLarge()

  densest, index := hierdenc.HierdencIndex(soybeans, 1)
  for key, value := range index {
      log.Println("Key:", key, "Value:", value)
  }
  log.Println(densest)
}
