package hierdenc

import (
    "math"

    "github.com/galvinma/agora/algorithm/common"
)

func CalculateSilhouetteScore(clusters map[int]int, objects map[int][]int) float64 {
    // For object calculate average distance between objects in it's cluster (ai)
    var ai map[int]float64
    ai = make(map[int]float64)

    // Find ai
    // Object ID, Cluster ID
    for k,v := range(clusters) {
        var distances []float64
        sisters := GetObjectsWithinCluster(v, clusters)
        for _, s := range(sisters) {
            // Calculate distance between objects
            hd, _ := common.HammingDistance(objects[k], objects[s])
            distances = append(distances, float64(hd))
        }

        // find average distance and add to map
        average := AvgScore(distances)
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
        unrelated := GetUnrelatedClusters(v, clusters)
        for _,c := range(unrelated) {
            var distances []float64
            sisters := GetObjectsWithinCluster(c, clusters)
            for _, s := range(sisters) {
                // Calculate distance between objects
                hd, _ := common.HammingDistance(objects[k], objects[s])
                distances = append(distances, float64(hd))
            }
            // find average distance and add to map
            average := AvgScore(distances)
            minbi = append(minbi, average)
        }
        if len(minbi) > 0 {
          min := FindMin(minbi)
          bi[k] = min
        }
    }



    // For each object, calculate a silouette coefficient (si).å
    // si = (bi - ai) / max(ai, bi)
    var sc []float64
    for k,_ := range(clusters) {
        si := ((bi[k] - ai[k]) / math.Max(bi[k],ai[k]))
        sc = append(sc, si)
    }

    // Average the list and return score
    average :=  AvgScore(sc)
    return average
}
