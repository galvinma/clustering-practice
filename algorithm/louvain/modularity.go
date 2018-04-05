package louvain

import (
    // "log"

    // "github.com/galvinma/agora/algorithm/common"
)

/* Modularity Calculation

"The modularity (Q) of a partition is a scalar value betwee -1 and 1 that measures the density of links inside communitiesn as compared to linkes between communities."

For weighted networks:

Q = (1/(2*m)*SUM(Aij-((ki*kj)/(2*m)))*(DELTA(ci,cj))

Where:

Q = modularity of a partition
m = (1/2)*SUM(Aij)
Aij = weight of the edge between i and j
ki = sum of the weights of the links indicent to node i.
kj = sum of the weights of the links indicent to node j.
DELTA(ci,cj) = 1 if ci = cj and 0 otherwise

Gain in modularity (DELTA Q) is obtained by testing the move of an isolated node (i) into another community (C):

DELTA Q = [((SUM,in + ki,in)/(2*m)) - ((SUM,tot + ki)/(2*m))**2] - [((SUM,in)/(2*m)) - ((SUM,tot)/(2*m))**2 - ((ki/(2*m))**2]

Where:

DELTA Q = change in modularity
SUM,in = sum of the weights of the links inside C
SUM,tot = sum of the weights of the links indicent to C
ki = sum of the weights of the links indicent to node i
ki,in = sum of the weights of the links from i to nodes in C
m = sum of the weights of all the links in the network

*/


// Calculate the change in modularity by placing node i into community C.
// func ModularityDelta(data map[int][]int, network map[int]int, weights map[int][]int, community int, object int) float64 {
//     // ki,in = sum of the weights of the links from i to nodes in C
//     SumCluster := GetCommunityWeights(data, network, community, object)
//     // SUM,tot = sum of the weights of the links indicent to C
//     SumIncident := SumIncidentWeight(data, object)
//     // ki,in = sum of the weights of the links from i to nodes in C
//
//
//
//
//     // m = sum of the weights of all the links in the network
//     // This will never change. Move this outside of delta calculation.
//     m := SumNetworkLinks(weights)
//
//
// }

// func Modularity() {
//     // community ID, modularity delta
//     var q map[int]int
//     q = make(map[int]int)
//
//     // For each node, calculate the potential modularity delta for each community.
//     for k,v := range(data) {
//         change := ModularityDelta()
//     }

    // Get max modularity delta out of map.

    // Merge node into community based upon max modularity delta.

// }
