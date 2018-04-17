package hierdenc

import (
    "log"

)

func HIERDENC(objects map[int][]int) map[int]int{
    // Initialize variables
    log.Println("Entering HIERDENC function...")
    r := 1            // radius of hypercubes
    id := 1           // cluster ID counter

    // Dictionary to keep all clusters and hypercubes
    var clusters map[int]int
    clusters = make(map[int]int)

    // Dependent on connectivity score check
    proceed := true
    // Worst silhouette score is -1
    // var sil float64

    // Get HIERDENC density index
    index := HierdencIndex(objects, r)
    n := len(index)           // number of objects
    m := len(objects[0])        // number of catagorical attributes
    log.Println("There are", n, "objects in S with", m, "catagorical attributes.")

    // Continue until edge of cube OR 99% coverage
    for r < 7 && float64(len(clusters)/len(objects)) < 0.99 && proceed == true {
        clusters, index, id = Cluster(clusters, objects, index, r, id)
        MergeClusters(clusters, objects, r)
        r = r + 1

    }

    // Assign leftover points to the closest cluster
        // TO DO
    return clusters
}
