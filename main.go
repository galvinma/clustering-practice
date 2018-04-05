package main

import (
    "log"

    "github.com/galvinma/agora/algorithm/common"
    "github.com/galvinma/agora/algorithm/louvain"


)

func main() {
  log.Println("Initializing data...")
  data := common.InitSoybeanLarge()
  weighted := louvain.GetWeights(soybeans)
  network := louvain.InitialCommunities(soybeans)
  // for k,v := range(weighted) {
  //     log.Println("For object",k,"the weight vector is",v)
  // }

}
