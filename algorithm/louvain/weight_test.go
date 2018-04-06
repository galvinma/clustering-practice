package louvain

import (
    "log"

    "github.com/galvinma/agora/algorithm/common"
)

func TestHappyPath(t *testing.T) {
    data := common.InitSimple()
    network := louvain.InitialCommunities(data)
    weights := louvain.GetWeights(data)
}
// sumin := louvain.SumCommunityWeights(data, network, weights, 1)
// sumtot := louvain.SumIncidentCommunityWeights(data, network, weights, 1)
// ki := louvain.GetIncidentWeights(data, weights, 1)
// kiin := louvain.GetIncidentCommunityWeights(network, weights, 1, 1)
// m := louvain.SumNetworkLinks(weights)
//
// log.Println("sumin:", sumin, "sumtot:", sumtot, "ki:", ki, "kiin:", kiin, "m:", m)
