package hierdenc

func MergeClusters(clusters map[int]int, objects map[int][]int, r int) {
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
