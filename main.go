package main

import (
    "log"

    "github.com/galvinma/agora/algorithm/common"
    "github.com/galvinma/agora/algorithm/louvain"
)

func main() {
  log.Println("Initializing data...")
  data := common.InitSoybeanLarge()
  network := louvain.InitialCommunities(data)
  weights := louvain.GetWeights(data)
  // for k,v := range(network) {
  //     log.Println("For object",k,"the assigned cluster is",v)
  // }

}
