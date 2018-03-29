package main

import (
    "log"

    "github.com/galvinma/agora/algorithm"
)

func main() {

  soybeans := hierdenc.InitSoybeanLarge()
  clusters := hierdenc.HIERDENC(soybeans)
  for k, v := range clusters {
    log.Println("Object ID: ", k, "Cluster ID: ", v)
  }
}
