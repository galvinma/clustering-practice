package hierdenc
//
// import (
//     "log"
// )

// m & n represent the clusters to be merged
func mergeMaps(m map[int][]int, n map[int][]int) map[int][]int {
    for k, v := range(n) {
        m[k] = v
    }
    return  m
}

// a & b are cluster IDs in cluster map
func mergeClusters(o map[int]map[int][]int, a int, b int) map[int]map[int][]int {
    joined = mergeMaps(o[a], o[b])
    delete(o, a)
    delete(o, b)
    o[a] = joined
    return o
}

func FindClusters(objects map[int][]int, r int) {
    // Initialize variables
    r := 1            // radius of hypercubes
    k := 0            // number of leaf clusters
    kr := 0           // number of clusters at level r
    U := [][]int        // set of hypercube centers

    // Dictionary to keep all clusters and hypercubes
    var clusters = map[int]map[int][]int

    // Get densest hypercube and HIERDENC density index
    index := hierdenc.HierdencIndex(objects, r)
    n := len(index)           // number of objects
    m := len(index[0])        // number of catagorical attributes

    for r < m {
        // If density is 1, increase r
        if index[densest] <= 1 {
            r = r + 1
            // update the index
        }
        // Merge clusters connected via distance r...
              // code here...

        // Find a new leaf cluster
        k = k + 1           // expand the number of leaf clusters by 1
        Gk = append(Gk, )




    }


}




// Create a new leaf cluster.
//
//
// Try and move to the denest hypercube with radius r.
// Repeat until no available hypercube's within radius r.
//
//     Collect the hypercube's cells.
//
//
// Cluster until r = m, or  density(R) < 1%.
//
//
//
// Globals
// r    // radius of hypercubes
// R    // set of unclustered cells
// k    // number of leaf clusters
// kr   // number of clusters at level r
// Gk   // kth cluster
// U    // set of hypercube centers
