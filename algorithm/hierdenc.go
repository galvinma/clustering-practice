package hierdenc

import (
    "log"
    // "reflect"

)



func HIERDENC(objects map[int][]int) map[int]int{
    // Initialize variables
    log.Println("Entering HIERDENC function...")
    r := 1            // radius of hypercubes
    // k := 0            // number of leaf clusters
    // kr := 0           // number of clusters at level r
    // var U []int        // set of IDs already sorted into clusters

    // Dictionary to keep all clusters and hypercubes
    var clusters map[int]int
    clusters = make(map[int]int)

    // Get HIERDENC density index
    index := HierdencIndex(objects, r)
    log.Println(index)
    // n := len(index)           // number of objects
    // m := len(index[0])        // number of catagorical attributes

    // for r < 2 {
        id := 1
        for _, object := range index {
            // If density is 1, increase r
            // Density of 1 implies no other objects within r
            // if index[object].Density <= 1 {
            //     r = r + 1
            //     // Update the index
            //     index = HierdencIndex(objects, r)
            // } else {
            // Cluster ID in the map

            // If object not in map, create a new cluster
            if clusters[object.ID] == 0 {
                clusters[object.ID] = id
                for k,v := range objects {
                    hd, _ := HammingDistance(objects[object.ID], v)
                    if hd <= r {
                        // Add object to current cluster
                        clusters[k] = id
                    }
                }
                id++
            }


            // }
        }
    // }
    // log.Println(U)
    return clusters
}
