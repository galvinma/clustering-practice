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

  sumin := louvain.SumCommunityWeights(network, weights, 1)
  sumtot := louvain.SumIncidentCommunityWeights(data, network, weights, 1)
  ki := louvain.GetIncidentWeights(data, weights, 1)
  kiin := louvain.GetIncidentCommunityWeights(network, weights, 1, 1)
  m := louvain.SumNetworkLinks(weights)

  log.Println("sumin:", sumin, "sumtot:", sumtot, "ki:", ki, "kiin:", kiin, "m:", m)

  log.Println(weights)
}
