package hierdenc

import (
    "math"
)

/* Determine HIERDENC cut-off level. */

func InList(i int, c []int) bool {
    for _, cluster := range(c) {
        if i == cluster {
            return true
        }
    }
    return false
}

func findMin(minbi []float64) float64 {
    min := minbi[0]
    for _, v := range(minbi) {
        if v < min {
            min = v
        }
    }
    return min
}

func avgScore(scores []float64) float64{
    total := 0.0
    for _, v := range scores {
        total += v
    }
    avg := total/float64(len(scores))
    return avg
}

// Return a list of objects contained within a given cluster
func getObjectsWithinCluster(target int, clusters map[int]int) []int {
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
func getUnrelatedClusters(target int, clusters map[int]int) []int {
    var unrelated []int
    // Object ID, Cluster ID
    for _,v := range(clusters) {
        if v != target && InList(v, unrelated) == false {
            unrelated = append(unrelated, v)
        }
    }
    return unrelated
}

func CalculateSilhouetteScore(clusters map[int]int, objects map[int][]int) float64 {
    // For object calculate average distance between objects in it's cluster (ai)
    var ai map[int]float64
    ai = make(map[int]float64)

    // Find ai
    // Object ID, Cluster ID
    for k,v := range(clusters) {
        var distances []float64
        sisters := getObjectsWithinCluster(v, clusters)
        for _, s := range(sisters) {
            // Calculate distance between objects
            hd, _ := HammingDistance(objects[k], objects[s])
            distances = append(distances, float64(hd))
        }

        // find average distance and add to map
        average := avgScore(distances)
        ai[k] = average
    }

    // For object and any cluster NOT containing the object, calculate the
    // objects average distance of all objects in the given cluster. Find the
    // minimum such value with respoect to all clusters (bi)
    var bi map[int]float64
    bi = make(map[int]float64)

    // Find bi
    // Object ID, Cluster ID
    for k,v := range(clusters) {
        var minbi []float64
        // Object ID, Cluster ID
        unrelated := getUnrelatedClusters(v, clusters)
        for _,c := range(unrelated) {
            var distances []float64
            sisters := getObjectsWithinCluster(c, clusters)
            for _, s := range(sisters) {
                // Calculate distance between objects
                hd, _ := HammingDistance(objects[k], objects[s])
                distances = append(distances, float64(hd))
            }
            // find average distance and add to map
            average := avgScore(distances)
            minbi = append(minbi, average)
        }
        min := findMin(minbi)
        bi[k] = min
    }



    // For each object, calculate a silouette coefficient (si).å
    // si = (bi - ai) / max(ai, bi)
    var sc []float64
    for k,_ := range(clusters) {
        si := ((bi[k] - ai[k]) / math.Max(bi[k],ai[k]))
        sc = append(sc, si)
    }

    // Average the list and return score
    average :=  avgScore(sc)
    return average
}
