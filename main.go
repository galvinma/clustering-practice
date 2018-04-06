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

  // test
  sumin := louvain.SumCommunityWeights(data, network, weights, 1)
  sumtot := louvain.SumIncidentCommunityWeights(data, network, weights, 1)
  ki := louvain.GetIncidentWeights(data, weights, 1)
  kiin := louvain.GetIncidentCommunityWeights(network, weights, 1, 1)
  m := louvain.SumNetworkLinks(weights)

  log.Println("sumin:", sumin, "sumtot:", sumtot, "ki:", ki, "kiin:", kiin, "m:", m)

}
