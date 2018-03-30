package hierdenc

func InList(i int, c []int) bool {
    for _, cluster := range(c) {
        if i == cluster {
            return true
        }
    }
    return false
}

func FindMin(minbi []float64) float64 {
    min := minbi[0]
    for _, v := range(minbi) {
        if v < min {
            min = v
        }
    }
    return min
}

func AvgScore(scores []float64) float64{
    total := 0.0
    for _, v := range scores {
        total += v
    }
    avg := total/float64(len(scores))
    return avg
}

// Return a list of objects contained within a given cluster
func GetObjectsWithinCluster(target int, clusters map[int]int) []int {
    var sisters []int
    // Object ID, Cluster ID
    for k,v := range(clusters) {
        if v == target {
            sisters = append(sisters, k)
        }
    }
    return sisters
}

// Return a list of clusters an object does NOT belong to
func GetUnrelatedClusters(target int, clusters map[int]int) []int {
    var unrelated []int
    // Object ID, Cluster ID
    for _,v := range(clusters) {
        if v != target && InList(v, unrelated) == false {
            unrelated = append(unrelated, v)
        }
    }
    return unrelated
}
