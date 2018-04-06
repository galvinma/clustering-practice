package main

import (
    "log"

    "github.com/galvinma/agora/algorithm/common"
    "github.com/galvinma/agora/algorithm/louvain"
)

func main() {
  log.Println("Initializing data...")
  data := common.InitSoybeanLarge()
  // weighted := louvain.GetWeights(data)
  network := louvain.InitialCommunities(data)
  for k,v := range(network) {
      log.Println("For object",k,"the assigned cluster is",v)
  }

}
