package main

import (
    "log"

    "github.com/galvinma/agora/algorithm"
)

func main() {
  log.Println("Initializing data...")
  soybeans := hierdenc.InitSoybeanLarge()
  clusters := hierdenc.HIERDENC(soybeans)
  var unique []int
  for k, v := range clusters {
    log.Println("Object ID: ", k, "Cluster ID: ", v)
    if hierdenc.InList(v, unique) == false {
        unique = append(unique, v)
    }
  }
  count := len(unique)
  log.Println("Clustered", len(clusters),"/", len(soybeans), "objects into", count, "clusters.")
}
