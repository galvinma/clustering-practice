package hierdenc

import (
    "log"

)

func mergeClusters(clusters map[int]int, objects map[int][]int, r int) {
    // object id, cluster id
    for k, v := range(clusters) {
        // object id, cluster id
        for m, n := range(clusters) {
            // Test if object is within r
            hd, _ := HammingDistance(objects[k], objects[m])
            if hd <= r {
                // Check if objects are in the same cluster
                if v != n {
                    // If not in the same cluster, merge
                    // object id, cluster id
                    for x, y := range(clusters) {
                        if y == n {
                            clusters[x] = v
                        }
                    }
                }
            }
        }
    }
}

func HIERDENC(objects map[int][]int) map[int]int{
    // Initialize variables
    log.Println("Entering HIERDENC function...")
    r := 1            // radius of hypercubes

    // Dictionary to keep all clusters and hypercubes
    var clusters map[int]int
    clusters = make(map[int]int)

    // Get HIERDENC density index
    index := HierdencIndex(objects, r)
    n := len(index)           // number of objects
    m := len(objects[0])        // number of catagorical attributes
    log.Println("There are", n, "objects in S with", m, "catagorical attributes.")

    // Continue until edge of cube OR 98% coverage
    for r < m && float64(len(clusters)/len(objects)) < 0.98 {
        // Start cluster count @ 1.
        // Map lookup relies on nil value of 0 for non-existant key
        id := 1
        for _, object := range index {
            // If density is 1, increase r
            // Density of 1 implies no other objects within r
            if object.Density <= 1 {
                r = r + 1
                // Update the index
                index = HierdencIndex(objects, r)

                // Merge clusters here, then break
                mergeClusters(clusters, objects, r)
                break
            }

            // If object not in map, check if it can be added to an existing cluster
            if clusters[object.ID] == 0 {
              // for obj #, cluster #
              for k,v := range clusters {
                  hd, _ := HammingDistance(objects[object.ID], objects[k])
                  if hd <= r {
                      // Add object to found cluster
                      clusters[object.ID] = v
                      break
                  }
              }
            }

            // If object cannot be added to an existing cluster, create a new cluster
            if clusters[object.ID] == 0 {
                clusters[object.ID] = id
                // for obj #, attributes
                for k,v := range objects {
                    hd, _ := HammingDistance(objects[object.ID], v)
                    if hd <= r {
                        // Add object to current cluster
                        clusters[k] = id
                    }
                }
                // Increment the counter
                id++
            }
        }
    }
    return clusters
}
