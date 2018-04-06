package louvain

import (
    "log"

    "github.com/galvinma/agora/algorithm/common"
)

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



// // Sum of the weights of the links incident to node i.
// func SumIncidentWeight(weighted [][]int , object int) int {
//     var sum int
//     for k,v := range(data) {
//         if k == object {
//             for _,y := range(v) {
//                 sum += y
//             }
//             break
//         }
//     }
//     return sum
// }

// // Sum of the weights between all nodes in the network.
// // Accepts a symmetric weight matrix
// func SumNetworkLinks(network map[int][]int) int {
//     var sum int
//     for _,v := range(network) {
//         for _,y := range(v) {
//             sum += y
//           }
//     }
//     return sum
// }


// /* For a given node (i) in community (C), return the sum of the weights of the links from i to nodes in C. */
// func GetCommunityWeights(data map[int[]int, network map[int]int, community int, object int) int {
//     /* Note, community variable IS NOT the current community of object. */
//     var sum int
//     members := GetMembers(network, community)
//     // Get attributes for object
//     i := data[object]
//     for _,m := range(members) {
//         // Calculate hamming distance
//         dif, err := common.HammingDistance(i, network[m])
//         if err != nil {
//             log.Fatal(err)
//         }
//         sum += dif
//     }
//     return sum
// }
