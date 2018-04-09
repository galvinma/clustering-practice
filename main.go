package main

import (
    "log"

    "github.com/galvinma/agora/algorithm/common"
    "github.com/galvinma/agora/algorithm/louvain"
)

func main() {
  log.Println("Initializing data...")
  data := common.InitSimple()
  network := louvain.InitialCommunities(data)
  weights := louvain.GetWeights(data)

  assigned := louvain.AssignCommunities(data, network, weights)
  for k,v := range(assigned) {
      log.Println("Object",k,"assigned to cluster",v)
  }
}
