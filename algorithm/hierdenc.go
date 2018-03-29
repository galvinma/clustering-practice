package hierdenc

import (
    "log"
    // "reflect"

)

// // m & n represent the clusters to be merged
// func mergeMaps(m map[int][]int, n map[int][]int) map[int][]int {
//     for k, v := range(n) {
//         m[k] = v
//     }
//     return  m
// }
//
// // a & b are cluster IDs in cluster map
// func mergeClusters(o map[int]map[int][]int, a int, b int) map[int]map[int][]int {
//     joined = mergeMaps(o[a], o[b])
//     delete(o, a)
//     delete(o, b)
//     o[a] = joined
//     return o
// }

// Check if a given object has been assigned a cluster
// func checkAssigned(id int, l []int) bool {
//   for i := range l {
//       if i == id {
//           return true
//       }
//   }
//   return false
// }


func HIERDENC(objects map[int][]int) map[int][]int{
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



            if clusters[object.ID] == 0 {
                clusters[object.ID] = id
                id++
            }

            for k,v := range objects {
                hd, _ := HammingDistance(objects[object.ID], v)
                if hd <= r {
                    // Add object to the list of sorted IDs
                    U = append(U, k)
                    // append to the cluster ID the new object
                    clusters[id] = append(clusters[id], k)
                }
            }


            // }
        }
    // }
    // log.Println(U)
    return clusters
}
