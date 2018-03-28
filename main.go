package main

import (
    "log"

    "github.com/galvinma/agora/algorithm"
)

func main() {

  soybeans := hierdenc.InitSoybeanLarge()
  index := hierdenc.HierdencIndex(soybeans, 1)
  log.Println(index)

}
