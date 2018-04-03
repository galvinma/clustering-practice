package louvain

import (
    "log"
)

func SumNetworkLinks(network map[int][]int) int {
    var sum int
    for _,v := range(network) {
        for _,y := range(v) {
            sum += y
          }
    }
    return sum
}


// Sum of the weights of the links incident to node i.
func SumIncidentWeight(network map[int][]int, object int) {
    var sum int
    for k,v := range(network) {
        if k == object {
            for _,y := range(v) {
                sum += y
            }
        }
    }
    return sum
}

// Parse objects in map for a given community and return a member list
func GetMembers(network map[int]int, community int) []int {
    var members []int
    // object ID, community ID
    for k,v := range(network) {
        if v == community {
            members = append(members, k)
        }
    }
    return members
}

/* For a given object (i) in community (C), return the sum of the weights of the links from i to nodes in C. */
func GetCommunityWeights(network map[int]int, community int) int {
    members := GetMembers(network, community)


}

func InitialCommunities(data map[int][]int) map[int]int {
    // Each object is assigned to its own community.
    log.Println("Assigning objects to initial communities..")
    var assigned map[int]int
    assigned = make(map[int]int)
    length := len(data)
    for k,_ := range(data) {
        assigned[k] = k
    }
    return assigned
}
