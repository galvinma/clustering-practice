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

}
