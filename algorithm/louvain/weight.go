package louvain

import (
    "log"

    "github.com/galvinma/agora/algorithm/common"
)

func GetWeights(data map[int][]int) map[int][]int {
    log.Println("Getting vertex weights...")
    var weights map[int][]int
    weights = make(map[int][]int)

    for k,v := range(data) {
        var hd []int
        for _,m := range(data) {
            dif, err := common.HammingDistance(v,m)
            if err != nil {
                log.Fatal(err)
            }
            hd = append(hd, dif)
        }
        weights[k] = hd
    }
    return weights
}

// Sum of the weights of the links incident to node i.
func SumIncidentWeight(data map[int][]int, object int) int {
    var sum int
    for k,v := range(data) {
        if k == object {
            for _,y := range(v) {
                sum += y
            }
            break
        }
    }
    return sum
}

// Sum of the weights between all nodes in the network.
// Accepts a symmetric weight matrix
func SumNetworkLinks(network map[int][]int) int {
    var sum int
    for _,v := range(network) {
        for _,y := range(v) {
            sum += y
          }
    }
    return sum
}

// Sum of the weights for a given community
func SumCommunityWeights(data map[int[]int, community int) int {
    var sum int
    members := GetMembers(network, community)
    for _,m := range(members) {
        attributes := data[m]
        for _,n := range(members) {
            dif, err := common.HammingDistance(attributes,n)
            if err != nil {
                log.Fatal(err)
            }
            hd = append(hd, dif)

        }



        }



        // Calculate hamming distance
        dif, err := common.HammingDistance(i, network[m])
        if err != nil {
            log.Fatal(err)
        }
        sum += dif
    }
    return sum
}

/* For a given node (i) in community (C), return the sum of the weights of the links from i to nodes in C. */
func GetCommunityWeights(data map[int[]int, network map[int]int, community int, object int) int {
    /* Note, community variable IS NOT the current community of object. */
    var sum int
    members := GetMembers(network, community)
    // Get attributes for object
    i := data[object]
    for _,m := range(members) {
        // Calculate hamming distance
        dif, err := common.HammingDistance(i, network[m])
        if err != nil {
            log.Fatal(err)
        }
        sum += dif
    }
    return sum
}
