package louvain

import (
    "math"
    "log"
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
func ModularityDelta(data map[int][]int, network map[int]int, weights [][]int, community int, object int) float64 {
    sumin := SumCommunityWeights(network, weights, community)
    sumtot := SumIncidentCommunityWeights(data, network, weights, community)
    ki := GetIncidentWeights(data, weights, object)
    kiin := GetIncidentCommunityWeights(network, weights, community, object)
    m := SumNetworkLinks(weights)

    q := (((sumin + kiin) / (2 * m)) - (math.Pow(((sumtot + ki) / (2 * m)), 2))) - ((sumin / (2 * m)) - math.Pow((sumtot / (2 * m)), 2) - math.Pow((ki / (2 * m)), 2))
    // if community == 1 && object == 0 {
        log.Println("Q for placing object",object,"in community",community,"is",q)
        log.Println("sumin:",sumin,"sumtot:",sumtot,"ki:",ki,"kiin:",kiin,"m:",m)
    // }


    return q

}

func AssignCommunities(data map[int][]int, network map[int]int, weights [][]int) map[int]int {
    // Obj ID : Cluster ID
    coms := []int{0,1,2,3,4,5,6,7,8,9}
    // For each node, calculate the potential modularity delta for each community.
    for k,_ := range(data) {
        // community ID, modularity delta
        var best int
        modularity := -1.0

        for _,c := range(coms) {
            var test map[int]int
            test = make(map[int]int)
            for a,b := range(network) {
                test[a] = b
            }
            test[k] = c
            delta := ModularityDelta(data, test, weights, c, k)
            if delta > modularity && delta > 0.0 {
                modularity = delta
                network[k] = best
            }
        }

        log.Println("Determined best fit for object",k,"is community",network[k])
    }
    return network
}
