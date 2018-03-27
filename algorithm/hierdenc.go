package hierdenc

// import (
//     "bufio"
//     "fmt"
//     "os"
//     "log"
//     "path/filepath"
// )

func FindClusters(objects map[int][]int, r int) {
    // Initialize variables
    r := 1            // radius of hypercubes
    k := 0            // number of leaf clusters
    kr := 0           // number of clusters at level r
    Gk := []int       // kth cluster
    U := [][]int        // set of hypercube centers

    // Get densest hypercube and HIERDENC density index
    densest, index := hierdenc.HierdencIndex(objects, r)
    n := len(index)           // number of objects
    // should come up with a better way of grabbing number of attributes. This is
    // unnecessary and will increase runtime.
    m := len(index[0])        // number of catagorical attributes

    for r < m {
        // If density is 1, increase r
        if index[densest] <= 1 {
            r = r + 1
        }
        // Merge clusters connected via distance r...
              // code here...

        // Find a new leaf cluster
        k = k + 1           // expand the number of leaf clusters by 1
        Gk = append(Gk, )




    }


}




// Create a new leaf cluster.


// Try and move to the denest hypercube with radius r.
// Repeat until no available hypercube's within radius r.

    // Collect the hypercube's cells.


// Cluster until r = m, or  density(R) < 1%.



// Globals
// r    // radius of hypercubes
// R    // set of unclustered cells
// k    // number of leaf clusters
// kr   // number of clusters at level r
// Gk   // kth cluster
// U    // set of hypercube centers
