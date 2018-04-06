package louvain

import (
    "log"

    "github.com/galvinma/agora/algorithm/common"
)

// Create a matrix of node weights
func GetWeights(data map[int][]int) [][]int {
    log.Println("Getting vertex weights...")
    weights := make([][]int, len(data))
    for i := 0; i < len(data); i++ {
        weights[i] = make([]int, len(data))
    }
    for k,v := range(data) {
        for m,n := range(data) {
            dif, err := common.HammingDistance(v,n)
            if err != nil {
                log.Fatal(err)
            }
            weights[k][m] = dif
        }
    }
    return weights
}

// SUM,in
// Sum of the weights for a given community
func SumCommunityWeights(data map[int][]int, network map[int]int, weights [][]int, community int) float64 {
    var sum float64
    members := GetMembers(network, community)

    for _,m := range(members) {
        for _,j := range(members) {
            sum += float64(weights[m][j])
        }
    }
    // Symmetric array, get half of weight sum
    sum = sum/2.0
    return sum
}

// SUM,tot
// Sum of the weights of the links incident to nodes in a community
func SumIncidentCommunityWeights(data map[int][]int, network map[int]int, weights [][]int, community int) float64 {
    var sum float64
    members := GetMembers(network, community)
    communityweights := SumCommunityWeights(data, network, weights, community)
    for _,m := range(members) {
        for k,_ := range(data) {
            sum += float64(weights[m][k])
        }
    }
    // Sum columns, subtract community weights
    sum = (sum/2.0) - communityweights
    return sum
}


// ki
// Sum of the weights of the links incident to node i.
func GetIncidentWeights(data map[int][]int, weights [][]int, object int) float64 {
    var sum float64
    for k,_ := range(data) {
        sum += float64(weights[object][k])
    }
    return sum
}


// ki,in
/* For a given node (i) in community (C), return the sum of the weights of the links from i to nodes in C. */
func GetIncidentCommunityWeights(network map[int]int, weights [][]int, community int, object int) float64 {
    /* Note, community variable IS NOT the current community of object. */
    var sum float64
    members := GetMembers(network, community)
    // Get attributes for object
    for _,m := range(members) {
        sum += float64(weights[object][m])
    }
    return sum
}

// m
// Sum of the weights between all nodes in the network.
func SumNetworkLinks(weights [][]int) float64 {
    var sum float64
    for _,v := range(weights) {
        for _,y := range(v) {
            sum += float64(y)
          }
    }
    sum = sum/2.0
    return sum
}
