package louvain

import (
    "log"

)

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

// Assign each node to its own community.
func InitialCommunities(data map[int][]int) map[int]int {
    // Each object is assigned to its own community.
    log.Println("Assigning objects to initial communities..")
    var assigned map[int]int
    assigned = make(map[int]int)
    for k,_ := range(data) {
        assigned[k] = k
    }
    return assigned
}
