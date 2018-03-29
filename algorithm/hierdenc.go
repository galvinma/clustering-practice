package hierdenc

import (
    "log"

)

func mergeClusters() {
    // Get all objects associated with both input clusters



    // If any two objects are within r, merge clusters and return

}

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
    n := len(index)           // number of objects
    m := len(objects[0])        // number of catagorical attributes
    log.Println("There are", n, "objects in S with", m, "catagorical attributes.")

    for r < 5 {
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
